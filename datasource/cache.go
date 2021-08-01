package datasource

import (
	"fmt"
	"time"
)

// DistributedCache very fast key value store
type DistributedCache struct {
	data map[string]string
}

// NewDistributedCache returns a new
func NewDistributedCache() DistributedCache {
	return DistributedCache{
		data: make(map[string]string),
	}
}

// Value returns a value in 100ms
func (dc *DistributedCache) Value(key string) (string, error) {
	// simulate 100ms roundtrip to the distributed cache
	time.Sleep(100 * time.Millisecond)
	val, found := dc.data[key]
	if !found {
		return "", fmt.Errorf("key not found in cache: %q", key)
	}
	return val, nil
}

// Store saves a value in 100ms
func (dc *DistributedCache) Store(key string, value string) error {
	// simulate 100ms round trip to the distributed cache
	time.Sleep(100 * time.Millisecond)

	dc.data[key] = value
	return nil
}
