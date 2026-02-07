package main

import (
	"time"

	ratelimiterservice "github.com/samualhalder/lld/ratelimiter/ratelimiter_service"
	"github.com/samualhalder/lld/ratelimiter/registry"
)

func main() {
	rls := ratelimiterservice.NewRateLimiterService()
	tokenBucket, err := registry.NewRateLimiter("token-bucket", 10, time.Second*1)
	if err != nil {
		print(err.Error())
		return
	}

	rls.AddUser("1", tokenBucket)
	print(rls.AllowRequest("1"))
	print(rls.AllowRequest("2"))

}
