package ratelimiters

import (
	"time"
)

type TokenBucket struct {
	limit    int
	duration time.Duration
	tokens   int
	lastfill time.Time
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
	return true
}
