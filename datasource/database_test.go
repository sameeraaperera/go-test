package datasource_test

import (
	"github.com/sap4001/nearmap-go/datasource"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_DBGetValueDoesNotExist(t *testing.T) {
	db := datasource.NewDatabase()
	got, err := db.Value("nokey")

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "key not found in db: \"nokey\"")
	assert.Empty(t, got)
}

func Test_DBStoreAndGetValue(t *testing.T) {
	db := datasource.NewDatabase()
	db.Store("yeskey", "dbvalue")
	got, err := db.Value("yeskey")

	assert.Nil(t, err)
	assert.Equal(t, got, "dbvalue")
}
