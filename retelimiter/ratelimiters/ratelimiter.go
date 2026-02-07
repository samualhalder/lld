package ratelimiters

type RateLimiter interface {
	Allow() bool
}
