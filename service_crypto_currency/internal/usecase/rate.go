package usecase

import (
	"context"
	"fmt"
	"github.com/alseiitov/yacb/service_crypto_currency/internal/entity"
	"github.com/robfig/cron/v3"
	"log"
	"sync"
	"time"
)

const (
	cronEveryFiveMinute = "*/5 * * * *"
)

type RateUseCase struct {
	rateRepo           RateRepo
	cryptoCurrencyRepo CryptoCurrencyRepo
	rateClient         RateClient
}

func NewRatesUseCase(rateRepo RateRepo, cryptoCurrencyRepo CryptoCurrencyRepo, rateClient RateClient) *RateUseCase {
	return &RateUseCase{
		rateRepo:           rateRepo,
		cryptoCurrencyRepo: cryptoCurrencyRepo,
		rateClient:         rateClient,
	}
}

func (uc *RateUseCase) RunRatesParser() error {

	currencies, err := uc.cryptoCurrencyRepo.GetCryptoCurrencies(context.Background())
	if err != nil {
		return err
	}

	cron := cron.New()
	defer cron.Start()

	saveCurrentRate := func(wg *sync.WaitGroup, id int64, symbol string, date *time.Time) {
		defer wg.Done()

		rate, err := uc.GetCurrentRate(context.Background(), symbol)
		if err != nil {
			log.Println(fmt.Errorf("failed to parse rate: %w", err))
			return
		}
		rate.Date = date
		rate.CryptoCurrencyID = id

		err = uc.Create(context.Background(), rate)
		if err != nil {
			log.Println(fmt.Errorf("failed to create rate: %w", err))
			return
		}
	}

	cron.AddFunc(cronEveryFiveMinute, func() {
		var wg sync.WaitGroup
		// current time with truncated seconds
		// to not depend on the execution time of the request
		now := time.Now().UTC().Truncate(time.Minute)
		for _, currency := range currencies {
			wg.Add(1)
			go saveCurrentRate(&wg, currency.ID, currency.Symbol, &now)
		}
		wg.Wait()
	})
	return nil
}

func (uc *RateUseCase) Create(ctx context.Context, rate entity.Rate) error {
	return uc.rateRepo.Create(ctx, rate)
}

func (uc *RateUseCase) GetCurrentRate(ctx context.Context, symbol string) (entity.Rate, error) {

	rate, err := uc.rateClient.GetCurrentRate(ctx, symbol)
	if err != nil {
		return entity.Rate{}, err
	}

	return rate, nil
}

func (uc *RateUseCase) GetRateChange(ctx context.Context, symbol string, period time.Duration) (entity.RateChange, error) {

	currTime := time.Now().UTC().Truncate(5 * time.Minute)
	prevTime := currTime.Add(-period)

	currRate, err := uc.rateRepo.GetRateByDate(ctx, symbol, currTime)
	if err != nil {
		return entity.RateChange{}, fmt.Errorf("failed to get current rate: %w", err)
	}

	prevRate, err := uc.rateRepo.GetRateByDate(ctx, symbol, prevTime)
	if err != nil {
		return entity.RateChange{}, fmt.Errorf("failed to get previous rate: %w", err)
	}

	return entity.NewRateChange(currRate, prevRate), nil
}
