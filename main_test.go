package main

import (
	"fmt"
	"github.com/sap4001/nearmap-go/datasource"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func Test_GetRandKey(t *testing.T) {
	got := getRandKey()
	assert.NotNil(t, got)
	assert.Regexp(t, "key[0-9]", got)
}

func Test_InitDataStore(t *testing.T) {
	ds := initDataStore()
	AssertPopulatedDataSource(t, &ds)
}

func AssertPopulatedDataSource(t *testing.T, db datasource.DataSource) {
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		want := fmt.Sprintf("value%d", i)
		got, err := db.Value(key)

		assert.Nil(t, err)
		assert.Equal(t, want, got)
	}
}

func Benchmark_RandomRequestWorker(b *testing.B) {
	ds := initDataStore()
	var wg sync.WaitGroup
	for n := 0; n < b.N; n++ {
		wg.Add(1)
		RandomRequestWorker(&wg, ds, n, 50)
	}
}
