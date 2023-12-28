package cacheit

import (
	"cacheit/cache"
	"cacheit/expiration"
	"cacheit/shared"
)

func New(cc shared.CacheConfig) (cacheInstance *cache.Cache) {
	cacheInstance = &cache.Cache{
		ExpirationModule: expiration.New(),
		CacheData:        map[string]cache.CacheItem{},
	}

	go cacheInstance.CheckForExpiration(cacheInstance)

	return
}
