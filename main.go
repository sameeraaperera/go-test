package main

import (
	"fmt"
	"github.com/sap4001/nearmap-go/datasource"
	"math/rand"
	"sync"
	"time"
)

func main() {
	db := getPopulatedDB()
	dc := datasource.NewDistributedCache()
	ds := datasource.NewDataStore(db, dc)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go makeRandomRequests(&wg, ds, i+1, 50)
	}
	fmt.Println("Main: Waiting for workers to finish")
	wg.Wait()
	fmt.Println("Main: Completed")
}

func getPopulatedDB() datasource.Database {
	db := datasource.NewDatabase()
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		db.Store(key, value)
	}
	return db
}

func makeRandomRequests(wg *sync.WaitGroup, ds datasource.DataStore, id, count int) {
	//fmt.Printf("Worker %v: Started\n", id)
	defer wg.Done()
	for i := 0; i < 50; i++ {
		randKey := getRandKey()
		start := time.Now()
		ds.Value(randKey)
		elapsed := time.Since(start)
		fmt.Printf("RoutineID:%d [%d] Request %q time: %s\n", id, i, randKey, elapsed)
	}
	//fmt.Printf("Worker %v: Finished\n", id)
}

func getRandKey() string {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 10
	randI := rand.Intn(max-min) + min
	//randI = 2 //////TODO
	return fmt.Sprintf("key%d", randI)
}
