package telegram_handler

import (
	"context"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/entity"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/pkg/tgbotapi_handler"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (h *handler) cmdCurrentRates(w tgbotapi_handler.Writer, r *tgbotapi_handler.Request) {
	rates, err := h.rateUseCase.GetCurrentRates(context.Background())
	if err != nil {
		log.Println(err)
		return
	}
	msg := entity.Rates(rates).String()
	w.Send(tgbotapi.NewMessage(r.ChatID, msg))
	w.Request(tgbotapi.NewDeleteMessage(r.ChatID, r.Message.MessageID))
}
