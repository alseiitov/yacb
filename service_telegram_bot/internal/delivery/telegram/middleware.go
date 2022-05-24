package telegram_handler

import (
	"context"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/pkg/tgbotapi_handler"
	"log"
)

func (h *handler) registerUserMdw(fn tgbotapi_handler.HandlerFunc) tgbotapi_handler.HandlerFunc {
	return func(w tgbotapi_handler.Writer, r *tgbotapi_handler.Request) {
		registered, err := h.userUseCase.IsRegistered(context.Background(), r.ChatID)
		if err != nil {
			log.Println(err)
			return
		}
		if !registered {
			err = h.userUseCase.Register(context.Background(), r.ChatID)
			if err != nil {
				log.Println(err)
				return
			}
		}
		fn(w, r)
	}
}
