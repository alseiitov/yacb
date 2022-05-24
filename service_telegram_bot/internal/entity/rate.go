package entity

import (
	"fmt"
	"strings"
	"time"
)

type (
	Rate struct {
		Currency Currency
		Price    float32
		Date     time.Time
	}

	Rates []Rate

	RateChange struct {
		Currency              Currency
		PrevPrice             float32
		CurrPrice             float32
		PriceChange           float32
		PriceChangePercentage float32
		PrevDate              time.Time
		CurrDate              time.Time
	}
)

func (r Rates) String() string {
	var sb strings.Builder
	for _, rate := range r {
		sb.WriteString(rate.String())
	}
	return sb.String()
}

func (r Rate) String() string {
	return fmt.Sprintf(
		"%s (%s):\n%s USDT\n\n",
		r.Currency.Name, r.Currency.Symbol, floatToStr(r.Price),
	)
}

func (r RateChange) String() string {
	var sign string
	if r.PriceChange > 0 {
		sign = "+"
	}
	return fmt.Sprintf("%s / %s\n%s USDT / %s\n%s USDT / %s\n%s%s USDT\n%s%.2f%%",
		r.Currency.Name,
		r.Currency.Symbol,
		floatToStr(r.PrevPrice), r.PrevDate.Format(time.Kitchen),
		floatToStr(r.CurrPrice), r.CurrDate.Format(time.Kitchen),
		sign, floatToStr(r.PriceChange),
		sign, r.PriceChangePercentage,
	)
}

func floatToStr(price float32) string {
	if abs(price) < 0.1 {
		return fmt.Sprintf("%.8f", price)
	}
	if abs(price) < 1 {
		return fmt.Sprintf("%.4f", price)
	}
	return fmt.Sprintf("%.2f", price)
}

func abs(f float32) float32 {
	if f < 0 {
		return -f
	}
	return f
}
