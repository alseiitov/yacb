package entity

import (
	"github.com/alseiitov/yacb/service_crypto_currency/proto/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type (
	Rate struct {
		CryptoCurrencyID   int64
		CryptoCurrencyName string
		Price              float32
		Date               *time.Time
	}

	RateChange struct {
		CryptoCurrencyName    string
		PrevPrice             float32
		CurrPrice             float32
		PriceChange           float32
		PriceChangePercentage float32
		PrevDate              *time.Time
		CurrDate              *time.Time
	}
)

func NewRateChange(curr, prev Rate) RateChange {

	priceChange := curr.Price - prev.Price
	priceChangePercentage := (curr.Price * 100 / prev.Price) - 100

	return RateChange{
		CryptoCurrencyName:    curr.CryptoCurrencyName,
		PrevPrice:             prev.Price,
		CurrPrice:             curr.Price,
		PriceChange:           priceChange,
		PriceChangePercentage: priceChangePercentage,
		PrevDate:              prev.Date,
		CurrDate:              curr.Date,
	}
}

func RateToGrpc(rate Rate) *pb.Rate {
	return &pb.Rate{
		CryptoCurrencyID: rate.CryptoCurrencyID,
		Price:            rate.Price,
		Date:             timestamppb.New(*rate.Date),
	}
}

func RateChangeToGrpc(change RateChange) *pb.GetRateChangeResponse {
	return &pb.GetRateChangeResponse{
		CryptoCurrencyName:    change.CryptoCurrencyName,
		PrevPrice:             change.PrevPrice,
		CurrPrice:             change.CurrPrice,
		PriceChange:           change.PriceChange,
		PriceChangePercentage: change.PriceChangePercentage,
		PrevDate:              timestamppb.New(*change.PrevDate),
		CurrDate:              timestamppb.New(*change.CurrDate),
	}
}
