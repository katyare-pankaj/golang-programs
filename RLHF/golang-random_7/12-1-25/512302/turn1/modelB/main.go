package main

import (
	"fmt"
	"sync"
)

type Data struct {
	Key   int
	Value string
}

// Global cache variables
var cache1 = make(map[int]*Data)
var cache2 = make(map[int]*Data)
var cache3 = make(map[int]*Data)

var wg sync.WaitGroup

func main() {
	keys := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	wg.Add(len(keys))

	// Create a channel to receive data from the function
	dataCh := make(chan *Data)

	go retrieveData(keys, dataCh)

	for data := range dataCh {
		fmt.Printf("Key: %d, Value: %s\n", data.Key, data.Value)
	}

	wg.Wait()
}

func retrieveData(keys []int, dataCh chan *Data) {
	defer close(dataCh)

	// Range loop over the keys
	for _, key := range keys {
		// Check if the data is in the first-level cache
		if data, found := cache1[key]; found {
			fmt.Println("Data found in cache1.")
			dataCh <- data
			wg.Done()
			continue
		}

		// Check if the data is in the second-level cache
		if data, found := cache2[key]; found {
			fmt.Println("Data found in cache2.")
			cache1[key] = data // Move the data to the first-level cache
			dataCh <- data
			wg.Done()
			continue
		}

		// Check if the data is in the third-level cache
		if data, found := cache3[key]; found {
			fmt.Println("Data found in cache3.")
			cache1[key] = data // Move the data to the first-level cache
			cache2[key] = data // Move the data to the second-level cache
			dataCh <- data
			wg.Done()
			continue
		}

		// If the data is not found in any cache, retrieve it from the database (simulated here)
		fmt.Println("Data not found in any cache. Retrieving from database...")
		data := &Data{Key: key, Value: fmt.Sprintf("Value for key %d", key)}
		cache1[key] = data // Move the data to the first-level cache
		cache2[key] = data // Move the data to the second-level cache
		cache3[key] = data // Move the data to the third-level cache
		dataCh <- data
		wg.Done()
	}
}
