package telegram_handler

import (
	"github.com/alseiitov/yacb/service_telegram_bot/internal/pkg/tgbotapi_handler"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/usecase"
)

type handler struct {
	userUseCase         usecase.User
	subscriptionUseCase usecase.Subscription
	rateUseCase         usecase.Rate
}

func New(
	userUseCase usecase.User,
	subscriptionUseCase usecase.Subscription,
	rateUseCase usecase.Rate,
) *handler {
	return &handler{
		userUseCase:         userUseCase,
		subscriptionUseCase: subscriptionUseCase,
		rateUseCase:         rateUseCase,
	}
}

func (h *handler) InitRoutes(mux tgbotapi_handler.Handler) {
	mux.HandleCommand(cmdStart, h.registerUserMdw(h.cmdStart))
	mux.HandleCommand(cmdMyUserInfo, h.cmdMyUserInfo)
	mux.HandleCommand(cmdCurrentRates, h.cmdCurrentRates)

	mux.HandleCommand(cmdSubscribe, h.cmdSubscribe)
	mux.HandleCallback(msgChooseCurrencyToSubscribe, h.chooseIntervalToSubscribe)
	mux.HandleCallback(msgChooseIntervalToSubscribe, h.createSubscription)

	mux.HandleCommand(cmdUnsubscribe, h.cmdUnsubscribe)
	mux.HandleCallback(msgChooseCurrencyToUnsubscribe, h.deleteSubscription)

}
