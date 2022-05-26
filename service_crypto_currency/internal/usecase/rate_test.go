package usecase

import (
	"bou.ke/monkey"
	"context"
	"github.com/alseiitov/yacb/service_crypto_currency/internal/entity"
	mock_usecase "github.com/alseiitov/yacb/service_crypto_currency/internal/usecase/mocks"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
	"reflect"
	"testing"
	"time"
)

func TestNewRatesUseCase(t *testing.T) {
	type args struct {
		rateRepo           RateRepo
		cryptoCurrencyRepo CryptoCurrencyRepo
		rateClient         RateClient
	}
	tests := []struct {
		name string
		args args
		want *RateUseCase
	}{
		{
			name: "ok",
			args: args{
				rateRepo:           &mock_usecase.MockRateRepo{},
				cryptoCurrencyRepo: &mock_usecase.MockCryptoCurrencyRepo{},
				rateClient:         &mock_usecase.MockRateClient{},
			},
			want: &RateUseCase{
				rateRepo:           &mock_usecase.MockRateRepo{},
				cryptoCurrencyRepo: &mock_usecase.MockCryptoCurrencyRepo{},
				rateClient:         &mock_usecase.MockRateClient{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRatesUseCase(tt.args.rateRepo, tt.args.cryptoCurrencyRepo, tt.args.rateClient); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRatesUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRateUseCase_Create(t *testing.T) {
	type fields struct {
		rateRepo           *mock_usecase.MockRateRepo
		cryptoCurrencyRepo *mock_usecase.MockCryptoCurrencyRepo
		rateClient         *mock_usecase.MockRateClient
	}
	type args struct {
		ctx  context.Context
		rate entity.Rate
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
				ctx:  context.Background(),
				rate: entity.Rate{},
			},
			prepare: func(f *fields) {
				gomock.InOrder(
					f.rateRepo.EXPECT().
						Create(context.Background(), entity.Rate{}).
						Return(errors.New("error")),
				)
			},
			wantErr: true,
		},
		{
			name: "repo returned error",
			args: args{
				ctx:  context.Background(),
				rate: entity.Rate{},
			},
			prepare: func(f *fields) {
				gomock.InOrder(
					f.rateRepo.EXPECT().
						Create(context.Background(), entity.Rate{}).
						Return(nil),
				)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				rateRepo:           mock_usecase.NewMockRateRepo(ctrl),
				cryptoCurrencyRepo: mock_usecase.NewMockCryptoCurrencyRepo(ctrl),
				rateClient:         mock_usecase.NewMockRateClient(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			uc := &RateUseCase{
				rateRepo:           f.rateRepo,
				cryptoCurrencyRepo: f.cryptoCurrencyRepo,
				rateClient:         f.rateClient,
			}
			if err := uc.Create(tt.args.ctx, tt.args.rate); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRateUseCase_GetCurrentRate(t *testing.T) {
	now := time.Now()

	type fields struct {
		rateRepo           *mock_usecase.MockRateRepo
		cryptoCurrencyRepo *mock_usecase.MockCryptoCurrencyRepo
		rateClient         *mock_usecase.MockRateClient
	}
	type args struct {
		ctx    context.Context
		symbol string
	}
	tests := []struct {
		name    string
		args    args
		prepare func(f *fields)
		want    entity.Rate
		wantErr bool
	}{
		{
			name: "client returned error",
			args: args{
				ctx:    context.Background(),
				symbol: "BTC",
			},
			prepare: func(f *fields) {
				gomock.InOrder(
					f.rateClient.EXPECT().
						GetCurrentRate(context.Background(), "BTC").
						Return(entity.Rate{}, errors.New("client error")),
				)
			},
			want:    entity.Rate{},
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				ctx:    context.Background(),
				symbol: "BTC",
			},
			prepare: func(f *fields) {
				gomock.InOrder(
					f.rateClient.EXPECT().
						GetCurrentRate(context.Background(), "BTC").
						Return(
							entity.Rate{
								CryptoCurrencyID:   1,
								CryptoCurrencyName: "Bitcoin",
								Price:              42500,
								Date:               &now,
							},
							nil,
						),
				)
			},
			want: entity.Rate{
				CryptoCurrencyID:   1,
				CryptoCurrencyName: "Bitcoin",
				Price:              42500,
				Date:               &now,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				rateRepo:           mock_usecase.NewMockRateRepo(ctrl),
				cryptoCurrencyRepo: mock_usecase.NewMockCryptoCurrencyRepo(ctrl),
				rateClient:         mock_usecase.NewMockRateClient(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			uc := &RateUseCase{
				rateRepo:           f.rateRepo,
				cryptoCurrencyRepo: f.cryptoCurrencyRepo,
				rateClient:         f.rateClient,
			}

			got, err := uc.GetCurrentRate(tt.args.ctx, tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCurrentRate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCurrentRate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRateUseCase_GetRateChange(t *testing.T) {
	now := time.Now()
	currTime := now.UTC().Truncate(5 * time.Minute)
	prevTime := currTime.Add(-5 * time.Minute)

	type fields struct {
		rateRepo           *mock_usecase.MockRateRepo
		cryptoCurrencyRepo *mock_usecase.MockCryptoCurrencyRepo
		rateClient         *mock_usecase.MockRateClient
	}
	type args struct {
		ctx    context.Context
		symbol string
		period time.Duration
	}
	tests := []struct {
		name    string
		args    args
		prepare func(f *fields)
		want    entity.RateChange
		wantErr bool
	}{
		{
			name: "failed to get current rate",
			args: args{
				ctx:    context.Background(),
				symbol: "BTC",
				period: 5 * time.Minute,
			},
			prepare: func(f *fields) {
				gomock.InOrder(
					f.rateRepo.EXPECT().
						GetRateByDate(context.Background(), "BTC", now.UTC().Truncate(5*time.Minute)).
						Return(entity.Rate{}, errors.New("error")),
				)
			},
			want:    entity.RateChange{},
			wantErr: true,
		},
		{
			name: "failed to get previous rate",
			args: args{
				ctx:    context.Background(),
				symbol: "BTC",
				period: 5 * time.Minute,
			},
			prepare: func(f *fields) {
				currTime := now.UTC().Truncate(5 * time.Minute)
				prevTime := currTime.Add(-5 * time.Minute)
				gomock.InOrder(
					f.rateRepo.EXPECT().
						GetRateByDate(context.Background(), "BTC", currTime).
						Return(entity.Rate{}, nil),
					f.rateRepo.EXPECT().
						GetRateByDate(context.Background(), "BTC", prevTime).
						Return(entity.Rate{}, errors.New("error")),
				)
			},
			want:    entity.RateChange{},
			wantErr: true,
		},
		{
			name: "failed to get previous rate",
			args: args{
				ctx:    context.Background(),
				symbol: "BTC",
				period: 5 * time.Minute,
			},
			prepare: func(f *fields) {
				gomock.InOrder(
					f.rateRepo.EXPECT().
						GetRateByDate(context.Background(), "BTC", currTime).
						Return(
							entity.Rate{
								CryptoCurrencyID:   1,
								CryptoCurrencyName: "Bitcoin",
								Price:              38000,
								Date:               &currTime,
							},
							nil,
						),
					f.rateRepo.EXPECT().
						GetRateByDate(context.Background(), "BTC", prevTime).
						Return(
							entity.Rate{
								CryptoCurrencyID:   1,
								CryptoCurrencyName: "Bitcoin",
								Price:              40000,
								Date:               &prevTime,
							},
							nil,
						),
				)
			},
			want: entity.RateChange{
				CryptoCurrencyName:    "Bitcoin",
				PrevPrice:             40000,
				CurrPrice:             38000,
				PriceChange:           -2000,
				PriceChangePercentage: -5,
				PrevDate:              &prevTime,
				CurrDate:              &currTime,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				rateRepo:           mock_usecase.NewMockRateRepo(ctrl),
				cryptoCurrencyRepo: mock_usecase.NewMockCryptoCurrencyRepo(ctrl),
				rateClient:         mock_usecase.NewMockRateClient(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			uc := &RateUseCase{
				rateRepo:           f.rateRepo,
				cryptoCurrencyRepo: f.cryptoCurrencyRepo,
				rateClient:         f.rateClient,
			}

			got, err := uc.GetRateChange(tt.args.ctx, tt.args.symbol, tt.args.period)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRateChange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRateChange() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRateUseCase_RunRatesParser(t *testing.T) {
	c := cron.New()
	monkey.Patch(cron.New, func(...cron.Option) *cron.Cron { return c })
	defer monkey.UnpatchAll()

	type fields struct {
		rateRepo           *mock_usecase.MockRateRepo
		cryptoCurrencyRepo *mock_usecase.MockCryptoCurrencyRepo
		rateClient         *mock_usecase.MockRateClient
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		after   func()
		wantErr bool
	}{
		{
			name: "get crypto currencies returned error",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.cryptoCurrencyRepo.EXPECT().
						GetCryptoCurrencies(context.Background()).
						Return(nil, errors.New("error")),
				)
			},
			wantErr: true,
		},
		{
			name: "ok",
			prepare: func(f *fields) {
				now := time.Now().Truncate(time.Minute).UTC()
				gomock.InOrder(
					f.cryptoCurrencyRepo.EXPECT().
						GetCryptoCurrencies(context.Background()).
						Return(
							[]entity.CryptoCurrency{
								{
									ID:     1,
									Symbol: "BTC",
									Name:   "Bitcoin",
								},
							},
							nil,
						),
					f.rateClient.EXPECT().
						GetCurrentRate(context.Background(), "BTC").
						Return(entity.Rate{
							CryptoCurrencyID:   1,
							CryptoCurrencyName: "Bitcoin",
							Price:              40000,
							Date:               &now,
						}, nil),

					f.rateRepo.EXPECT().
						Create(context.Background(), entity.Rate{
							CryptoCurrencyID:   1,
							CryptoCurrencyName: "Bitcoin",
							Price:              40000,
							Date:               &now,
						}).
						Return(nil),
				)
			},
			after: func() {
				for _, entry := range c.Entries() {
					entry.Job.Run()
				}
				time.Sleep(1 * time.Second)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				rateRepo:           mock_usecase.NewMockRateRepo(ctrl),
				cryptoCurrencyRepo: mock_usecase.NewMockCryptoCurrencyRepo(ctrl),
				rateClient:         mock_usecase.NewMockRateClient(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			uc := &RateUseCase{
				rateRepo:           f.rateRepo,
				cryptoCurrencyRepo: f.cryptoCurrencyRepo,
				rateClient:         f.rateClient,
			}

			if err := uc.RunRatesParser(); (err != nil) != tt.wantErr {
				t.Errorf("RunRatesParser() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.after != nil {
				tt.after()
			}
		})
	}
}
