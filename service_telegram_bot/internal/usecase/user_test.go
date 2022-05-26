package usecase

import (
	"bou.ke/monkey"
	"context"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/entity"
	mock_usecase "github.com/alseiitov/yacb/service_telegram_bot/internal/usecase/mocks"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/sethvargo/go-password/password"
	"reflect"
	"testing"
)

func TestNewUsersUseCase(t *testing.T) {
	type args struct {
		repo *mock_usecase.MockUserRepo
	}
	tests := []struct {
		name string
		args args
		want *UserUseCase
	}{
		{
			name: "ok",
			args: args{
				repo: &mock_usecase.MockUserRepo{},
			},
			want: &UserUseCase{
				repo: &mock_usecase.MockUserRepo{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUsersUseCase(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUsersUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserUseCase_GetUserInfo(t *testing.T) {
	type fields struct {
		repo *mock_usecase.MockUserRepo
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		args    args
		prepare func(f *fields)
		want    entity.User
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
					GetUserInfo(context.Background(), int64(11)).
					Return(entity.User{}, errors.New("error"))
			},
			want:    entity.User{},
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
					GetUserInfo(context.Background(), int64(11)).
					Return(
						entity.User{
							ID:       int64(11),
							Password: "pass",
						},
						nil,
					)
			},
			want: entity.User{
				ID:       int64(11),
				Password: "pass",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				repo: mock_usecase.NewMockUserRepo(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			uc := &UserUseCase{
				repo: f.repo,
			}
			got, err := uc.GetUserInfo(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserUseCase_IsRegistered(t *testing.T) {
	type fields struct {
		repo *mock_usecase.MockUserRepo
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		args    args
		prepare func(f *fields)
		want    bool
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
					IsRegistered(context.Background(), int64(11)).
					Return(false, errors.New("error"))
			},
			want:    false,
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
					IsRegistered(context.Background(), int64(11)).
					Return(true, nil)
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				repo: mock_usecase.NewMockUserRepo(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			uc := &UserUseCase{
				repo: f.repo,
			}
			got, err := uc.IsRegistered(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsRegistered() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsRegistered() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserUseCase_Register(t *testing.T) {
	type fields struct {
		repo *mock_usecase.MockUserRepo
	}
	type args struct {
		ctx    context.Context
		userID int64
	}
	type passGen struct {
		pass  string
		error error
	}
	tests := []struct {
		name    string
		args    args
		passGen passGen
		prepare func(f *fields)
		wantErr bool
	}{
		{
			name: "repo returned error",
			args: args{
				ctx:    context.Background(),
				userID: int64(11),
			},
			passGen: passGen{
				pass:  "pass",
				error: nil,
			},
			prepare: func(f *fields) {
				f.repo.EXPECT().
					Create(context.Background(), entity.User{
						ID:       int64(11),
						Password: "pass",
					}).
					Return(errors.New("error"))
			},
			wantErr: true,
		},
		{
			name: "password generate error",
			args: args{
				ctx:    context.Background(),
				userID: int64(11),
			},
			passGen: passGen{
				pass:  "",
				error: errors.New("error"),
			},
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				ctx:    context.Background(),
				userID: int64(11),
			},
			passGen: passGen{
				pass:  "pass",
				error: nil,
			},
			prepare: func(f *fields) {
				f.repo.EXPECT().
					Create(context.Background(), entity.User{
						ID:       int64(11),
						Password: "pass",
					}).
					Return(nil)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				repo: mock_usecase.NewMockUserRepo(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			uc := &UserUseCase{
				repo: f.repo,
			}

			monkey.Patch(
				password.Generate,
				func(length, numDigits, numSymbols int, noUpper, allowRepeat bool) (string, error) {
					return tt.passGen.pass, tt.passGen.error
				},
			)
			defer monkey.UnpatchAll()

			if err := uc.Register(tt.args.ctx, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserUseCase_Unregister(t *testing.T) {
	type fields struct {
		repo *mock_usecase.MockUserRepo
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
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				repo: mock_usecase.NewMockUserRepo(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			uc := &UserUseCase{
				repo: f.repo,
			}

			if err := uc.Unregister(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Unregister() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
