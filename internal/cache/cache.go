package cache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu    *sync.RWMutex
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(expiration time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mu:    &sync.RWMutex{},
	}

	go c.reapLoop(expiration)
	return c
}

func (c *Cache) Add(url string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	newEntry := cacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
	c.cache[url] = newEntry
}

func (c *Cache) Get(url string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, exists := c.cache[url]
	return entry.val, exists
}

func (c *Cache) reapLoop(expiration time.Duration) {
	ticker := time.NewTicker(expiration)
	for range ticker.C {
		c.reap(time.Now().UTC(), expiration)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cache, k)
		}
	}
}
