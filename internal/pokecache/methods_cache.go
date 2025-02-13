package pokecache

import (
	"time"
	"sync"
)

func (c *Cache) Add(key string, val []byte) {
	c.mut.Lock()
	c.entries[key] = CacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.mut.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mut.RLock()
	defer c.mut.RUnlock()

	result, ok := c.entries[key]
	if !ok {
		return nil, ok
	}
	return result.val, ok
} 


func (c *Cache) reapLoop() {

	// Ticker to regularly poll the cache loop
	ticker := time.NewTicker(c.interval)

	// Don't forget to stop ticker when loop is done
	defer ticker.Stop()

	// For polling
	for {
		select {
			// If cache / program is complete
		case <- c.Done:
			break
			// If ticker 
		case <-ticker.C:
			c.removeOldEntries()
		}
	}
}

func (c *Cache) removeOldEntries() {

	var entriesToDelete []string

	// Collect all old keys to delete
	c.mut.RLock()

	for key := range c.entries {
		entry := c.entries[key]

		// If cacheEntry is old, add to slice to delete
		if time.Since(entry.createdAt) > c.interval {
			entriesToDelete = append(entriesToDelete, key)
		}
	}
	c.mut.RUnlock()

	// Performing key deletion
	c.mut.Lock()
	for _, entry := range entriesToDelete {
		delete(c.entries, entry)
	}
	c.mut.Unlock()
}

func NewCache(duration time.Duration) Cache {
	cache := Cache{
		entries : make(map[string]CacheEntry),
		mut : &sync.RWMutex{},
		Done : make(chan bool),
		interval: duration,
	}
	go cache.reapLoop()
	return cache
}