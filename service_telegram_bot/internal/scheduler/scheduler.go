package scheduler

import (
	"context"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/entity"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/pkg/currencies"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/pkg/intervals"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/robfig/cron/v3"
	"log"
	"sync"
	"time"
)

//go:generate mockgen -source=scheduler.go -destination=mocks/mock.go

type Scheduler interface {
	AddSubscriptions(subs ...entity.Subscription)
	AddSubscription(sub entity.Subscription)
	DeleteSubscription(id int64)
	InitJobs(ctx context.Context, cron *cron.Cron) error
}

type CryptoCurrencyClient interface {
	GetRateChange(ctx context.Context, symbol string, period time.Duration) (entity.RateChange, error)
}

type scheduler struct {
	sync.Mutex
	bot                  *tgbotapi.BotAPI
	subscriptions        map[currencies.Currency]map[intervals.Interval][]entity.Subscription
	cryptoCurrencyClient CryptoCurrencyClient
}

func New(bot *tgbotapi.BotAPI, cryptoCurrencyClient CryptoCurrencyClient) *scheduler {
	subscriptions := make(map[currencies.Currency]map[intervals.Interval][]entity.Subscription)
	for curr := range currencies.Currencies {
		subscriptions[curr] = make(map[intervals.Interval][]entity.Subscription)
	}

	return &scheduler{
		bot:                  bot,
		Mutex:                sync.Mutex{},
		subscriptions:        subscriptions,
		cryptoCurrencyClient: cryptoCurrencyClient,
	}
}

func (s *scheduler) AddSubscriptions(subs ...entity.Subscription) {
	for _, sub := range subs {
		s.AddSubscription(sub)
	}
}

func (s *scheduler) AddSubscription(sub entity.Subscription) {
	s.Lock()
	defer s.Unlock()
	currency, _ := currencies.GetBySymbol(sub.Symbol)
	interval, _ := intervals.GetByDuration(sub.UpdateInterval)
	users := append(s.subscriptions[currency][interval], sub)
	s.subscriptions[currency][interval] = users
}

func (s *scheduler) DeleteSubscription(id int64) {
	s.Lock()
	defer s.Unlock()

	// TODO: refactor
	for curr, currSubs := range s.subscriptions {
		for inter, interSubs := range currSubs {
			for i, sub := range interSubs {
				if sub.ID == id {
					s.subscriptions[curr][inter] = removeInt(interSubs, i)
					return
				}
			}
		}
	}
}

func (s *scheduler) InitJobs(ctx context.Context, cron *cron.Cron) error {

	worker := func(inter intervals.Interval, curr currencies.Currency) {
		<-time.After(30 * time.Second)

		duration := intervals.Intervals[inter].Duration
		symbol := currencies.Currencies[curr].Symbol

		subs := s.subscriptions[curr][inter]
		if len(subs) == 0 {
			return
		}

		rateChange, err := s.cryptoCurrencyClient.GetRateChange(ctx, symbol, duration)
		if err != nil {
			log.Println(err)
			return
		}

		text := rateChange.String()
		for _, sub := range subs {
			msg := tgbotapi.NewMessage(sub.UserID, text)
			s.bot.Send(msg)
		}
	}

	for inter, spec := range сronSpecs {
		err := func(inter intervals.Interval, spec string) error {
			_, err := cron.AddFunc(spec, func() {
				for curr := range currencies.Currencies {
					go worker(inter, curr)
				}
			})
			return err
		}(inter, spec)
		if err != nil {
			return err
		}
	}
	return nil
}

func removeInt(s []entity.Subscription, i int) []entity.Subscription {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
