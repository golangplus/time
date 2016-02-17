package timep

import (
	"time"
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
