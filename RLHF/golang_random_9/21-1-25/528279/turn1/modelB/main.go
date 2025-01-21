package main

import (
	"fmt"
	"sync"
)

func processData(data []int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, value := range data {
		fmt.Println("Processing:", value)
	}
}

func main() {
	var wg sync.WaitGroup
	dataStream := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	for _, data := range dataStream {
		wg.Add(1)
		go processData(data, &wg)
	}

	wg.Wait()
	fmt.Println("All data processed.")
}
