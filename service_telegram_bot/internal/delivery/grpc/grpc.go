package grpc

import (
	"github.com/alseiitov/yacb/service_crypto_currency/proto/pb"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/usecase"
	"google.golang.org/grpc"
)

type ServiceServers struct {
	User         pb.UserServiceServer
	Subscription pb.SubscriptionServiceServer
}

func NewServiceServers(userUseCase usecase.User, subscriptionUseCase usecase.Subscription) *ServiceServers {
	return &ServiceServers{
		User:         NewUserServer(userUseCase),
		Subscription: NewSubscriptionServer(subscriptionUseCase),
	}
}

func (s *ServiceServers) Register(server *grpc.Server) {
	pb.RegisterUserServiceServer(server, s.User)
	pb.RegisterSubscriptionServiceServer(server, s.Subscription)
}
