package pokecache

import (
	"time"
	"sync"
)

// Cache struct
type Cache struct {
	entries map[string]CacheEntry
	mut *sync.RWMutex
	Done chan bool
	interval time.Duration
}

// Cache entries
type CacheEntry struct {
	createdAt time.Time
	val []byte
}