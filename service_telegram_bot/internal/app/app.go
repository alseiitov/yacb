package app

import (
	"context"
	"fmt"
	"github.com/alseiitov/yacb/pkg/postgres"
	grpc_service "github.com/alseiitov/yacb/service_telegram_bot/internal/delivery/grpc"
	http_service "github.com/alseiitov/yacb/service_telegram_bot/internal/delivery/http"
	"github.com/robfig/cron/v3"
	"net"

	"github.com/alseiitov/yacb/service_telegram_bot/config"
	telegram_handler "github.com/alseiitov/yacb/service_telegram_bot/internal/delivery/telegram"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/pkg/tgbotapi_handler"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/scheduler"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/usecase"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/usecase/clients"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/usecase/repo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"google.golang.org/grpc"
	"log"
)

func Run(cfg *config.Config) {
	ctx := context.Background()

	pgxConn, err := postgres.NewPgxConn(cfg.Postgres)
	if err != nil {
		log.Fatalln(err)
	}
	defer pgxConn.Close()

	bot, err := tgbotapi.NewBotAPI(cfg.Telegram.APIKey)
	if err != nil {
		log.Fatalf("can't create bot: %s\n", err.Error())
	}

	cryptoCurrencyClient, err := clients.NewCryptoCurrencyServiceClient(
		cfg.CryptoCurrencyService.Host, cfg.CryptoCurrencyService.Port,
	)
	if err != nil {
		log.Fatalln(err)
	}

	scheduler := scheduler.New(bot, cryptoCurrencyClient)

	usersUseCase := usecase.NewUsersUseCase(
		repo.NewUserRepo(pgxConn),
	)
	subscriptionsUseCase := usecase.NewSubscriptionsUseCase(
		repo.NewSubscriptionRepo(pgxConn), scheduler,
	)

	subs, err := subscriptionsUseCase.GetAllSubscriptions(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	scheduler.AddSubscriptions(subs...)

	c := cron.New()
	err = scheduler.InitJobs(ctx, c)
	if err != nil {
		log.Fatalln(err)
	}
	c.Start()
	defer c.Stop()

	tgHandler := telegram_handler.New(
		usersUseCase,
		subscriptionsUseCase,
		usecase.NewRatesUseCase(cryptoCurrencyClient),
	)

	updateCfg := tgbotapi.NewUpdate(0)
	updateCfg.Timeout = cfg.Telegram.UpdateTimeout

	tgMux := tgbotapi_handler.New(bot, updateCfg)
	tgHandler.InitRoutes(tgMux)

	go func() {
		tgMux.ListenAndServe()
	}()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	server := grpc.NewServer()
	grpcServiceServers := grpc_service.NewServiceServers(subscriptionsUseCase)
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
