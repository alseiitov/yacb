package usecase

import (
	"context"
	"github.com/alseiitov/yacb/service_crypto_currency/internal/entity"
)

type CryptoCurrencyUseCase struct {
	repo CryptoCurrencyRepo
}

func NewCryptoCurrencyUseCase(r CryptoCurrencyRepo) *CryptoCurrencyUseCase {
	return &CryptoCurrencyUseCase{
		repo: r,
	}
}

func (uc *CryptoCurrencyUseCase) GetCryptoCurrencies(ctx context.Context) ([]entity.CryptoCurrency, error) {
	return uc.repo.GetCryptoCurrencies(ctx)
}
