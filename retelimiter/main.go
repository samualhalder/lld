package main

import (
	"time"

	ratelimiterservice "github.com/samualhalder/lld/ratelimiter/ratelimiter_service"
	"github.com/samualhalder/lld/ratelimiter/registry"
)

func main() {
	rls := ratelimiterservice.NewRateLimiterService()
	tokenBucket1, err := registry.NewRateLimiter("token-bucket", 10, time.Microsecond*1)
	if err != nil {
		print(err.Error())
		return
	}
	tokenBucket2, err := registry.NewRateLimiter("token-bucket", 1, time.Minute*1)
	if err != nil {
		print(err.Error())
		return
	}

	rls.AddUser("1", tokenBucket1)
	rls.AddUser("2", tokenBucket2)

	println(rls.AllowRequest("1"))
	println(rls.AllowRequest("2"))

	// time.Sleep(time.Second * 2)
	println(rls.AllowRequest("1"))
	println(rls.AllowRequest("2"))
	println(rls.AllowRequest("1"))
	println(rls.AllowRequest("2"))
	println(rls.AllowRequest("1"))
	println(rls.AllowRequest("2"))
	println(rls.AllowRequest("1"))
	println(rls.AllowRequest("2"))
	println(rls.AllowRequest("1"))
	println(rls.AllowRequest("2"))

}
