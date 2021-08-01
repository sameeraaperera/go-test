package datasource

import (
	"fmt"
	"time"
)

//Database persistent key value store
type Database struct {
	data map[string]string
}

// NewDatabase returns a new db instance
func NewDatabase() Database {
	return Database{
		data: make(map[string]string),
	}
}

// Value returns a value in DB in 500ms
func (db *Database) Value(key string) (string, error) {
	// simulate 500ms roundtrip to the distributed cache
	time.Sleep(500 * time.Millisecond)
	val, found := db.data[key]
	if !found {
		return "", fmt.Errorf("key not found in db: %q", key)
	}
	return val, nil
}

// Store saves a value in DB in 500ms
func (db *Database) Store(key string, value string) error {
	// simulate 500ms roundtrip to the distributed cache
	time.Sleep(500 * time.Millisecond)

	db.data[key] = value
	return nil
}
