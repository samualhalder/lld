package main

import (
	"fmt"

	"github.com/samualhalder/lld/cache_system/storage"
)

type Cache struct {
	store storage.Storage
}

func (c *Cache) Get(s string) {
	val := c.store.Get(s)
	if val != nil {
		fmt.Println(*val)
	} else {
		fmt.Println("Not presant")
	}
}
func (c *Cache) Put(s string, val string) {
	c.store.Put(s, val)
}

func main() {
	mapStore := storage.NewMapStorage(0)
	cache := Cache{
		store: &mapStore,
	}
	cache.Put("1", "Hello") // 1
	cache.Put("2", "hi")    // 1 2
	cache.Put("3", "NICE")  // 1 2 3
	cache.Get("1")          // hello, 2 3 1
	cache.Put("4", "four")  //  3 1 4
	cache.Get("2")          // np
	cache.Get("3")          // nice 1 4 3
	cache.Put("5", "five")  // 4 3 5
	cache.Get("1")          // np
	cache.Put("5", "edit 5")
	cache.Get("4")        //four  3 5 4
	cache.Put("6", "six") // 5 4 6
	cache.Get("3")        // np
}
