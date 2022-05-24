package http

import (
	"context"
	"fmt"
	"github.com/alseiitov/yacb/service_crypto_currency/proto/pb"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

type Gateway struct {
	mux    *runtime.ServeMux
	conn   *grpc.ClientConn
	Server *http.Server
}

func NewGatewayServer(target string, port int) (*Gateway, error) {

	mux := runtime.NewServeMux()

	conn, err := grpc.DialContext(
		context.Background(),
		target,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to dial server: %w", err)
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	return &Gateway{
		mux:    mux,
		conn:   conn,
		Server: server,
	}, nil
}

func (g *Gateway) RegisterHandlers() error {

	err := pb.RegisterCryptoCurrencyServiceHandler(context.Background(), g.mux, g.conn)
	if err != nil {
		return fmt.Errorf("failed to register crypto currency service handler: %w", err)
	}

	err = pb.RegisterRateServiceHandler(context.Background(), g.mux, g.conn)
	if err != nil {
		return fmt.Errorf("failed to register rate service handler: %w", err)
	}

	err = pb.RegisterSubscriptionServiceHandler(context.Background(), g.mux, g.conn)
	if err != nil {
		return fmt.Errorf("failed to register subscription service handler: %w", err)
	}

	err = pb.RegisterUserServiceHandler(context.Background(), g.mux, g.conn)
	if err != nil {
		return fmt.Errorf("failed to register user service handler: %w", err)
	}

	return nil
}
