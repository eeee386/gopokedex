package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
  val []byte
}

type Cache struct {
  entries map[string]CacheEntry
	mutex *sync.RWMutex
}


func NewCache(interval time.Duration) Cache {
  var cache Cache = Cache{
		entries: make(map[string]CacheEntry),
	  mutex: &sync.RWMutex{},
	}
	return cache
}

func (c Cache) readLoop() {
	
}

func (c Cache) Add(key string, entry CacheEntry) {
	c.mutex.Lock()
	c.entries[key] = entry
	c.mutex.Unlock()
}

func (c Cache) Get(key string) (CacheEntry, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	val, ok := c.entries[key]
	return val, ok
}


