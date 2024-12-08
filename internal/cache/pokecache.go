package cache

// DONE

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mu      *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		entries: make(map[string]cacheEntry),
		mu:      &sync.Mutex{},
	}
	go cache.startTimer(interval)
	return cache
}

func (c *Cache) Add(k string, v []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[k] = cacheEntry{
		val:       v,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(k string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.entries[k]
	return entry.val, ok
}

func (c *Cache) startTimer(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, entry := range c.entries {
		if entry.createdAt.Before(now.Add(-last)) {
			delete(c.entries, k)
		}
	}
}
