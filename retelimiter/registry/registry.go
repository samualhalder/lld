package registry

import (
	"fmt"
	"time"

	"github.com/samualhalder/lld/ratelimiter/ratelimiters"
)

func NewRateLimiter(name string, limit int, duration time.Duration) (ratelimiters.RateLimiter, error) {
	switch name {
	case "token-bucket":
		return ratelimiters.NewTokenBucket(limit, duration), nil
	default:
		return nil, fmt.Errorf("no such limiter")
	}
}
