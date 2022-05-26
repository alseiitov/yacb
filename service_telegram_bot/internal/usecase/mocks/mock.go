// Code generated by MockGen. DO NOT EDIT.
// Source: service_telegram_bot/internal/usecase/interfaces.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"
	time "time"

	entity "github.com/alseiitov/yacb/service_telegram_bot/internal/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// GetUserInfo mocks base method.
func (m *MockUser) GetUserInfo(ctx context.Context, id int64) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserInfo", ctx, id)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserInfo indicates an expected call of GetUserInfo.
func (mr *MockUserMockRecorder) GetUserInfo(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserInfo", reflect.TypeOf((*MockUser)(nil).GetUserInfo), ctx, id)
}

// IsRegistered mocks base method.
func (m *MockUser) IsRegistered(ctx context.Context, id int64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRegistered", ctx, id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsRegistered indicates an expected call of IsRegistered.
func (mr *MockUserMockRecorder) IsRegistered(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRegistered", reflect.TypeOf((*MockUser)(nil).IsRegistered), ctx, id)
}

// Register mocks base method.
func (m *MockUser) Register(ctx context.Context, userID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockUserMockRecorder) Register(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockUser)(nil).Register), ctx, userID)
}

// Unregister mocks base method.
func (m *MockUser) Unregister(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unregister", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unregister indicates an expected call of Unregister.
func (mr *MockUserMockRecorder) Unregister(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unregister", reflect.TypeOf((*MockUser)(nil).Unregister), ctx, id)
}

// MockUserRepo is a mock of UserRepo interface.
type MockUserRepo struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepoMockRecorder
}

// MockUserRepoMockRecorder is the mock recorder for MockUserRepo.
type MockUserRepoMockRecorder struct {
	mock *MockUserRepo
}

// NewMockUserRepo creates a new mock instance.
func NewMockUserRepo(ctrl *gomock.Controller) *MockUserRepo {
	mock := &MockUserRepo{ctrl: ctrl}
	mock.recorder = &MockUserRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepo) EXPECT() *MockUserRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserRepo) Create(ctx context.Context, user entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUserRepoMockRecorder) Create(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserRepo)(nil).Create), ctx, user)
}

// Delete mocks base method.
func (m *MockUserRepo) Delete(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserRepoMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserRepo)(nil).Delete), ctx, id)
}

// GetUserInfo mocks base method.
func (m *MockUserRepo) GetUserInfo(ctx context.Context, id int64) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserInfo", ctx, id)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserInfo indicates an expected call of GetUserInfo.
func (mr *MockUserRepoMockRecorder) GetUserInfo(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserInfo", reflect.TypeOf((*MockUserRepo)(nil).GetUserInfo), ctx, id)
}

// IsRegistered mocks base method.
func (m *MockUserRepo) IsRegistered(ctx context.Context, id int64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRegistered", ctx, id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsRegistered indicates an expected call of IsRegistered.
func (mr *MockUserRepoMockRecorder) IsRegistered(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRegistered", reflect.TypeOf((*MockUserRepo)(nil).IsRegistered), ctx, id)
}

// MockSubscription is a mock of Subscription interface.
type MockSubscription struct {
	ctrl     *gomock.Controller
	recorder *MockSubscriptionMockRecorder
}

// MockSubscriptionMockRecorder is the mock recorder for MockSubscription.
type MockSubscriptionMockRecorder struct {
	mock *MockSubscription
}

// NewMockSubscription creates a new mock instance.
func NewMockSubscription(ctrl *gomock.Controller) *MockSubscription {
	mock := &MockSubscription{ctrl: ctrl}
	mock.recorder = &MockSubscriptionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubscription) EXPECT() *MockSubscriptionMockRecorder {
	return m.recorder
}

// GetAllSubscriptions mocks base method.
func (m *MockSubscription) GetAllSubscriptions(ctx context.Context) ([]entity.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSubscriptions", ctx)
	ret0, _ := ret[0].([]entity.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSubscriptions indicates an expected call of GetAllSubscriptions.
func (mr *MockSubscriptionMockRecorder) GetAllSubscriptions(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSubscriptions", reflect.TypeOf((*MockSubscription)(nil).GetAllSubscriptions), ctx)
}

// GetUserSubscriptions mocks base method.
func (m *MockSubscription) GetUserSubscriptions(ctx context.Context, userID int64) ([]entity.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserSubscriptions", ctx, userID)
	ret0, _ := ret[0].([]entity.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserSubscriptions indicates an expected call of GetUserSubscriptions.
func (mr *MockSubscriptionMockRecorder) GetUserSubscriptions(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserSubscriptions", reflect.TypeOf((*MockSubscription)(nil).GetUserSubscriptions), ctx, userID)
}

// Subscribe mocks base method.
func (m *MockSubscription) Subscribe(ctx context.Context, subscription entity.Subscription) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", ctx, subscription)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockSubscriptionMockRecorder) Subscribe(ctx, subscription interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockSubscription)(nil).Subscribe), ctx, subscription)
}

// Unsubscribe mocks base method.
func (m *MockSubscription) Unsubscribe(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unsubscribe", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unsubscribe indicates an expected call of Unsubscribe.
func (mr *MockSubscriptionMockRecorder) Unsubscribe(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockSubscription)(nil).Unsubscribe), ctx, id)
}

// MockSubscriptionRepo is a mock of SubscriptionRepo interface.
type MockSubscriptionRepo struct {
	ctrl     *gomock.Controller
	recorder *MockSubscriptionRepoMockRecorder
}

// MockSubscriptionRepoMockRecorder is the mock recorder for MockSubscriptionRepo.
type MockSubscriptionRepoMockRecorder struct {
	mock *MockSubscriptionRepo
}

// NewMockSubscriptionRepo creates a new mock instance.
func NewMockSubscriptionRepo(ctrl *gomock.Controller) *MockSubscriptionRepo {
	mock := &MockSubscriptionRepo{ctrl: ctrl}
	mock.recorder = &MockSubscriptionRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubscriptionRepo) EXPECT() *MockSubscriptionRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockSubscriptionRepo) Create(ctx context.Context, subscription entity.Subscription) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, subscription)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockSubscriptionRepoMockRecorder) Create(ctx, subscription interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSubscriptionRepo)(nil).Create), ctx, subscription)
}

// Delete mocks base method.
func (m *MockSubscriptionRepo) Delete(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockSubscriptionRepoMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSubscriptionRepo)(nil).Delete), ctx, id)
}

// GetAllSubscriptions mocks base method.
func (m *MockSubscriptionRepo) GetAllSubscriptions(ctx context.Context) ([]entity.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSubscriptions", ctx)
	ret0, _ := ret[0].([]entity.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSubscriptions indicates an expected call of GetAllSubscriptions.
func (mr *MockSubscriptionRepoMockRecorder) GetAllSubscriptions(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSubscriptions", reflect.TypeOf((*MockSubscriptionRepo)(nil).GetAllSubscriptions), ctx)
}

// GetUserSubscriptions mocks base method.
func (m *MockSubscriptionRepo) GetUserSubscriptions(ctx context.Context, userID int64) ([]entity.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserSubscriptions", ctx, userID)
	ret0, _ := ret[0].([]entity.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserSubscriptions indicates an expected call of GetUserSubscriptions.
func (mr *MockSubscriptionRepoMockRecorder) GetUserSubscriptions(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserSubscriptions", reflect.TypeOf((*MockSubscriptionRepo)(nil).GetUserSubscriptions), ctx, userID)
}

// MockRate is a mock of Rate interface.
type MockRate struct {
	ctrl     *gomock.Controller
	recorder *MockRateMockRecorder
}

// MockRateMockRecorder is the mock recorder for MockRate.
type MockRateMockRecorder struct {
	mock *MockRate
}

// NewMockRate creates a new mock instance.
func NewMockRate(ctrl *gomock.Controller) *MockRate {
	mock := &MockRate{ctrl: ctrl}
	mock.recorder = &MockRateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRate) EXPECT() *MockRateMockRecorder {
	return m.recorder
}

// GetCurrentRates mocks base method.
func (m *MockRate) GetCurrentRates(ctx context.Context) ([]entity.Rate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentRates", ctx)
	ret0, _ := ret[0].([]entity.Rate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentRates indicates an expected call of GetCurrentRates.
func (mr *MockRateMockRecorder) GetCurrentRates(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentRates", reflect.TypeOf((*MockRate)(nil).GetCurrentRates), ctx)
}

// MockCryptoCurrencyClient is a mock of CryptoCurrencyClient interface.
type MockCryptoCurrencyClient struct {
	ctrl     *gomock.Controller
	recorder *MockCryptoCurrencyClientMockRecorder
}

// MockCryptoCurrencyClientMockRecorder is the mock recorder for MockCryptoCurrencyClient.
type MockCryptoCurrencyClientMockRecorder struct {
	mock *MockCryptoCurrencyClient
}

// NewMockCryptoCurrencyClient creates a new mock instance.
func NewMockCryptoCurrencyClient(ctrl *gomock.Controller) *MockCryptoCurrencyClient {
	mock := &MockCryptoCurrencyClient{ctrl: ctrl}
	mock.recorder = &MockCryptoCurrencyClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCryptoCurrencyClient) EXPECT() *MockCryptoCurrencyClientMockRecorder {
	return m.recorder
}

// GetCurrentRate mocks base method.
func (m *MockCryptoCurrencyClient) GetCurrentRate(ctx context.Context, currency entity.Currency) (entity.Rate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentRate", ctx, currency)
	ret0, _ := ret[0].(entity.Rate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentRate indicates an expected call of GetCurrentRate.
func (mr *MockCryptoCurrencyClientMockRecorder) GetCurrentRate(ctx, currency interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentRate", reflect.TypeOf((*MockCryptoCurrencyClient)(nil).GetCurrentRate), ctx, currency)
}

// GetRateChange mocks base method.
func (m *MockCryptoCurrencyClient) GetRateChange(ctx context.Context, symbol string, period time.Duration) (entity.RateChange, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRateChange", ctx, symbol, period)
	ret0, _ := ret[0].(entity.RateChange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRateChange indicates an expected call of GetRateChange.
func (mr *MockCryptoCurrencyClientMockRecorder) GetRateChange(ctx, symbol, period interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRateChange", reflect.TypeOf((*MockCryptoCurrencyClient)(nil).GetRateChange), ctx, symbol, period)
}
