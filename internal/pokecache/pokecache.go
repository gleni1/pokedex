package pokecache

import (
	_ "fmt"
	"sync"
	"time"
)

type Cache struct {
	mu       sync.Mutex
	store    map[string]CacheEntry
	interval time.Duration
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	return &Cache{
		store:    make(map[string]CacheEntry),
		interval: interval,
	}
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	entry := CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	cache.store[key] = entry
	cache.mu.Unlock()
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	entry, exists := cache.store[key]
	if exists {
		return entry.val, true
	}
	return nil, false
}

func (cache *Cache) ReapLoop() {
	for {
		time_interval := cache.interval
		time.Sleep(time_interval)
		curr_time := time.Now()
		for key, value := range cache.store {
			if curr_time.Sub(value.createdAt) >= time_interval {
				cache.mu.Lock()
				delete(cache.store, key)
				cache.mu.Unlock()
			}
		}
	}
}
