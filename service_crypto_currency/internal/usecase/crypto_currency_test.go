package usecase

import (
	"context"
	"github.com/alseiitov/yacb/service_crypto_currency/internal/entity"
	mock_usecase "github.com/alseiitov/yacb/service_crypto_currency/internal/usecase/mocks"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"reflect"
	"testing"
)

func TestNewCryptoCurrencyUseCase(t *testing.T) {
	type args struct {
		r CryptoCurrencyRepo
	}
	tests := []struct {
		name string
		args args
		want *CryptoCurrencyUseCase
	}{
		{
			name: "ok",
			args: args{
				r: &mock_usecase.MockCryptoCurrencyRepo{},
			},
			want: &CryptoCurrencyUseCase{
				&mock_usecase.MockCryptoCurrencyRepo{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCryptoCurrencyUseCase(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCryptoCurrencyUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCryptoCurrencyUseCase_GetCryptoCurrencies(t *testing.T) {
	type fields struct {
		repo *mock_usecase.MockCryptoCurrencyRepo
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		prepare func(f *fields)
		want    []entity.CryptoCurrency
		wantErr bool
	}{
		{
			name: "repo returned error",
			args: args{context.Background()},
			prepare: func(f *fields) {
				gomock.InOrder(
					f.repo.EXPECT().
						GetCryptoCurrencies(context.Background()).
						Return(nil, errors.New("error")),
				)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ok",
			args: args{context.Background()},
			prepare: func(f *fields) {
				gomock.InOrder(
					f.repo.EXPECT().
						GetCryptoCurrencies(context.Background()).
						Return(
							[]entity.CryptoCurrency{
								{
									ID:     1,
									Symbol: "BTC",
									Name:   "Bitcoin",
								},
								{
									ID:     2,
									Symbol: "ETH",
									Name:   "Ethereum",
								},
							},
							nil,
						),
				)
			},
			want: []entity.CryptoCurrency{
				{
					ID:     1,
					Symbol: "BTC",
					Name:   "Bitcoin",
				},
				{
					ID:     2,
					Symbol: "ETH",
					Name:   "Ethereum",
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
				repo: mock_usecase.NewMockCryptoCurrencyRepo(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			uc := &CryptoCurrencyUseCase{
				repo: f.repo,
			}
			got, err := uc.GetCryptoCurrencies(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCryptoCurrencies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCryptoCurrencies() got = %v, want %v", got, tt.want)
			}
		})
	}
}
