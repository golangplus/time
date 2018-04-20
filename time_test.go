package timep

import (
	"context"
	"testing"
	"time"

	"github.com/golangplus/testing/assert"
)

func TestSleepUntil(t *testing.T) {
	tm := time.Now().Add(time.Second)
	SleepUntil(tm)
	assert.ValueShould(t, "now", time.Now(), tm.Before, "SleepUntil returns earlier than expected")
}

func TestSleepUntil_SleepBefore(t *testing.T) {
	start := time.Now()
	SleepUntil(start)
	assert.ValueShould(t, "now", time.Now(), start.Add(time.Millisecond).After, "SleepUntil returns later than expected")
}

func TestSleepUntilWithCancel(t *testing.T) {
	ctx := context.Background()

	tm := time.Now().Add(time.Second)
	SleepUntilWithCancel(ctx, tm)
	assert.ValueShould(t, "now", time.Now(), tm.Before, "SleepUntil returns earlier than expected")
}

func TestSleepUntilWithCancel_Timeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	tm := time.Now().Add(time.Minute)
	SleepUntilWithCancel(ctx, tm)
	assert.ValueShould(t, "now", time.Now(), tm.After, "SleepUntilWithCancel should returns earlier than expected")
}

func TestPresetNow(t *testing.T) {
	tm := time.Now().Add(Day)
	now := PresetNow(tm)
	assert.Equal(t, "now", now(), tm)
}
