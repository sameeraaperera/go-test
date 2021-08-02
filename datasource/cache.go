package datasource

import (
	"fmt"
	"sync"
	"time"
)

// DistributedCache very fast key value store
type DistributedCache struct {
	mu *sync.RWMutex
	data map[string]string
}

// NewDistributedCache returns a new
func NewDistributedCache() DistributedCache {
	return DistributedCache{
		mu:   &sync.RWMutex{},
		data: make(map[string]string),
	}
}

// Value returns a value in 100ms
func (dc *DistributedCache) Value(key string) (string, error) {
	// simulate 100ms roundtrip to the distributed cache
	time.Sleep(100 * time.Millisecond)
	dc.mu.RLock()
	defer dc.mu.RUnlock()

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

	dc.mu.Lock()
	defer dc.mu.Unlock()

	dc.data[key] = value
	return nil
}
