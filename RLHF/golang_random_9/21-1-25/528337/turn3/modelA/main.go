package main

import (
	"fmt"
	"sync"
)

func processData(data []int, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for _, value := range data {
		fmt.Printf("Goroutine %d processing: %d\n", id, value)
	}
}

func main() {
	var wg sync.WaitGroup

	dataSlices := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	}

	for id, data := range dataSlices {
		wg.Add(1)
		go processData(data, &wg, id+1) // id+1 to make identifiers start from 1
	}

	wg.Wait()
	fmt.Println("All data processed.")
}
