package cacheit

import (
	"cacheit/cache"
	"cacheit/expiration"
	"cacheit/shared"
	"sync"
)

func New(cc shared.CacheConfig) (cacheInstance *cache.Cache) {
	cacheInstance = &cache.Cache{
		ExpirationModule: expiration.New(),
		CacheData:        map[string]cache.CacheItem{},
		Mtx:              &sync.RWMutex{},
	}

	go cacheInstance.CheckForExpirationAndRemove(cacheInstance)

	return
}
