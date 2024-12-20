package main

import (
	"fmt"
	"sync"
)

func square(num int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	result := num * num
	ch <- result
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	results := make([]int, len(nums))
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(len(nums))

	for i, num := range nums {
		go square(num, &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for i := range nums {
		results[i] = <-ch
	}

	fmt.Println("Original slice:", nums)
	fmt.Println("Squared slice: ", results)
}
