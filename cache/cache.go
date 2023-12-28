package cache

import (
	"time"
)

var DefaultExpirationTime = time.Minute * 5

type Cache struct {
	DefaultExpirationTime time.Duration
}
