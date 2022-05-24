package grpc

import (
	"context"
	"github.com/alseiitov/yacb/service_crypto_currency/internal/entity"
	"github.com/alseiitov/yacb/service_crypto_currency/internal/usecase"
	"github.com/alseiitov/yacb/service_crypto_currency/proto/pb"
)

type rateServer struct {
	uc usecase.Rate
	pb.UnimplementedRateServiceServer
}

func NewRateServer(uc usecase.Rate) *rateServer {
	return &rateServer{
		uc: uc,
	}
}

func (r *rateServer) GetCurrentRate(ctx context.Context, in *pb.GetCurrentRateRequest) (*pb.Rate, error) {

	rate, err := r.uc.GetCurrentRate(ctx, in.GetCryptoCurrencySymbol())
	if err != nil {
		return nil, err
	}

	return entity.RateToGrpc(rate), nil
}

func (r *rateServer) GetRateChange(ctx context.Context, in *pb.GetRateChangeRequest) (*pb.GetRateChangeResponse, error) {

	rateChange, err := r.uc.GetRateChange(ctx, in.GetCryptoCurrencySymbol(), in.GetPeriod().AsDuration())
	if err != nil {
		return nil, err
	}

	return entity.RateChangeToGrpc(rateChange), nil
}
