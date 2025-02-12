package pokecache

import (
	"time"
	"sync"
)

// Cache struct
type Cache struct {
	Entries map[string]CacheEntry
	Mut *sync.RWMutex
}

// Cache entries
type CacheEntry struct {
	createdAt time.Time
	val []byte
}