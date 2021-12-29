package cache

import (
	"github.com/patrickmn/go-cache"
	"github.com/shenhailuanma/miniTranscoder/models"
)

const CacheKeyPlaaylist = "CachePlaylist"

func GetCachePlaylist() ([]models.Job, bool) {
	cacheData, found := cacheCache.Get(CacheKeyPlaaylist)
	if found {
		cacheValue, ok := cacheData.([]models.Job)
		if ok {
			return cacheValue, true
		}
	}

	return []models.Job{}, false
}

func SetCachePlaylist(accounts []models.Job) {
	cacheCache.Set(CacheKeyPlaaylist, accounts, cache.DefaultExpiration)
}

func CleanCachePlaylist()  {
	cacheCache.Delete(CacheKeyPlaaylist)
}
