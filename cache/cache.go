package cache

import (
	"cacheit/err"
	"errors"
	"time"
)

// Default Expiration Time. If not set explicitly, expiration time will be default to 5 mins
var Det = time.Minute * 5

type ExpirationModule interface {
	CheckForExpiration(*Cache)
}

type CacheItem struct {
	expirationTime time.Time
	data           any
}

type Cache struct {
	CacheData map[string]CacheItem
	ExpirationModule
}

func (c *Cache) checkExpired(t time.Time) (expired bool) {
	expired = time.Now().After(t)
	return
}

// SetDefaultExpirationTime sets the default expiration time for cache items.
func (c *Cache) SetDefaultExpirationTime(t time.Duration) {
	Det = t
}

// Get Default Expiration Time
func (c *Cache) GetDefaultExpirationTime() time.Duration {
	return Det
}

// Add data with specified key and expiration time
func (c *Cache) Add(key string, data interface{}, expirationTime time.Duration) {
	validTime := time.Now().Add(expirationTime)
	cacheItem := CacheItem{
		expirationTime: validTime,
		data:           data,
	}
	c.CacheData[key] = cacheItem

}

// Check if Item with key is present in cache
func (c *Cache) Has(key string) (found bool) {
	_, found = c.CacheData[key]
	return
}

// Get Data with specific key
func (c *Cache) Get(key string) (interface{}, error) {
	item, found := c.CacheData[key]

	if !found {
		return nil, errors.New(err.CACHE_NOT_FOUND)
	}

	expired := c.checkExpired(item.expirationTime)

	if expired {
		return nil, errors.New(err.CACHE_EXPIRED)
	}
	return item, nil
}
