package usecase

import (
	"context"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/entity"
	"github.com/alseiitov/yacb/service_telegram_bot/internal/pkg/currencies"
	mock_usecase "github.com/alseiitov/yacb/service_telegram_bot/internal/usecase/mocks"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"reflect"
	"testing"
	"time"
)

func TestNewRatesUseCase(t *testing.T) {
	type args struct {
		cryptoCurrencyClient CryptoCurrencyClient
	}
	tests := []struct {
		name string
		args args
		want *RateUseCase
	}{
		{
			name: "ok",
			args: args{
				cryptoCurrencyClient: &mock_usecase.MockCryptoCurrencyClient{},
			},
			want: &RateUseCase{
				cryptoCurrencyClient: &mock_usecase.MockCryptoCurrencyClient{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRatesUseCase(tt.args.cryptoCurrencyClient); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRatesUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRateUseCase_GetCurrentRates(t *testing.T) {
	ctx, _ := context.WithCancel(context.Background())
	now := time.Now()

	type fields struct {
		cryptoCurrencyClient *mock_usecase.MockCryptoCurrencyClient
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		prepare func(f *fields)
		want    []entity.Rate
		wantErr bool
	}{
		{
			name: "empty resp",
			args: args{
				ctx: ctx,
			},
			prepare: func(f *fields) {
				currencies.List = []currencies.Currency{currencies.Bitcoin}
				f.cryptoCurrencyClient.EXPECT().
					GetCurrentRate(ctx, currencies.Currencies[currencies.Bitcoin]).
					Return(entity.Rate{}, errors.New("error"))
			},
			want:    []entity.Rate{},
			wantErr: false,
		},
		{
			name: "ok",
			args: args{
				ctx: ctx,
			},
			prepare: func(f *fields) {
				currencies.List = []currencies.Currency{currencies.Bitcoin, currencies.Ethereum}
				btc := currencies.Currencies[currencies.Bitcoin]
				eth := currencies.Currencies[currencies.Ethereum]
				f.cryptoCurrencyClient.EXPECT().
					GetCurrentRate(ctx, btc).
					Return(entity.Rate{
						Currency: btc,
						Price:    42000,
						Date:     now,
					}, nil)

				f.cryptoCurrencyClient.EXPECT().
					GetCurrentRate(ctx, eth).
					Return(entity.Rate{
						Currency: eth,
						Price:    3000,
						Date:     now,
					}, nil)
			},
			want: []entity.Rate{
				{
					Currency: currencies.Currencies[currencies.Bitcoin],
					Price:    42000,
					Date:     now,
				},
				{
					Currency: currencies.Currencies[currencies.Ethereum],
					Price:    3000,
					Date:     now,
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
				cryptoCurrencyClient: mock_usecase.NewMockCryptoCurrencyClient(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			uc := &RateUseCase{
				cryptoCurrencyClient: f.cryptoCurrencyClient,
			}

			got, err := uc.GetCurrentRates(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCurrentRates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCurrentRates() got = %v, want %v", got, tt.want)
			}
		})
	}
}
