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
	cache.readLoop(interval)
	return cache
}

func (c Cache) readLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for {
			select {
			case <- ticker.C:
				for key, entry := range(c.entries) {
					diff := time.Now().Sub(entry.createdAt)
					if diff > interval {
						delete(c.entries, key)
					}
				}
			}
		}
	}()
}

func (c Cache) Add(key string, entry CacheEntry) {
	c.mutex.Lock()
	c.entries[key] = entry
	c.mutex.Unlock()
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	val, ok := c.entries[key]
	if(!ok) {
		return nil, ok
	} else {
	  return val.val, ok
	}
}


