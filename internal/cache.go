package pokecache

import (
	"log"
	"sync"
	"time"
)

type Cache struct {
	Content map[string]CacheEntry
	Mut     sync.Mutex
	ttl     time.Duration
}

type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

func NewCache(t int) Cache {
	c := Cache{
		Content: map[string]CacheEntry{},
		Mut:     sync.Mutex{},
		ttl:     time.Duration(t) * time.Second,
	}
	return c
}

func (c *Cache) Add(key string, val []byte) {
	log.Printf("Received requst to add %v to cache\n",key)
	log.Printf("Checking if %v exists in cache\n",key)
	c.Mut.Lock()
	defer c.Mut.Unlock()
	entry, ok := c.Content[key]
	if ok {
		log.Printf("%v exists in cache. Updating createdAt time\n",key)
		entry.CreatedAt = time.Now()
		return
	}
	log.Printf("%v doesnt exists in cache. Adding\n",key)
	c.Content[key] = CacheEntry{
		CreatedAt: time.Now(),
		Val:       val,
	}
	log.Printf("Cache after adding %v\n",c.Content)
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mut.Lock()
	defer c.Mut.Unlock()
	entry, ok := c.Content[key]
	if ok {
		return entry.Val, true
	}
	return []byte{}, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}
func (c *Cache) reap(now time.Time, last time.Duration) {
	c.Mut.Lock()
	defer c.Mut.Unlock()
	for k, v := range c.Content {
		if v.CreatedAt.Before(now.Add(-last)) {
			delete(c.Content, k)
		}
	}
}