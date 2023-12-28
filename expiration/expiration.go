package expiration

import (
	"cacheit/cache"
	"log"
	"time"
)

// Default Expiration Time. If not set explicitly, expiration time will be default to 5 mins
var Det = time.Minute * 5

type Expiration struct {
}

func (e *Expiration) CheckForExpirationAndRemove(c *cache.Cache) {
	log.Println("Cache Expiration check is running on Background")

	for key, cacheItem := range c.CacheData {
		if e.ItemExpired(cacheItem.ExpirationTime) {
			c.Remove(key)
		}
	}
}

func (e *Expiration) SetTime(t time.Duration) {
	Det = t
}

func (e *Expiration) GetTime() time.Duration {
	return Det
}

func (e *Expiration) ItemExpired(t time.Time) (expired bool) {
	expired = time.Now().After(t)
	return
}

func New() *Expiration {
	return &Expiration{}
}
