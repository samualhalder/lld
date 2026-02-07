package ratelimiterservice

import (
	"github.com/samualhalder/lld/ratelimiter/ratelimiters"
)

type RateLimiterService struct {
	mp map[string]ratelimiters.RateLimiter
}

func NewRateLimiterService() *RateLimiterService {
	return &RateLimiterService{
		mp: make(map[string]ratelimiters.RateLimiter),
	}
}

func (r *RateLimiterService) AddUser(userId string, limiter ratelimiters.RateLimiter) {
	r.mp[userId] = limiter
}

func (r *RateLimiterService) AllowRequest(userId string) bool {
	rl, ok := r.mp[userId]
	if !ok {
		return false
	}
	return rl.Allow()
}
