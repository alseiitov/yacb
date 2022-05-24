package buttons

import (
	"encoding/json"
	"fmt"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/entity"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/pkg/currencies"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/pkg/intervals"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type SubButtonData struct {
	UserID   int64               `json:"userID"`
	Currency currencies.Currency `json:"currency"`
	Interval intervals.Interval  `json:"interval"`
}

type UnsubscribeButtonData struct {
	SubscriptionID int64 `json:"id"`
}

func GetSubscribeCurrenciesButtons(userID int64) tgbotapi.InlineKeyboardMarkup {
	rows := make([][]tgbotapi.InlineKeyboardButton, 0, len(currencies.List))
	for _, curr := range currencies.List {
		s := SubButtonData{UserID: userID, Currency: curr}
		b, _ := json.Marshal(s)
		btn := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(
			fmt.Sprintf("%s (%s)", currencies.Currencies[curr].Name, currencies.Currencies[curr].Symbol), string(b),
		))
		rows = append(rows, btn)
	}
	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}

func GetSubscribeIntervalButtons(userID int64, currency currencies.Currency) tgbotapi.InlineKeyboardMarkup {
	rows := make([][]tgbotapi.InlineKeyboardButton, 0, len(intervals.List))
	for _, inter := range intervals.List {
		s := SubButtonData{UserID: userID, Currency: currency, Interval: inter}
		b, _ := json.Marshal(s)
		btn := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(
			intervals.Intervals[inter].Name, string(b),
		))
		rows = append(rows, btn)
	}
	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}

func GetUnsubscribeCurrenciesButtons(subscriptions []entity.Subscription) tgbotapi.InlineKeyboardMarkup {
	rows := make([][]tgbotapi.InlineKeyboardButton, 0, len(subscriptions))
	for _, sub := range subscriptions {
		s := UnsubscribeButtonData{SubscriptionID: sub.ID}
		b, _ := json.Marshal(s)
		_, curr := currencies.GetBySymbol(sub.Symbol)
		_, inter := intervals.GetByDuration(sub.UpdateInterval)
		btn := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(
			fmt.Sprintf("%s (%s) - %s", curr.Name, sub.Symbol, inter.Name), string(b),
		))
		rows = append(rows, btn)
	}
	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}
