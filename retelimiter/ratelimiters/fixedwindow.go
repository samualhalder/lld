package ratelimiters

import (
	"sync"
	"time"
)

type FixedWindow struct {
	limit        int
	window       time.Duration
	requests     int
	windowStarts time.Time
	mu           sync.Mutex
}

func NewFixedWindow(limit int, window time.Duration) *FixedWindow {
	return &FixedWindow{
		limit:        limit,
		window:       window,
		windowStarts: time.Now(),
		requests:     0,
	}
}

func (f *FixedWindow) Allow() bool {
	now := time.Now()
	f.mu.Lock()
	defer f.mu.Unlock()
	timePassed := now.Sub(f.windowStarts)
	if timePassed >= f.window {
		f.windowStarts = now
		f.requests = 1
	} else {
		f.requests++
	}
	return f.requests <= f.limit
}
