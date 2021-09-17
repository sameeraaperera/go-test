package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/sap4001/nearmap-go/datasource"
)

func main() {
	ds := initDataStore()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go RandomRequestWorker(&wg, ds, i+1, 50)
	}
	wg.Wait()
	fmt.Println("End of Main.")
}

func RandomRequestWorker(wg *sync.WaitGroup, ds datasource.DataStore, id, count int) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		randKey := getRandKey()
		start := time.Now()
		ds.Value(randKey)
		elapsed := time.Since(start)
		fmt.Printf("RoutineIDEdited:%d [%d] Request %q time: %s\n", id, i, randKey, elapsed)
	}
}

func initDataStore() datasource.DataStore {
	db := initPopulatedDB()
	dc := datasource.NewDistributedCache()
	return datasource.NewDataStore(db, dc)
}

func initPopulatedDB() datasource.Database {
	db := datasource.NewDatabase()
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		db.Store(key, value)
	}
	return db
}

func getRandKey() string {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 10
	randI := rand.Intn(max-min) + min
	return fmt.Sprintf("key%d", randI)
}
