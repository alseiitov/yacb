package usecase

type UseCases struct {
	Rate           Rate
	CryptoCurrency CryptoCurrency
}

func New(
	rateRepo RateRepo,
	cryptoCurrencyRepo CryptoCurrencyRepo,
	rateClient RateClient,
) *UseCases {
	return &UseCases{
		Rate:           NewRatesUseCase(rateRepo, cryptoCurrencyRepo, rateClient),
		CryptoCurrency: NewCryptoCurrencyUseCase(cryptoCurrencyRepo),
	}
}
