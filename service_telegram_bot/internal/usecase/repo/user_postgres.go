package repo

import (
	"context"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/entity"
	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepo struct {
	*pgxpool.Pool
}

func NewUserRepo(pgxConn *pgxpool.Pool) *userRepo {
	return &userRepo{pgxConn}
}

func (r *userRepo) IsRegistered(ctx context.Context, id int64) (bool, error) {
	var exist bool

	err := r.QueryRow(ctx, queryUserExist, id).Scan(&exist)
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (r *userRepo) Create(ctx context.Context, user entity.User) error {

	_, err := r.Exec(ctx, queryInsertUser, user.ID, user.Password)
	return err
}

func (r *userRepo) Delete(ctx context.Context, id int64) error {

	_, err := r.Exec(ctx, queryDeleteUser, id)
	return err
}

func (r *userRepo) GetUserInfo(ctx context.Context, id int64) (entity.User, error) {

	var user entity.User

	err := r.QueryRow(ctx, queryGetUser, id).Scan(&user.ID, &user.Password)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
