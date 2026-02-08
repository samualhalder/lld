package ratelimiters

import (
	"sync"
	"time"
)

type SlidingWindow struct {
	limit  int
	window time.Duration
	logs   []time.Time
	mu     sync.Mutex
}

func NewSlidingWindow(limit int, window time.Duration) *SlidingWindow {
	return &SlidingWindow{
		limit:  limit,
		window: window,
		logs:   make([]time.Time, 0),
	}
}

func (s *SlidingWindow) Allow() bool {
	now := time.Now()
	s.mu.Lock()
	defer s.mu.Unlock()
	startTime := now.Add(-s.window)
	i := 0
	for i < len(s.logs) && s.logs[i].Before(startTime) {
		i++
	}
	s.logs = s.logs[i:]
	if len(s.logs) < s.limit {
		s.logs = append(s.logs, now)
		return true
	}
	return false
}
