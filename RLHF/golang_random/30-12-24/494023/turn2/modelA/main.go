package main

import (
	"fmt"
	"sync"
)

func incrementCounter(wg *sync.WaitGroup, key interface{}, countMap *sync.Map) {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		value, loaded := countMap.LoadOrStore(key, 0)
		countMap.Store(key, value.(int)+1)
	}
}

func main() {
	var wg sync.WaitGroup
	keys := []interface{}{"apple", "banana", "orange"}
	countMap := new(sync.Map)

	for _, key := range keys {
		wg.Add(1)
		go incrementCounter(&wg, key, countMap)
	}

	wg.Wait()

	fmt.Println("Final Counter Values:")
	countMap.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value.(int))
		return true
	})
}
