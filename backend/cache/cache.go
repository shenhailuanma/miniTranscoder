package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var cacheCache *cache.Cache

func init() {
	cacheCache = cache.New(time.Minute, time.Minute)
}

func ItemCount() int {
	return cacheCache.ItemCount()
}
