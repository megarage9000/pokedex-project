package pokecache

import (
	"time"
)

func (c *Cache) Add(key string, val []byte) {

}

func (c *Cache) Get(key string) ([]byte, bool) {
	return nil, false
} 

// Todo
func (c *Cache) reapLoop() {

}

func NewCache(duration time.Duration) *Cache {
	return nil
}