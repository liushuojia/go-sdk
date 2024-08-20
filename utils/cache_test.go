package utils

import (
	"fmt"
	"github.com/coocood/freecache"
	"testing"
)

// go test -run TestCache
func TestCache(t *testing.T) {
	// In bytes, where 1024 * 1024 represents a single Megabyte, and 100 * 1024*1024 represents 100 Megabytes.
	cacheSize := 100 * 1024 * 1024
	cache := freecache.NewCache(cacheSize)
	key := []byte("abc")
	val := []byte("def")
	expire := 60 // expire in 60 seconds
	_ = cache.Set(key, val, expire)
	got, err := cache.Get(key)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", got)
	}
	affected := cache.Del(key)
	fmt.Println("deleted key ", affected)
	fmt.Println("entry count ", cache.EntryCount())
}
