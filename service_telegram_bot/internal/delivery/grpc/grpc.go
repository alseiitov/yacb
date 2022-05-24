package grpc

import (
	"github.com/alseiitov/yacb/service_telegram_bot/internal/usecase"
	"github.com/alseiitov/yacb/service_telegram_bot/proto/pb"
	"google.golang.org/grpc"
)

type ServiceServers struct {
	Subscription pb.SubscriptionServiceServer
}

func NewServiceServers(subscriptionUseCase usecase.Subscription) *ServiceServers {
	return &ServiceServers{
		Subscription: NewSubscriptionServer(subscriptionUseCase),
	}
}

func (s *ServiceServers) Register(server *grpc.Server) {
	pb.RegisterSubscriptionServiceServer(server, s.Subscription)
}
