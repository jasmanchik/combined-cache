package cache

import (
	"sync"
	"time"
)

type Item struct {
	Key       string
	Val       interface{}
	CreatedAt time.Time
}

type CombinedCache struct {
	TTL    time.Duration
	MaxLen int
	mutex  sync.Mutex
	Cache  map[string]Item
	HistoryList
}

func NewCombinedCache(cap int, ttl time.Duration) *CombinedCache {
	c := CombinedCache{
		ttl,
		cap,
		sync.Mutex{},
		make(map[string]Item, cap),
		HistoryList{},
	}

	go func() {
		for {
			<-time.After(c.TTL)
			now := time.Now()
			for key, val := range c.Cache {
				if now.Sub(val.CreatedAt) > c.TTL {
					c.mutex.Lock()
					delete(c.Cache, key)
					if h, ok := c.Find(key); ok {
						c.Remove(h)
					}
					c.mutex.Unlock()
				}
			}
		}
	}()

	return &c
}

func (c *CombinedCache) Add(key string, value interface{}) error {

	val, ok := c.Get(key)
	if ok {
		c.mutex.Lock()
		c.Cache[key] = val
		c.mutex.Unlock()
		if h, ok := c.Find(key); ok {
			c.HistoryList.MoveToEnd(h)
		}
		return nil
	}
	c.mutex.Lock()
	if len(c.Cache) >= c.MaxLen {
		if oldNode, ok := c.HistoryList.PopFront(); ok {
			delete(c.Cache, oldNode.Value)
		}
	}

	c.Cache[key] = Item{
		key,
		value,
		time.Now().Add(c.TTL),
	}
	c.HistoryList.Append(key)

	c.mutex.Unlock()

	return nil
}

func (c *CombinedCache) Get(key string) (Item, bool) {
	val, ok := c.Cache[key]
	if ok {
		if h, ok := c.HistoryList.Find(key); ok {
			c.HistoryList.MoveToEnd(h)
		}
	}
	return val, ok
}
