package telegram_handler

import (
	"context"
	"fmt"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/pkg/tgbotapi_handler"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (h *handler) cmdMyUserInfo(w tgbotapi_handler.Writer, r *tgbotapi_handler.Request) {
	user, err := h.userUseCase.GetUserInfo(context.Background(), r.ChatID)
	if err != nil {
		log.Println(err)
		return
	}

	text := fmt.Sprintf(msgUserInfo, user.ID, user.Password)
	msg := tgbotapi.NewMessage(r.ChatID, text)
	msg.ParseMode = "Markdown"
	w.Send(msg)
	w.Request(tgbotapi.NewDeleteMessage(r.ChatID, r.Message.MessageID))
}
