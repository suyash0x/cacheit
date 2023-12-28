package cacheit

import (
	"cacheit/cache"
	"cacheit/shared"
)

func New(cc shared.CacheConfig) *cache.Cache {

	return &cache.Cache{}
}
