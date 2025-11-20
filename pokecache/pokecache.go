package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createAt time.Time
	val      []byte
}

// mutex for cache map
type Cache struct {
	m  map[string]CacheEntry
	mu *sync.Mutex
}

// function to create cache
func NewCache(interval time.Duration) Cache {

	c := Cache{
		m:  map[string]CacheEntry{},
		mu: &sync.Mutex{},
	}

	go c.ReapLoop(interval)

	return c
}

// methods for cache map
func (cm *Cache) Add(key string, val []byte) {
	(*cm).mu.Lock()
	defer (*cm).mu.Unlock()
	(*cm).m[key] = CacheEntry{
		createAt: time.Now(),
		val:      val,
	}
}

func (cm *Cache) Get(key string) ([]byte, bool) {
	(*cm).mu.Lock()
	defer (*cm).mu.Unlock()
	entry, ok := (*cm).m[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (cm *Cache) ReapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for {
		<-ticker.C
		cm.Reap(interval)
	}

}

func (cm *Cache) Reap(interval time.Duration) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	// current time
	now := time.Now()

	for k, v := range cm.m {
		if now.Sub(v.createAt) > interval { // now.Sub(v.createAt) =  v.createAT - now in duration
			delete((*cm).m, k)
		}
	}

}
