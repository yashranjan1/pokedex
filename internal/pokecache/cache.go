package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	data  map[string]cacheEntry
	mutex *sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		data:  make(map[string]cacheEntry),
		mutex: &sync.Mutex{},
	}
	go cache.reapLoop(interval)
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	data, exists := c.data[key]
	if !exists {
		return nil, false
	}
	return data.val, true
}

func (c *Cache) reapLoop(t time.Duration) {
	tick := time.NewTicker(t)
	for {
		// block until you see a tick
		<-tick.C

		c.mutex.Lock()

		for key, entry := range c.data {
			lapsedTime := time.Since(entry.createdAt)
			if lapsedTime >= t {
				delete(c.data, key)
			}
		}

		c.mutex.Unlock()
	}
}
