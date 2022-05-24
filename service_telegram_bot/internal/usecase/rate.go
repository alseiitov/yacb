package usecase

import (
	"context"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/entity"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/pkg/currencies"
	"log"
	"sort"
	"sync"
)

type RateUseCase struct {
	cryptoCurrencyClient CryptoCurrencyClient
}

func NewRatesUseCase(cryptoCurrencyClient CryptoCurrencyClient) *RateUseCase {
	return &RateUseCase{
		cryptoCurrencyClient: cryptoCurrencyClient,
	}
}

func (uc *RateUseCase) GetCurrentRates(ctx context.Context) ([]entity.Rate, error) {

	rates := make([]entity.Rate, 0, len(currencies.Currencies))
	wg := sync.WaitGroup{}

	for _, c := range currencies.Currencies {
		wg.Add(1)
		go func(currency entity.Currency) {
			defer wg.Done()
			rate, err := uc.cryptoCurrencyClient.GetCurrentRate(ctx, currency)
			if err != nil {
				log.Println(err)
				return
			}
			rate.Currency = currency
			rates = append(rates, rate)
		}(c)
	}

	wg.Wait()

	sort.Slice(rates, func(i, j int) bool {
		return rates[i].Price > rates[j].Price
	})

	return rates, nil
}
