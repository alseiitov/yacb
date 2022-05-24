package grpc

import (
	"github.com/alseiitov/yacb/service_crypto_currency/internal/usecase"
	"github.com/alseiitov/yacb/service_crypto_currency/proto/pb"
	"google.golang.org/grpc"
)

type ServiceServers struct {
	Rate           pb.RateServiceServer
	CryptoCurrency pb.CryptoCurrencyServiceServer
}

func NewServiceServers(useCases *usecase.UseCases) *ServiceServers {
	return &ServiceServers{
		Rate:           NewRateServer(useCases.Rate),
		CryptoCurrency: NewCryptoCurrencyServer(useCases.CryptoCurrency),
	}
}

func (s *ServiceServers) Register(server *grpc.Server) {
	pb.RegisterRateServiceServer(server, s.Rate)
	pb.RegisterCryptoCurrencyServiceServer(server, s.CryptoCurrency)
}
