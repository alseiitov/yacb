package app

import (
	"fmt"
	"github.com/alseiitov/yacb/pkg/postgres"
	"github.com/alseiitov/yacb/service_crypto_currency/config"
	grpc_service "github.com/alseiitov/yacb/service_crypto_currency/internal/delivery/grpc"
	http_service "github.com/alseiitov/yacb/service_crypto_currency/internal/delivery/http"
	"github.com/alseiitov/yacb/service_crypto_currency/internal/usecase"
	"github.com/alseiitov/yacb/service_crypto_currency/internal/usecase/clients"
	"github.com/alseiitov/yacb/service_crypto_currency/internal/usecase/repo"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Run(cfg *config.Config) {

	pgxConn, err := postgres.NewPgxConn(cfg.Postgres)
	if err != nil {
		log.Fatalln(err)
	}
	defer pgxConn.Close()

	useCases := usecase.New(
		repo.NewRateRepo(pgxConn),
		repo.NewCryptoCurrencyRepo(pgxConn),
		clients.NewBinanceClient(cfg.Binance.URL),
	)

	err = useCases.Rate.RunRatesParser()
	if err != nil {
		log.Fatalln(err)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	server := grpc.NewServer()
	grpcServiceServers := grpc_service.NewServiceServers(useCases)
	grpcServiceServers.Register(server)

	// Serve gRPC server
	grpcAddress := fmt.Sprintf("0.0.0.0:%d", cfg.Server.Port)
	log.Printf("Serving gRPC on %s", grpcAddress)
	go func() {
		log.Fatalln(server.Serve(listener))
	}()

	gatewayServer, err := http_service.NewGatewayServer(grpcAddress, cfg.Server.GatewayPort)
	if err != nil {
		log.Fatalln(err)
	}

	err = gatewayServer.RegisterHandlers()
	if err != nil {
		log.Fatalln(err)
	}

	gatewayAddress := fmt.Sprintf("0.0.0.0:%d", cfg.Server.GatewayPort)
	log.Printf("Serving gRPC-Gateway on %s", gatewayAddress)
	log.Fatalln(gatewayServer.Server.ListenAndServe())
}
