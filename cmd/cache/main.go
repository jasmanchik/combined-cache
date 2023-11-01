package main

import (
	"cache/internal/cache"
	"fmt"
	"time"
)

func main() {
	c := cache.NewCombinedCache(2, time.Second)
	err := c.Add("key1", "val")
	if err != nil {
		fmt.Errorf("can't add value to cache %v", err)
	}

	res, ok := c.Get("key1")
	if ok {
		fmt.Println(res)
	}

	res, ok = c.Get("key2")
	if !ok {
		fmt.Println("key2 doesn't exists")
	}

	time.Sleep(2 * time.Second)

	res, ok = c.Get("key1")
	if !ok {
		fmt.Println("key1 doesn't exists")
	}

}
