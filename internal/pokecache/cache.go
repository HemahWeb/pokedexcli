package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu       sync.RWMutex
	entries  map[string]cacheEntry
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(ttl time.Duration) *Cache {
	c := &Cache{
		entries:  make(map[string]cacheEntry),
		interval: ttl,
	}
	go func() {
		ticker := time.NewTicker(ttl)
		for {
			<-ticker.C
			c.DeleteExpired()
		}
	}()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	if time.Since(entry.createdAt) > c.interval {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) DeleteExpired() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, entry := range c.entries {
		if time.Since(entry.createdAt) > c.interval {
			delete(c.entries, key)
		}
	}
}
