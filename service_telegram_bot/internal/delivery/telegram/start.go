package telegram_handler

import (
	"github.com/alseiitov/yacb/service_telegram_bot/internal/pkg/tgbotapi_handler"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *handler) cmdStart(w tgbotapi_handler.Writer, r *tgbotapi_handler.Request) {
	w.Send(tgbotapi.NewMessage(r.ChatID, msgChooseCmdToStart))
	w.Request(tgbotapi.NewDeleteMessage(r.ChatID, r.Message.MessageID))
}
