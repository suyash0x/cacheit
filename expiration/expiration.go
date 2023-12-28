package expiration

import (
	"cacheit/cache"
	"log"
)

type Expiration struct{}

func (e *Expiration) CheckForExpiration(c *cache.Cache) {
	log.Println("Cache Expiration check is running on Background")
}

func New() *Expiration {
	return &Expiration{}
}
