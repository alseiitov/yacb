package app

import (
	"context"
	"github.com/alseiitov/yacb/pkg/postgres"
	"github.com/alseiitov/yacb/service_telegram_bot/config"
	telegram_handler "github.com/alseiitov/yacb/service_telegram_bot/internal/delivery/telegram"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/pkg/tgbotapi_handler"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/scheduler"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/usecase"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/usecase/clients"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/usecase/repo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func Run(cfg *config.Config) {

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

	usersUseCase := usecase.NewUsersUseCase(repo.NewUserRepo(pgxConn))
	subscriptionsUseCase := usecase.NewSubscriptionsUseCase(repo.NewSubscriptionRepo(pgxConn), scheduler)

	subs, err := subscriptionsUseCase.GetAllSubscriptions(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	scheduler.AddSubscriptions(subs...)

	go scheduler.Run(context.Background())

	tgHandler := telegram_handler.New(
		usersUseCase,
		subscriptionsUseCase,
		usecase.NewRatesUseCase(cryptoCurrencyClient),
	)

	updateCfg := tgbotapi.NewUpdate(0)
	updateCfg.Timeout = cfg.Telegram.UpdateTimeout

	tgMux := tgbotapi_handler.New(bot, updateCfg)
	tgHandler.InitRoutes(tgMux)

	tgMux.ListenAndServe()
}
