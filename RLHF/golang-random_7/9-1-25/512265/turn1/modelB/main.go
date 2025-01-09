package main

import (
	"fmt"
	"sync"
)

const numWorkers = 4
const sliceSize = 1000000

func main() {
	data := make([]int, sliceSize)
	for i := 0; i < sliceSize; i++ {
		data[i] = i + 1
	}
	result := make(chan int)
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for w := 1; w <= numWorkers; w++ {
		go worker(w, data, result, &wg)
	}
	go func() {
		wg.Wait()
		close(result)
	}()
	sum := 0
	for res := range result {
		sum += res
	}
	fmt.Println("Sum:", sum)
}

func worker(id int, data []int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	localSum := 0
	for _, num := range data {
		localSum += num
	}
	result <- localSum
}
