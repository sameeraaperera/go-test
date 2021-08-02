package datasource

import (
	"fmt"
	"sync"
)

// DataStore acts as a very fast key value store and acts as a wrapper for Database and DistributedCache
type DataStore struct {
	mu   *sync.Mutex
	data map[string]string
	db   *Database
	dc   *DistributedCache
}

func NewDataStore(db Database, dc DistributedCache) DataStore {
	return DataStore{
		mu:   &sync.Mutex{},
		data: make(map[string]string),
		db:   &db,
		dc:   &dc,
	}
}

// Value returns a value in max 500ms
func (ds *DataStore) Value(key string) (string, error) {
	// Initially check instant local data store
	ds.mu.Lock()
	value, found := ds.data[key]
	ds.mu.Unlock()
	if found {
		return value, nil
	}

	// Then check distributed cache
	value, err := ds.dc.Value(key)
	if err == nil {
		//store value in local store
		ds.Store(key, value)
		return value, nil
	}

	// Finally check the Database
	value, err = ds.db.Value(key)
	if err == nil {
		//store value in cache
		ds.dc.Store(key, value)
		return value, nil
	}

	return "", fmt.Errorf("key not found: %q", key)
}

// Store saves a value
func (ds *DataStore) Store(key string, value string) error {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	ds.data[key] = value
	return nil
}
