package main

import (
	"cache/internal/cache"
	"fmt"
	"time"
)

func main() {
	c := cache.NewCombinedCache(3, 1*time.Second) // Почему нельзя передавать "time.Second"?
	c.Add("key1", "val1")
	c.Add("key2", "val2")

	res, ok := c.Get("key1")
	fmt.Println("key1:", "Exists:", ok, "Value:", res)

	res, ok = c.Get("key2")
	fmt.Println("key2:", "Exists:", ok, "Value:", res)

	time.Sleep(2 * time.Second)

	c.Add("key3", "val3")

	res, ok = c.Get("key1")
	fmt.Println("key1:", "Exists:", ok, "Value:", res)

	res, ok = c.Get("key2")
	fmt.Println("key2:", "Exists:", ok, "Value:", res)

	res, ok = c.Get("key3")
	fmt.Println("key3:", "Exists:", ok, "Value:", res)
}
