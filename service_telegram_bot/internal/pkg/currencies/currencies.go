package currencies

import "github.com/alseiitov/yacb/service_telegram_bot/internal/entity"

// TODO: store and parse from database

type Currency uint8

const (
	Bitcoin Currency = iota
	Ethereum
	Cardano
	Solana
)

var (
	List = []Currency{Bitcoin, Ethereum, Cardano, Solana}

	Currencies = map[Currency]entity.Currency{
		Bitcoin: {
			Symbol: "BTC",
			Name:   "Bitcoin",
		},
		Ethereum: {
			Symbol: "ETH",
			Name:   "Ethereum",
		},
		Cardano: {
			Symbol: "ADA",
			Name:   "Cardano",
		},
		Solana: {
			Symbol: "SOL",
			Name:   "Solana",
		},
	}
)

func GetBySymbol(symbol string) (Currency, entity.Currency) {
	for curr, ent := range Currencies {
		if ent.Symbol == symbol {
			return curr, ent
		}
	}
	return 0, entity.Currency{}
}
