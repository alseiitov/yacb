package grpc

import (
	"context"
	"github.com/alseiitov/yacb/service_crypto_currency/internal/entity"
	"github.com/alseiitov/yacb/service_crypto_currency/internal/usecase"
	"github.com/alseiitov/yacb/service_crypto_currency/proto/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type cryptoCurrencyServer struct {
	uc usecase.CryptoCurrency
	pb.UnimplementedCryptoCurrencyServiceServer
}

func NewCryptoCurrencyServer(uc usecase.CryptoCurrency) *cryptoCurrencyServer {
	return &cryptoCurrencyServer{
		uc: uc,
	}
}

func (s *cryptoCurrencyServer) GetCryptoCurrencies(ctx context.Context, in *emptypb.Empty) (*pb.GetCurrenciesResponse, error) {

	currencies, err := s.uc.GetCryptoCurrencies(ctx)
	if err != nil {
		return nil, err
	}

	return entity.CryptoCurrencyListToGrpc(currencies), err
}
