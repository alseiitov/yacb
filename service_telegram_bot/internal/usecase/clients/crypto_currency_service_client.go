package clients

import (
	"context"
	"fmt"
	"github.com/alseiitov/yacb/service_crypto_currency/proto/pb"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/entity"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/durationpb"
	"time"
)

type cryptoCurrencyServiceClient struct {
	grpcConn *grpc.ClientConn
}

func NewCryptoCurrencyServiceClient(host string, port int) (*cryptoCurrencyServiceClient, error) {

	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", host, port),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}
	//defer conn.Close()

	client := &cryptoCurrencyServiceClient{
		grpcConn: conn,
	}

	return client, nil
}

func (c *cryptoCurrencyServiceClient) GetCurrentRate(ctx context.Context, currency entity.Currency) (entity.Rate, error) {

	client := pb.NewRateServiceClient(c.grpcConn)
	rate, err := client.GetCurrentRate(ctx, &pb.GetCurrentRateRequest{CryptoCurrencySymbol: currency.Symbol})
	if err != nil {
		return entity.Rate{}, err
	}

	return entity.Rate{
		Currency: currency,
		Price:    rate.Price,
		Date:     rate.Date.AsTime(),
	}, nil

}
func (c *cryptoCurrencyServiceClient) GetRateChange(ctx context.Context, symbol string, period time.Duration) (entity.RateChange, error) {

	client := pb.NewRateServiceClient(c.grpcConn)
	rateChange, err := client.GetRateChange(ctx, &pb.GetRateChangeRequest{
		CryptoCurrencySymbol: symbol,
		Period:               durationpb.New(period),
	})
	if err != nil {
		return entity.RateChange{}, err
	}

	return entity.RateChange{
		Currency: entity.Currency{
			Symbol: symbol,
			Name:   rateChange.CryptoCurrencyName,
		},
		PrevPrice:             rateChange.PrevPrice,
		CurrPrice:             rateChange.CurrPrice,
		PriceChange:           rateChange.PriceChange,
		PriceChangePercentage: rateChange.PriceChangePercentage,
		PrevDate:              rateChange.PrevDate.AsTime(),
		CurrDate:              rateChange.CurrDate.AsTime(),
	}, nil
}
