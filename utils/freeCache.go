package utils

import (
	"github.com/coocood/freecache"
)

const expire = 10 * 24 * 60 * 60

type FreeCache struct {
	cache *freecache.Cache
}

func (c *FreeCache) make() {
	if c.cache == nil {
		cacheSize := 100 * 1024 * 1024
		c.cache = freecache.NewCache(cacheSize)
	}
}
func (c *FreeCache) Load(key string) (any, bool) {
	c.make()
	v, err := c.cache.Get([]byte(key))
	if err != nil {
		return nil, false
	}

	var value any
	if err := GobByte2Any(v, &value); err != nil {
		return nil, false
	}

	return value, true
}
func (c *FreeCache) Store(key string, value any) {
	c.make()

	v, _ := GobAny2Byte(value)
	_ = c.cache.Set([]byte(key), v, expire)
}
func (c *FreeCache) Delete(key string) {
	c.make()
	_ = c.cache.Del([]byte(key))
}
func (c *FreeCache) LoadAndDelete(key string) (any, bool) {
	c.make()
	v, f := c.Load(key)
	if f {
		c.Delete(key)
	}
	return v, f
}
func (c *FreeCache) Range(f func(key string, value any) bool) {
	c.make()
	iterator := c.cache.NewIterator()
	for {
		entry := iterator.Next()
		if entry == nil {
			break
		}

		var v any
		if err := GobByte2Any(entry.Value, &v); err != nil {
			continue
		}

		if flag := f(string(entry.Key), v); !flag {
			break
		}
	}
}
func (c *FreeCache) Len() int {
	c.make()
	return int(c.cache.EntryCount())
}
