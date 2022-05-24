package repo

import (
	"context"
	"github.com/alseiitov/yacb/service_crypto_currency/internal/entity"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type rateRepo struct {
	*pgxpool.Pool
}

func NewRateRepo(pgxConn *pgxpool.Pool) *rateRepo {
	return &rateRepo{pgxConn}
}

func (r *rateRepo) Create(ctx context.Context, rate entity.Rate) error {

	_, err := r.Exec(
		ctx, queryInsertRate,
		rate.CryptoCurrencyID,
		rate.Price,
		rate.Date.UTC(),
	)

	return err
}

func (r *rateRepo) GetRateByDate(ctx context.Context, symbol string, date time.Time) (entity.Rate, error) {

	var rate entity.Rate

	err := r.QueryRow(
		ctx, queryGetRateAtDatetime,
		symbol, date,
	).Scan(
		&rate.CryptoCurrencyName,
		&rate.Price,
		&rate.Date,
	)

	if err != nil {
		return entity.Rate{}, err
	}

	return rate, nil
}
