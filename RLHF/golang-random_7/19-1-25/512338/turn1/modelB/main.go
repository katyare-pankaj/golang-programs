package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	callbackCount = 10000000
	cacheSize     = 10000
)

type Data struct {
	key   int
	value int
}

var dataCache map[int]*Data
var wg sync.WaitGroup

func callback(key int) {
	if _, ok := dataCache[key]; ok {
		return
	}
	// Simulate a cache miss by creating a new data structure and adding it to the cache
	data := &Data{key: key, value: key * 10}
	dataCache[key] = data
}
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	dataCache = make(map[int]*Data, cacheSize)
	startTime := time.Now()
	for i := 0; i < callbackCount; i++ {
		wg.Add(1)
		go callback(i)
	}
	wg.Wait()
	elapsedTime := time.Since(startTime)
	fmt.Printf("Execution time: %s\n", elapsedTime)
	// Analyze cache misses
	fmt.Printf("Cache size: %d, Cache misses: %d\n", len(dataCache), callbackCount-len(dataCache))
	// Optimize cache utilization using sync.Map
	optimizedCache := &sync.Map{}
	startTime = time.Now()
	for i := 0; i < callbackCount; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			_, _ = optimizedCache.LoadOrStore(key, &Data{key: key, value: key * 10})
		}(i)
	}
	wg.Wait()
	elapsedTime = time.Since(startTime)
	fmt.Printf("Optimized execution time: %s\n", elapsedTime)
}
