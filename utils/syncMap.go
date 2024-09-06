package utils

import (
	"sync"
)

type Cache struct {
	mu   sync.RWMutex
	data map[any]any
}

func (c *Cache) make() {
	c.data = make(map[any]any)
}
func (c *Cache) Load(key any) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.data == nil {
		c.make()
	}
	value, flag := c.data[key]
	return value, flag
}
func (c *Cache) Store(key any, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.data == nil {
		c.make()
	}
	c.data[key] = value
}
func (c *Cache) Delete(key any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.data == nil {
		c.make()
	}
	delete(c.data, key)
}
func (c *Cache) LoadAndDelete(key string) (any, bool) {
	v, f := c.Load(key)
	if f {
		c.Delete(key)
	}
	return v, f
}
func (c *Cache) Range(f func(key any, value any) bool) {
	if c.data == nil || len(c.data) <= 0 {
		return
	}
	for k, v := range c.data {
		if !f(k, v) {
			break
		}
	}
}
func (c *Cache) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.data)
}
