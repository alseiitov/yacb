package grpc

import (
	"context"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/entity"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/usecase"

	"github.com/alseiitov/yacb/service_crypto_currency/proto/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type subscriptionServer struct {
	uc usecase.Subscription
	pb.UnimplementedSubscriptionServiceServer
}

func NewSubscriptionServer(uc usecase.Subscription) *subscriptionServer {
	return &subscriptionServer{
		uc: uc,
	}
}

func (s *subscriptionServer) Subscribe(ctx context.Context, in *pb.SubscribeRequest) (*emptypb.Empty, error) {

	subscription := entity.Subscription{
		UserID: in.GetUserID(),
		// TODO
		Symbol:         "",
		UpdateInterval: in.GetInterval().AsDuration(),
	}

	id, err := s.uc.Subscribe(ctx, subscription)
	if err != nil {
		return nil, err
	}

	_ = id

	return &emptypb.Empty{}, nil
}

func (s *subscriptionServer) GetUserSubscriptions(ctx context.Context, in *pb.GetUserSubscriptionsRequest) (*pb.GetUserSubscriptionsResponse, error) {

	subscriptions, err := s.uc.GetUserSubscriptions(ctx, in.GetUserID())
	if err != nil {
		return nil, err
	}

	return entity.UserSubscriptionListToGrpc(subscriptions), nil
}

func (s *subscriptionServer) Unsubscribe(ctx context.Context, in *pb.UnsubscribeRequest) (*emptypb.Empty, error) {

	err := s.uc.Unsubscribe(ctx, in.GetSubscriptionID())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
