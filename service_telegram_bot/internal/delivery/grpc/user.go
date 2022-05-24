package grpc

import (
	"context"
	"github.com/alseiitov/yacb/service_crypto_currency/internal/entity"
	"github.com/alseiitov/yacb/service_crypto_currency/internal/usecase"
	"github.com/alseiitov/yacb/service_crypto_currency/proto/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type userServer struct {
	uc usecase.User
	pb.UnimplementedUserServiceServer
}

func NewUserServer(uc usecase.User) *userServer {
	return &userServer{
		uc: uc,
	}
}

func (s *userServer) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	user := entity.User{
		Login:    in.GetLogin(),
		Password: in.GetPassword(),
	}

	id, err := s.uc.Register(ctx, user)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{ID: id}, nil
}

func (s *userServer) Unregister(ctx context.Context, in *pb.UnregisterRequest) (*emptypb.Empty, error) {

	err := s.uc.Unregister(ctx, in.GetID())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, err
}
