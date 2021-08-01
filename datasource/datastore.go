package datasource

import (
	"fmt"
)

// DataStore acts as a very fast key value store and acts as a wrapper for Database and DistributedCache
type DataStore struct {
	data map[string]string
	db   *Database
	dc   *DistributedCache
}

func NewDataStore(db Database, dc DistributedCache) DataStore {
	return DataStore{
		make(map[string]string),
		&db,
		&dc,
	}
}

//func populateData(db Database) {
//	for i := 0; i < 10; i++ {
//		key := fmt.Sprintf("key%d", i)
//		value := fmt.Sprintf("value%d", i)
//		db.Store(key, value)
//	}
//}

// Value returns a value in max 500ms
func (ds *DataStore) Value(key string) (string, error) {
	// Initially check instant local data store
	value, found := ds.data[key]
	if found {
		return value, nil
	}

	// Then check distributed cache
	value, err := ds.dc.Value(key)
	if err == nil {
		//store value in local store
		ds.data[key] = value
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
	ds.data[key] = value

	return nil
}
