package tgbotapi_handler

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type (
	Command         string
	CallbackMessage string

	Request struct {
		tgbotapi.Update
		ChatID int64
	}

	Writer interface {
		Send(msg tgbotapi.Chattable) (tgbotapi.Message, error)
		Request(msg tgbotapi.Chattable) (*tgbotapi.APIResponse, error)
	}

	HandlerFunc func(w Writer, r *Request)

	handler struct {
		bot       *tgbotapi.BotAPI
		updates   tgbotapi.UpdatesChannel
		callbacks map[CallbackMessage]HandlerFunc
		commands  map[Command]HandlerFunc
	}

	Handler interface {
		HandleCommand(command Command, handlerFunc HandlerFunc)
		HandleCallback(message CallbackMessage, handlerFunc HandlerFunc)
		ListenAndServe()
	}
)

func New(bot *tgbotapi.BotAPI, updateCfg tgbotapi.UpdateConfig) Handler {
	return &handler{
		bot:       bot,
		updates:   bot.GetUpdatesChan(updateCfg),
		callbacks: make(map[CallbackMessage]HandlerFunc),
		commands:  make(map[Command]HandlerFunc),
	}
}

func (h *handler) HandleCommand(command Command, handlerFunc HandlerFunc) {
	h.commands[command] = handlerFunc
}

func (h *handler) HandleCallback(callbackMessage CallbackMessage, handlerFunc HandlerFunc) {
	h.callbacks[callbackMessage] = handlerFunc
}

func (h *handler) ListenAndServe() {
	for update := range h.updates {
		sender := update.SentFrom()
		if sender == nil {
			continue
		}

		req := &Request{
			Update: update,
			ChatID: sender.ID,
		}

		var fn HandlerFunc
		var found bool

		switch {
		case update.Message != nil:
			if !update.Message.IsCommand() {
				continue
			}
			fn, found = h.commands[Command(update.Message.Command())]
		case update.CallbackQuery != nil:
			fn, found = h.callbacks[CallbackMessage(update.CallbackQuery.Message.Text)]
		}

		if !found {
			continue
		}

		fn(h.bot, req)
	}
}
