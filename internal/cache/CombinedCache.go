package cache

import (
	"sync"
	"time"
)

type Item struct {
	key      string
	val      interface{}
	expireAt time.Time
}

type CombinedCache struct {
	ttl         time.Duration
	maxLen      int
	mutex       sync.Mutex
	cache       map[string]Item
	historyList HistoryList
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
			<-time.After(c.ttl)
			now := time.Now()
			for key, val := range c.cache {
				if val.expireAt.After(now) {
					c.mutex.Lock()
					delete(c.cache, key)
					if h, ok := c.historyList.Find(key); ok {
						c.historyList.Remove(h)
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
		c.cache[key] = val
		c.mutex.Unlock()
		if h, ok := c.historyList.Find(key); ok {
			c.historyList.MoveToEnd(h)
		}
		return nil
	}
	c.mutex.Lock()
	if len(c.cache) >= c.maxLen {
		if oldNode, ok := c.historyList.PopFront(); ok {
			delete(c.cache, oldNode.Value)
		}
	}

	c.cache[key] = Item{
		key,
		value,
		time.Now().Add(c.ttl),
	}
	c.historyList.Append(key)

	c.mutex.Unlock()

	return nil
}

func (c *CombinedCache) Get(key string) (Item, bool) {
	val, ok := c.cache[key]
	if ok {
		if h, ok := c.historyList.Find(key); ok {
			c.historyList.MoveToEnd(h)
		}
	}
	return val, ok
}
