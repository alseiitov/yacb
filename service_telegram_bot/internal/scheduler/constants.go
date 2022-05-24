package scheduler

import "github.com/alseiitov/yacb/service_telegram_bot/internal/pkg/intervals"

var (
	сronSpecs = map[intervals.Interval]string{
		intervals.FiveMinute:   "*/5 * * * *",
		intervals.TenMinute:    "*/10 * * * *",
		intervals.ThirtyMinute: "*/30 * * * *",
		intervals.Hour:         "0 * * * *",
		intervals.ThreeHours:   "0 */2 * * *",
		intervals.SixHours:     "0 */6 * * *",
		intervals.TwelveHours:  "0 */12 * * *",
		intervals.Day:          "0 0 * * *",
	}
)
