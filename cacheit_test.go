package cacheit_test

import (
	"cacheit"
	"cacheit/shared"
	"testing"
)

func TestNew(t *testing.T) {
	cache := cacheit.New(shared.CacheConfig{})

	if cache == nil {
		t.Error("Can not Initialize cache")
	}
}
