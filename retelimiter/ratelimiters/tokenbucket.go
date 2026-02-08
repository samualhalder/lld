package ratelimiters

import (
	"sync"
	"time"
)

type TokenBucket struct {
	limit    int
	duration time.Duration
	tokens   int
	lastfill time.Time
	mu       sync.Mutex
}

func NewTokenBucket(limit int, duration time.Duration) *TokenBucket {
	return &TokenBucket{
		limit:    limit,
		duration: duration,
		tokens:   0,
		lastfill: time.Now(),
	}
}

func (t *TokenBucket) Allow() bool {
	now := time.Now()
	t.mu.Lock()
	defer t.mu.Unlock()
	// cal time passed
	timePassed := now.Sub(t.lastfill)
	rate := float64(t.limit) / timePassed.Seconds()
	newTokens := timePassed.Seconds() * rate

	if newTokens > 0 {
		t.tokens = min(t.limit, t.tokens+int(newTokens))
		t.lastfill = now
	}
	if t.tokens >= 1 {
		t.tokens--
		return true

	}
	return false

}
