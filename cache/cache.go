package cache

import (
	"cacheit/err"
	"errors"
	"sync"
	"time"
)

type ExpirationModule interface {
	CheckForExpirationAndRemove(*Cache)
	SetTime(time.Duration)
	GetTime() time.Duration
	ItemExpired(time.Time) (expired bool)
}

type CacheItem struct {
	ExpirationTime time.Time
	Data           any
}

type Cache struct {
	CacheData map[string]CacheItem
	ExpirationModule
	Mtx *sync.RWMutex
}

// SetDefaultExpirationTime sets the default expiration time for cache items.
func (c *Cache) SetDefaultExpirationTime(t time.Duration) {
	c.ExpirationModule.SetTime(t)
}

// Get Default Expiration Time
func (c *Cache) GetDefaultExpirationTime() time.Duration {
	return c.ExpirationModule.GetTime()
}

// Add data with specified key and expiration time
func (c *Cache) Add(key string, data interface{}, expirationTime time.Duration) {
	validTime := time.Now().Add(expirationTime)
	cacheItem := CacheItem{
		ExpirationTime: validTime,
		Data:           data,
	}
	c.Mtx.Lock()
	defer c.Mtx.Unlock()
	c.CacheData[key] = cacheItem

}

// Check if Item with key is present in cache
func (c *Cache) Has(key string) (found bool) {
	c.Mtx.RLock()
	defer c.Mtx.RUnlock()

	_, found = c.CacheData[key]
	return
}

// Get Data with specific key
func (c *Cache) Get(key string) (interface{}, error) {

	if !c.Has(key) {
		return nil, errors.New(err.CACHE_NOT_FOUND)
	}

	c.Mtx.RLock()
	defer c.Mtx.RUnlock()

	item := c.CacheData[key]

	if c.ExpirationModule.ItemExpired(item.ExpirationTime) {
		return nil, errors.New(err.CACHE_EXPIRED)
	}

	return item, nil
}

// Remove cache with specified key
func (c *Cache) Remove(key string) error {

	if !c.Has(key) {
		return errors.New(err.CACHE_NOT_FOUND)
	}

	c.Mtx.Lock()
	defer c.Mtx.Unlock()

	delete(c.CacheData, key)
	return nil
}
