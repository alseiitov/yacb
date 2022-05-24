package repo

import (
	"context"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/entity"
	"github.com/jackc/pgx/v4/pgxpool"
)

type subscriptionRepo struct {
	*pgxpool.Pool
}

func NewSubscriptionRepo(pgxConn *pgxpool.Pool) *subscriptionRepo {
	return &subscriptionRepo{pgxConn}
}

func (r *subscriptionRepo) Create(ctx context.Context, subscription entity.Subscription) (id int64, err error) {

	err = r.QueryRow(
		ctx, queryInsertSubscription,
		subscription.UserID,
		subscription.Symbol,
		subscription.UpdateInterval,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return
}

func (r *subscriptionRepo) GetAllSubscriptions(ctx context.Context) ([]entity.Subscription, error) {
	var subscriptions []entity.Subscription

	rows, err := r.Query(ctx, queryGetAllSubscriptions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var subscription entity.Subscription

		err := rows.Scan(
			&subscription.ID,
			&subscription.UserID,
			&subscription.Symbol,
			&subscription.UpdateInterval,
		)
		if err != nil {
			return nil, err
		}

		subscriptions = append(subscriptions, subscription)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func (r *subscriptionRepo) GetUserSubscriptions(ctx context.Context, userID int64) ([]entity.Subscription, error) {

	var subscriptions []entity.Subscription

	rows, err := r.Query(ctx, queryGetUsersSubscriptions, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		subscription := entity.Subscription{UserID: userID}

		err := rows.Scan(
			&subscription.ID,
			&subscription.Symbol,
			&subscription.UpdateInterval,
		)
		if err != nil {
			return nil, err
		}

		subscriptions = append(subscriptions, subscription)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func (r *subscriptionRepo) Delete(ctx context.Context, id int64) error {

	_, err := r.Exec(ctx, queryDeleteSubscription, id)
	return err
}
