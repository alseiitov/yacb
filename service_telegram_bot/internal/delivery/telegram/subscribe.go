package telegram_handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/entity"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/pkg/buttons"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/pkg/currencies"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/pkg/intervals"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/pkg/tgbotapi_handler"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (h *handler) cmdSubscribe(w tgbotapi_handler.Writer, r *tgbotapi_handler.Request) {
	msg := tgbotapi.NewMessage(r.ChatID, msgChooseCurrencyToSubscribe)
	msg.ReplyMarkup = buttons.GetSubscribeCurrenciesButtons(r.ChatID)
	w.Send(msg)
	w.Request(tgbotapi.NewDeleteMessage(r.ChatID, r.Message.MessageID))
}

func (h *handler) chooseIntervalToSubscribe(w tgbotapi_handler.Writer, r *tgbotapi_handler.Request) {
	var btnData buttons.SubButtonData
	err := json.Unmarshal([]byte(r.CallbackData()), &btnData)
	if err != nil {
		log.Println(err)
		return
	}

	msg := tgbotapi.NewMessage(r.ChatID, msgChooseIntervalToSubscribe)
	msg.ReplyMarkup = buttons.GetSubscribeIntervalButtons(btnData.UserID, btnData.Currency)
	w.Send(msg)
	w.Request(tgbotapi.NewDeleteMessage(r.ChatID, r.CallbackQuery.Message.MessageID))
}

func (h *handler) createSubscription(w tgbotapi_handler.Writer, r *tgbotapi_handler.Request) {
	var btnData buttons.SubButtonData
	err := json.Unmarshal([]byte(r.CallbackData()), &btnData)
	if err != nil {
		log.Println(err)
		return
	}

	currency := currencies.Currencies[btnData.Currency]
	interval := intervals.Intervals[btnData.Interval]

	sub := entity.Subscription{
		UserID:         btnData.UserID,
		Symbol:         currency.Symbol,
		UpdateInterval: interval.Duration,
	}

	_, err = h.subscriptionUseCase.Subscribe(context.Background(), sub)
	if err != nil {
		log.Println(err)
		return
	}

	text := fmt.Sprintf(msgSubscribed, currency.Name, currency.Symbol, interval.Name)
	w.Send(tgbotapi.NewMessage(r.ChatID, text))
	w.Request(tgbotapi.NewDeleteMessage(r.ChatID, r.CallbackQuery.Message.MessageID))
}
