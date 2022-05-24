package usecase

import (
	"context"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/entity"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/scheduler"
)

type SubscriptionUseCase struct {
	repo      SubscriptionRepo
	scheduler scheduler.Scheduler
}

func NewSubscriptionsUseCase(r SubscriptionRepo, scheduler scheduler.Scheduler) *SubscriptionUseCase {
	return &SubscriptionUseCase{
		repo:      r,
		scheduler: scheduler,
	}
}

func (uc *SubscriptionUseCase) Subscribe(ctx context.Context, sub entity.Subscription) (int64, error) {
	id, err := uc.repo.Create(ctx, sub)
	if err != nil {
		return 0, err
	}

	sub.ID = id
	uc.scheduler.AddSubscription(sub)
	return id, nil
}

func (uc *SubscriptionUseCase) GetAllSubscriptions(ctx context.Context) ([]entity.Subscription, error) {
	return uc.repo.GetAllSubscriptions(ctx)
}

func (uc *SubscriptionUseCase) GetUserSubscriptions(ctx context.Context, userID int64) ([]entity.Subscription, error) {
	return uc.repo.GetUserSubscriptions(ctx, userID)
}

func (uc *SubscriptionUseCase) Unsubscribe(ctx context.Context, id int64) error {
	err := uc.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	uc.scheduler.DeleteSubscription(id)
	return nil
}
