package usecase

import (
	"context"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/entity"
	"time"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/mock.go

type (
	User interface {
		IsRegistered(ctx context.Context, id int64) (bool, error)
		Register(ctx context.Context, userID int64) error
		Unregister(ctx context.Context, id int64) error
		GetUserInfo(ctx context.Context, id int64) (entity.User, error)
	}

	UserRepo interface {
		IsRegistered(ctx context.Context, id int64) (bool, error)
		Create(ctx context.Context, user entity.User) error
		Delete(ctx context.Context, id int64) error
		GetUserInfo(ctx context.Context, id int64) (entity.User, error)
	}

	Subscription interface {
		Subscribe(ctx context.Context, subscription entity.Subscription) (int64, error)
		GetAllSubscriptions(ctx context.Context) ([]entity.Subscription, error)
		GetUserSubscriptions(ctx context.Context, userID int64) ([]entity.Subscription, error)
		Unsubscribe(ctx context.Context, id int64) error
	}

	SubscriptionRepo interface {
		Create(ctx context.Context, subscription entity.Subscription) (int64, error)
		GetAllSubscriptions(ctx context.Context) ([]entity.Subscription, error)
		GetUserSubscriptions(ctx context.Context, userID int64) ([]entity.Subscription, error)
		Delete(ctx context.Context, id int64) error
	}

	Rate interface {
		GetCurrentRates(ctx context.Context) ([]entity.Rate, error)
	}

	CryptoCurrencyClient interface {
		GetCurrentRate(ctx context.Context, currency entity.Currency) (entity.Rate, error)
		GetRateChange(ctx context.Context, symbol string, period time.Duration) (entity.RateChange, error)
	}
)
