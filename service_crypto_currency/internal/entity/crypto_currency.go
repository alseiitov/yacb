package entity

import "github.com/alseiitov/yacb/service_crypto_currency/proto/pb"

type (
	CryptoCurrency struct {
		ID     int64
		Symbol string
		Name   string
	}
)

func CryptoCurrencyListToGrpc(currencies []CryptoCurrency) *pb.GetCurrenciesResponse {

	currenciesResp := make([]*pb.CryptoCurrency, 0, len(currencies))

	for _, currency := range currencies {
		c := &pb.CryptoCurrency{
			Id:     currency.ID,
			Symbol: currency.Symbol,
			Name:   currency.Name,
		}

		currenciesResp = append(currenciesResp, c)
	}

	return &pb.GetCurrenciesResponse{CryptoCurrencies: currenciesResp}
}
