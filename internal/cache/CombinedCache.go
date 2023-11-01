package cache

import (
	"sync"
	"time"
)

type Item struct {
	Key      string
	Val      interface{}
	LiveTill time.Time
}

type CombinedCache struct {
	TTL   time.Duration
	Cache map[string]Item
}

func NewCombinedCache(cap int, ttl time.Duration) CombinedCache {
	c := CombinedCache{
		ttl,
		make(map[string]Item, cap),
	}
	var mu sync.Mutex
	go func() {
		for {
			for key, val := range c.Cache {
				if val.LiveTill.After(time.Now()) {
					mu.Lock()
					delete(c.Cache, key)
					mu.Unlock()
				}
			}
		}
	}()

	return c
}

func (c *CombinedCache) Add(key string, value interface{}) error {
	var mu sync.Mutex
	mu.Lock()
	c.Cache[key] = Item{
		key,
		value,
		time.Now().Add(c.TTL),
	}
	mu.Unlock()
	return nil
}

func (c *CombinedCache) Get(key string) (interface{}, bool) {
	val, ok := c.Cache[key]
	return val, ok
}
