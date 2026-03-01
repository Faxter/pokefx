package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	cache    map[string]cacheEntry
	lock     sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{cache: make(map[string]cacheEntry), lock: sync.Mutex{}, interval: interval}
	go cache.reapLoop()
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.lock.Lock()
	c.cache[key] = cacheEntry{createdAt: time.Now(), val: val}
	c.lock.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.lock.Lock()
	entry, ok := c.cache[key]
	c.lock.Unlock()
	if !ok {
		return nil, ok
	} else {
		fmt.Println("fetching from cache:", key)
		return entry.val, ok
	}
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for range ticker.C {
		c.lock.Lock()
		for key, entry := range c.cache {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.cache, key)
				fmt.Println("removed from cache:", key)
			}
		}
		c.lock.Unlock()
	}
}
