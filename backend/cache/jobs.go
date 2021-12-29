package cache

import (
	"github.com/patrickmn/go-cache"
	"github.com/shenhailuanma/miniTranscoder/models"
)

const CacheKeyJobs = "CacheJobs"

func GetCacheJobs() ([]models.Job, bool) {
	cacheData, found := cacheCache.Get(CacheKeyJobs)
	if found {
		cacheValue, ok := cacheData.([]models.Job)
		if ok {
			return cacheValue, true
		}
	}

	return []models.Job{}, false
}

func SetCacheJobs(data []models.Job) {
	cacheCache.Set(CacheKeyJobs, data, cache.DefaultExpiration)
}

func CleanCacheJobs()  {
	cacheCache.Delete(CacheKeyJobs)
}
