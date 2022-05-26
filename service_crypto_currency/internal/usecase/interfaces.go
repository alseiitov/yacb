package usecase

import (
	"context"
	"github.com/alseiitov/yacb/service_crypto_currency/internal/entity"
	"time"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/mock.go

type (
	CryptoCurrency interface {
		GetCryptoCurrencies(ctx context.Context) ([]entity.CryptoCurrency, error)
	}

	CryptoCurrencyRepo interface {
		GetCryptoCurrencies(ctx context.Context) ([]entity.CryptoCurrency, error)
	}

	Rate interface {
		RunRatesParser() error
		Create(ctx context.Context, rate entity.Rate) error
		GetCurrentRate(ctx context.Context, symbol string) (entity.Rate, error)
		GetRateChange(ctx context.Context, symbol string, period time.Duration) (entity.RateChange, error)
	}

	RateRepo interface {
		Create(ctx context.Context, rate entity.Rate) error
		GetRateByDate(ctx context.Context, symbol string, datetime time.Time) (entity.Rate, error)
	}

	RateClient interface {
		GetCurrentRate(ctx context.Context, symbol string) (entity.Rate, error)
	}
)
