package repo

import (
	"context"
	"github.com/alseiitov/yacb/service_crypto_currency/internal/entity"
	"github.com/jackc/pgx/v4/pgxpool"
)

type cryptoCurrencyRepo struct {
	*pgxpool.Pool
}

func NewCryptoCurrencyRepo(pgxConn *pgxpool.Pool) *cryptoCurrencyRepo {
	return &cryptoCurrencyRepo{pgxConn}
}

func (r *cryptoCurrencyRepo) GetCryptoCurrencies(ctx context.Context) ([]entity.CryptoCurrency, error) {

	var currencies []entity.CryptoCurrency

	rows, err := r.Query(ctx, queryGetCurrencies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var currency entity.CryptoCurrency

		err := rows.Scan(
			&currency.ID,
			&currency.Symbol,
			&currency.Name,
		)
		if err != nil {
			return nil, err
		}

		currencies = append(currencies, currency)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return currencies, nil
}
