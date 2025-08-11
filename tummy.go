package tummy

import (
	"sync"
	"time"
)

var (
	Now = time.Now

	global = new()
)

// Enable sets the global time.Now to a custom tummy.Now function.
func Enable() {
	global.Enable()
}

// Disable resets the global tummy.Now to the original time.Now function.
func Disable() {
	global.Disable()
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

// Pause pauses the global tummy.Now, freezing the time.
func Pause() {
	global.Pause()
}

// Resume resumes the global tummy.Now, unfreezing the time.
func Resume() {
	global.Resume()
}

// IsPaused checks if the global tummy.Now is currently paused.
func IsPaused() bool {
	return global.IsPaused()
}

// ////////////////////////////////////

type tummy struct {
	start time.Time
	time  time.Time

	paused     bool
	pausedTime time.Time

	m sync.RWMutex
}

func new() *tummy {
	return &tummy{
		start: time.Now(),
		time:  time.Now(),
	}
}

func (t *tummy) Enable() {
	t.m.Lock()
	defer t.m.Unlock()

	Now = t.Now
}

func (t *tummy) Disable() {
	t.m.Lock()
	defer t.m.Unlock()

	Now = time.Now
}

func (t *tummy) SetTime(t2 time.Time) {
	t.m.Lock()
	defer t.m.Unlock()

	t.start = time.Now()
	t.time = t2

	if t.paused {
		t.pausedTime = t.start
	}
}

func (t *tummy) Now() time.Time {
	t.m.RLock()
	defer t.m.RUnlock()

	if t.paused {
		return t.time.Add(t.pausedTime.Sub(t.start))
	}

	return t.time.Add(time.Since(t.start))
}

func (t *tummy) AddDuration(d time.Duration) {
	t.m.Lock()
	defer t.m.Unlock()

	t.time = t.time.Add(d)
}

func (t *tummy) AddDate(years, months, days int) {
	t.m.Lock()
	defer t.m.Unlock()
	t.time = t.time.AddDate(years, months, days)
}

func (t *tummy) Pause() {
	t.m.Lock()
	defer t.m.Unlock()

	if t.paused {
		return
	}

	t.pausedTime = time.Now()
	t.paused = true
}

func (t *tummy) Resume() {
	t.m.Lock()
	defer t.m.Unlock()

	if !t.paused {
		return
	}

	t.start = t.start.Add(time.Since(t.pausedTime))
	t.paused = false
}

func (t *tummy) IsPaused() bool {
	t.m.RLock()
	defer t.m.RUnlock()

	return t.paused
}
