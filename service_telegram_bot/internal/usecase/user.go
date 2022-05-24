package usecase

import (
	"context"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/entity"
	"github.com/sethvargo/go-password/password"
)

type UserUseCase struct {
	repo UserRepo
}

func NewUsersUseCase(r UserRepo) *UserUseCase {
	return &UserUseCase{
		repo: r,
	}
}

func (uc *UserUseCase) IsRegistered(ctx context.Context, id int64) (bool, error) {
	return uc.repo.IsRegistered(ctx, id)
}

func (uc *UserUseCase) Register(ctx context.Context, userID int64) error {
	pass, err := password.Generate(10, 5, 0, false, false)
	if err != nil {
		return err
	}
	user := entity.User{
		ID:       userID,
		Password: pass,
	}
	return uc.repo.Create(ctx, user)
}

func (uc *UserUseCase) Unregister(ctx context.Context, id int64) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *UserUseCase) GetUserInfo(ctx context.Context, id int64) (entity.User, error) {
	return uc.repo.GetUserInfo(ctx, id)
}
