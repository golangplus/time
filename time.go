package timep

import (
	"time"
)

const (
	Day  = 24 * time.Hour
	Week = 7 * Day
)

// SleepUntil sleeps until the specified time.
func SleepUntil(t time.Time) {
	now := time.Now()
	if !t.After(now) {
		// t <= now, the time has passed
		return
	}
	time.Sleep(t.Sub(now))
}
