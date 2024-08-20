package utils

import (
	"github.com/coocood/freecache"
)

func NewCache(cacheSize int) *freecache.Cache {
	return freecache.NewCache(cacheSize)
}
