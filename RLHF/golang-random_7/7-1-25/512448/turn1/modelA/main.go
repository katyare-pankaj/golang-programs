package main

import (
	"fmt"
	"sync"
)

func processData(data string, wg *sync.WaitGroup) {
	defer wg.Done()
	// Simulate processing data
	fmt.Println("Processing:", data)
}

func main() {
	var wg sync.WaitGroup
	dataList := []string{"data1", "data2", "data3", "data4"}

	for _, data := range dataList {
		wg.Add(1)
		go processData(data, &wg)
	}

	wg.Wait()
	fmt.Println("All data processing completed.")
}
