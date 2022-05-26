package usecase

import (
	"context"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/entity"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/scheduler"
	mock_scheduler "github.com/alseiitov/yacb/service_telegram_bot/internal/scheduler/mocks"
	mock_usecase "github.com/alseiitov/yacb/service_telegram_bot/internal/usecase/mocks"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"reflect"
	"testing"
	"time"
)

func TestNewSubscriptionsUseCase(t *testing.T) {
	type args struct {
		repo      SubscriptionRepo
		scheduler scheduler.Scheduler
	}
	tests := []struct {
		name string
		args args
		want *SubscriptionUseCase
	}{
		{
			name: "ok",
			args: args{
				repo:      &mock_usecase.MockSubscriptionRepo{},
				scheduler: &mock_scheduler.MockScheduler{},
			},
			want: &SubscriptionUseCase{
				repo:      &mock_usecase.MockSubscriptionRepo{},
				scheduler: &mock_scheduler.MockScheduler{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSubscriptionsUseCase(tt.args.repo, tt.args.scheduler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSubscriptionsUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubscriptionUseCase_GetAllSubscriptions(t *testing.T) {
	type fields struct {
		repo      *mock_usecase.MockSubscriptionRepo
		scheduler *mock_scheduler.MockScheduler
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		prepare func(f *fields)
		want    []entity.Subscription
		wantErr bool
	}{
		{
			name: "repo returned error",
			args: args{
				ctx: context.Background(),
			},
			prepare: func(f *fields) {
				f.repo.EXPECT().
					GetAllSubscriptions(context.Background()).
					Return(nil, errors.New("error"))
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				ctx: context.Background(),
			},
			prepare: func(f *fields) {
				f.repo.EXPECT().
					GetAllSubscriptions(context.Background()).
					Return(
						[]entity.Subscription{
							{
								ID:             1,
								UserID:         1,
								Symbol:         "BTC",
								UpdateInterval: 5 * time.Minute,
							},
						},
						nil,
					)
			},
			want: []entity.Subscription{
				{
					ID:             1,
					UserID:         1,
					Symbol:         "BTC",
					UpdateInterval: 5 * time.Minute,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				repo:      mock_usecase.NewMockSubscriptionRepo(ctrl),
				scheduler: mock_scheduler.NewMockScheduler(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			uc := &SubscriptionUseCase{
				repo:      f.repo,
				scheduler: f.scheduler,
			}

			got, err := uc.GetAllSubscriptions(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllSubscriptions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllSubscriptions() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubscriptionUseCase_GetUserSubscriptions(t *testing.T) {
	type fields struct {
		repo      *mock_usecase.MockSubscriptionRepo
		scheduler *mock_scheduler.MockScheduler
	}
	type args struct {
		ctx    context.Context
		userID int64
	}
	tests := []struct {
		name    string
		args    args
		prepare func(f *fields)
		want    []entity.Subscription
		wantErr bool
	}{
		{
			name: "repo returned error",
			args: args{
				ctx:    context.Background(),
				userID: int64(1),
			},
			prepare: func(f *fields) {
				f.repo.EXPECT().
					GetUserSubscriptions(context.Background(), int64(1)).
					Return(nil, errors.New("error"))
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				ctx:    context.Background(),
				userID: int64(1),
			},
			prepare: func(f *fields) {
				f.repo.EXPECT().
					GetUserSubscriptions(context.Background(), int64(1)).
					Return(
						[]entity.Subscription{
							{
								ID:             1,
								UserID:         int64(1),
								Symbol:         "BTC",
								UpdateInterval: 5 * time.Minute,
							},
						},
						nil,
					)
			},
			want: []entity.Subscription{
				{
					ID:             1,
					UserID:         int64(1),
					Symbol:         "BTC",
					UpdateInterval: 5 * time.Minute,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				repo:      mock_usecase.NewMockSubscriptionRepo(ctrl),
				scheduler: mock_scheduler.NewMockScheduler(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			uc := &SubscriptionUseCase{
				repo:      f.repo,
				scheduler: f.scheduler,
			}

			got, err := uc.GetUserSubscriptions(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserSubscriptions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserSubscriptions() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubscriptionUseCase_Subscribe(t *testing.T) {
	type fields struct {
		repo      *mock_usecase.MockSubscriptionRepo
		scheduler *mock_scheduler.MockScheduler
	}
	type args struct {
		ctx context.Context
		sub entity.Subscription
	}
	tests := []struct {
		name    string
		args    args
		prepare func(f *fields)
		want    int64
		wantErr bool
	}{
		{
			name: "repo returned error",
			args: args{
				ctx: context.Background(),
				sub: entity.Subscription{
					UserID:         1,
					Symbol:         "BTC",
					UpdateInterval: 5 * time.Minute,
				},
			},
			prepare: func(f *fields) {
				f.repo.EXPECT().
					Create(context.Background(), entity.Subscription{
						UserID:         1,
						Symbol:         "BTC",
						UpdateInterval: 5 * time.Minute,
					}).
					Return(int64(0), errors.New("error"))
			},
			want:    int64(0),
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				ctx: context.Background(),
				sub: entity.Subscription{
					UserID:         1,
					Symbol:         "BTC",
					UpdateInterval: 5 * time.Minute,
				},
			},
			prepare: func(f *fields) {
				id := int64(11)
				sub := entity.Subscription{
					UserID:         1,
					Symbol:         "BTC",
					UpdateInterval: 5 * time.Minute,
				}

				create := f.repo.EXPECT().
					Create(context.Background(), sub).
					Return(id, nil)

				sub.ID = id

				addSub := f.scheduler.EXPECT().
					AddSubscription(sub).
					Return()

				gomock.InOrder(create, addSub)
			},
			want:    int64(11),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				repo:      mock_usecase.NewMockSubscriptionRepo(ctrl),
				scheduler: mock_scheduler.NewMockScheduler(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			uc := &SubscriptionUseCase{
				repo:      f.repo,
				scheduler: f.scheduler,
			}

			got, err := uc.Subscribe(tt.args.ctx, tt.args.sub)
			if (err != nil) != tt.wantErr {
				t.Errorf("Subscribe() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Subscribe() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubscriptionUseCase_Unsubscribe(t *testing.T) {
	type fields struct {
		repo      *mock_usecase.MockSubscriptionRepo
		scheduler *mock_scheduler.MockScheduler
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		args    args
		prepare func(f *fields)
		wantErr bool
	}{
		{
			name: "repo returned error",
			args: args{
				ctx: context.Background(),
				id:  int64(11),
			},
			prepare: func(f *fields) {
				f.repo.EXPECT().
					Delete(context.Background(), int64(11)).
					Return(errors.New("error"))
			},
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				ctx: context.Background(),
				id:  int64(11),
			},
			prepare: func(f *fields) {
				f.repo.EXPECT().
					Delete(context.Background(), int64(11)).
					Return(nil)

				f.scheduler.EXPECT().
					DeleteSubscription(int64(11)).
					Return()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				repo:      mock_usecase.NewMockSubscriptionRepo(ctrl),
				scheduler: mock_scheduler.NewMockScheduler(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			uc := &SubscriptionUseCase{
				repo:      f.repo,
				scheduler: f.scheduler,
			}

			if err := uc.Unsubscribe(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Unsubscribe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
