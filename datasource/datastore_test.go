package datasource_test

import (
	"github.com/sap4001/nearmap-go/datasource"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetValueDoesNotExist(t *testing.T) {
	ds := datasource.NewDataStore(datasource.NewDatabase(), datasource.NewDistributedCache())
	got, err := ds.Value("nokey")

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "key not found: \"nokey\"")
	assert.Empty(t, got)
}

func Test_GetValueExistInCache(t *testing.T) {
	key := "cachekey"
	dc := datasource.NewDistributedCache()
	dc.Store(key, "some cache value")
	ds := datasource.NewDataStore(datasource.NewDatabase(), dc)
	got, err := ds.Value(key)

	assert.Nil(t, err)
	assert.Equal(t, "some cache value", got)
}

func Test_GetValueExistInDB(t *testing.T) {
	key := "dbkey"
	db := datasource.NewDatabase()
	dc := datasource.NewDistributedCache()
	db.Store(key, "some db value")
	ds := datasource.NewDataStore(db, dc)

	_, err := dc.Value(key)
	assert.NotNil(t, err, "cache was not empty")

	got, err := ds.Value("dbkey")
	assert.Nil(t, err)
	assert.Equal(t, "some db value", got)

	cachedVal, err := dc.Value(key)
	assert.Nil(t, err)
	assert.Equal(t, "some db value", cachedVal)
}

func Test_StoreValue(t *testing.T) {
	key := "storevaluekey"
	db := datasource.NewDatabase()
	dc := datasource.NewDistributedCache()
	ds := datasource.NewDataStore(db, dc)
	ds.Store(key, "stored value")

	val, err := ds.Value(key)
	assert.Nil(t, err)
	assert.Equal(t, "stored value", val)

}
