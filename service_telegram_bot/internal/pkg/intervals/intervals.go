package intervals

import (
	"github.com/alseiitov/yacb/service_telegram_bot/internal/entity"
	"time"
)

type Interval uint8

const (
	FiveMinute Interval = iota
	TenMinute
	ThirtyMinute
	Hour
	ThreeHours
	SixHours
	TwelveHours
	Day
)

var (
	List = []Interval{FiveMinute, TenMinute, ThirtyMinute, Hour, ThreeHours, SixHours, TwelveHours, Day}

	Intervals = map[Interval]entity.Interval{
		FiveMinute: {
			Name:     "5 minutes",
			Duration: 5 * time.Minute,
		},
		TenMinute: {
			Name:     "10 minutes",
			Duration: 10 * time.Minute,
		},
		ThirtyMinute: {
			Name:     "30 minutes",
			Duration: 30 * time.Minute,
		},
		Hour: {
			Name:     "1 hour",
			Duration: 1 * time.Hour,
		},
		ThreeHours: {
			Name:     "3 hours",
			Duration: 3 * time.Hour,
		},
		SixHours: {
			Name:     "6 hours",
			Duration: 6 * time.Hour,
		},
		TwelveHours: {
			Name:     "12 hours",
			Duration: 12 * time.Hour,
		},
		Day: {
			Name:     "24 hours",
			Duration: 24 * time.Hour,
		},
	}
)

func GetByDuration(duration time.Duration) (Interval, entity.Interval) {
	for inter, ent := range Intervals {
		if ent.Duration == duration {
			return inter, ent
		}
	}
	return 0, entity.Interval{}
}
