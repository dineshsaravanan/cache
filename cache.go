//inspired by Cache from SketchGround
//https://bitbucket.org/jzs/sketchground/src/4defb0a2ea64ed226680515efae4c5f8df5827a9/cache.go?at=default

package cache

import (
	"sync"
	"time"
)

type Content struct {
	Expire	time.Time
	Content []byte
}

type Cache struct {
	cache 			map[string]*Entry
	mutex 			*sync.RWMutex
	expiryLength	time.Duration
}

func (c *Cache) get(key string) []byte {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	if c.cache[key] == nil {
		return nil
	}
	if c.cache[key].Expire.Before(time.Now()) {
		return nil
	}
	return c.cache[key].Content
}

func (c *Cache) put(key string, data []byte) {
	c.mutex.Lock()
	c.cache[key] = &Entry{
		Content: data,
		Expire: time.Now().Add(dur),
	}
	c.mutex.Unlock()
	return
}

func (c *Cache) clear(key string) {
	c.mutex.Lock()
	c.cache[key] = &Entry{
		Content: nil,
		Expire: time.Now(),
	}
	c.mutex.Unlock()
	return
}


//Application Level Static Cache & expires every 4 hours by default
var StaticCache = newStaticCache(4 * time.Hour)
func newStaticCache(dur time.Duration) {
	return &Cache{
		cache: map[string]*Entry{},
		mutex: new(sync.RWMutex),
		expiryLength: dur,
	}
}
