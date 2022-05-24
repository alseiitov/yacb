package entity

import (
	"github.com/alseiitov/yacb/service_crypto_currency/proto/pb"
	"google.golang.org/protobuf/types/known/durationpb"
	"time"
)

type (
	Subscription struct {
		ID             int64
		UserID         int64
		Symbol         string
		UpdateInterval time.Duration
	}
)

func UserSubscriptionListToGrpc(subscriptions []Subscription) *pb.GetUserSubscriptionsResponse {

	subscriptionsResp := make([]*pb.Subscription, 0, len(subscriptions))

	for _, subscription := range subscriptions {
		s := &pb.Subscription{
			ID:                   subscription.ID,
			UserID:               subscription.UserID,
			CryptoCurrencySymbol: subscription.Symbol,
			Interval:             durationpb.New(subscription.UpdateInterval),
		}

		subscriptionsResp = append(subscriptionsResp, s)
	}

	return &pb.GetUserSubscriptionsResponse{Subscriptions: subscriptionsResp}
}
