package tummy

import (
	"time"
)

var (
	Now = time.Now

	global = new()
)

// Enable sets the global time.Now to a custom tummy.Now function.
func Enable() {
	Now = global.Now
}

// Disable resets the global tummy.Now to the original time.Now function.
func Disable() {
	Now = time.Now
}

// AddDuration adds a duration to the global tummy.Now.
func AddDuration(d time.Duration) {
	global.AddDuration(d)
}

// AddDate adds years, months, and days to the global tummy.Now.
func AddDate(years int, months int, days int) {
	global.AddDate(years, months, days)
}

// SetTime sets the global tummy.Now to a specific time.
func SetTime(t time.Time) {
	global.SetTime(t)
}

// ////////////////////////////////////

type tummy struct {
	Start time.Time
	Time  time.Time
}

func new() *tummy {
	return &tummy{
		Start: time.Now(),
		Time:  time.Now(),
	}
}

func (t *tummy) SetTime(t2 time.Time) {
	t.Start = time.Now()
	t.Time = t2
}

func (t *tummy) Now() time.Time {
	return t.Time.Add(time.Since(t.Start))
}

func (t *tummy) AddDuration(d time.Duration) {
	t.Time = t.Time.Add(d)
}

func (t *tummy) AddDate(years, months, days int) {
	t.Time = t.Time.AddDate(years, months, days)
}
