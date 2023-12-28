package cacheit_test

import (
	"cacheit"
	"cacheit/shared"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	cache := cacheit.New(shared.CacheConfig{})

	if cache == nil {
		t.Error("Can not Initialize cache")
	}
}

func TestSetDefaultExpirationTime(t *testing.T) {
	cache := cacheit.New(shared.CacheConfig{})

	if cache == nil {
		t.Error("Can not Initialize cache")
	}

	const timeToBeSet = time.Minute * 10
	cache.SetDefaultExpirationTime(timeToBeSet)
	defaultTime := cache.GetDefaultExpirationTime()

	if timeToBeSet != defaultTime {
		t.Error("Error setting default expiration time")
	}
}

func TestAdd(t *testing.T) {

	cache := cacheit.New(shared.CacheConfig{})

	if cache == nil {
		t.Error("Can not Initialize cache")
	}

	cache.Add("test", "testData", time.Minute*2)
	data, _ := cache.Get("test")
	if data == nil {
		t.Error("Can not retrieve data")
		return
	}
	t.Log("data retried")

}
