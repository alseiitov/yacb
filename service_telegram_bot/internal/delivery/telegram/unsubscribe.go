package telegram_handler

import (
	"context"
	"encoding/json"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/pkg/buttons"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/pkg/tgbotapi_handler"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (h *handler) cmdUnsubscribe(w tgbotapi_handler.Writer, r *tgbotapi_handler.Request) {
	subscriptions, err := h.subscriptionUseCase.GetUserSubscriptions(context.Background(), r.ChatID)
	if err != nil {
		log.Println(err)
		return
	}
	if len(subscriptions) == 0 {
		msg := tgbotapi.NewMessage(r.ChatID, msgYouHaveNoSubscriptions)
		w.Send(msg)
		w.Request(tgbotapi.NewDeleteMessage(r.ChatID, r.Message.MessageID))
		return
	}

	msg := tgbotapi.NewMessage(r.ChatID, msgChooseCurrencyToUnsubscribe)
	msg.ReplyMarkup = buttons.GetUnsubscribeCurrenciesButtons(subscriptions)
	w.Send(msg)
	w.Request(tgbotapi.NewDeleteMessage(r.ChatID, r.Message.MessageID))
}

func (h *handler) deleteSubscription(w tgbotapi_handler.Writer, r *tgbotapi_handler.Request) {
	var btnData buttons.UnsubscribeButtonData
	err := json.Unmarshal([]byte(r.CallbackData()), &btnData)
	if err != nil {
		log.Println(err)
		return
	}

	err = h.subscriptionUseCase.Unsubscribe(context.Background(), btnData.SubscriptionID)
	if err != nil {
		log.Println(err)
		return
	}

	msg := tgbotapi.NewMessage(r.ChatID, msgUnsubscribed)
	w.Send(msg)
	w.Request(tgbotapi.NewDeleteMessage(r.ChatID, r.CallbackQuery.Message.MessageID))
}
