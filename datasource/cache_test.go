package datasource_test

import (
	"github.com/sap4001/nearmap-go/datasource"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CacheGetValueDoesNotExist(t *testing.T) {
	dc := datasource.NewDistributedCache()
	got, err := dc.Value("nokey")

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "key not found in cache: \"nokey\"")
	assert.Empty(t, got)
}

func Test_CacheStoreAndGetValue(t *testing.T) {
	dc := datasource.NewDistributedCache()
	dc.Store("yeskey", "somevalue")
	got, err := dc.Value("yeskey")

	assert.Nil(t, err)
	assert.Equal(t, got, "somevalue")
}
