package timep

import (
	"context"
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

// SleepUntilWithCancel sleep until a specify time and returns nil, unless
// the context is cancelled, in which case ctx.Err() is returned.
func SleepUntilWithCancel(ctx context.Context, t time.Time) error {
	now := time.Now()
	if now.After(t) {
		// Already after, directly return
		return nil
	}
	timer := time.NewTimer(t.Sub(now))
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

// NowFunc has the signature of time.Now().
type NowFunc func() time.Time

// PresetNow Returns a NowFunc which returns the specific time.
func PresetNow(t time.Time) NowFunc {
	return func() time.Time { return t }
}
