package main

import (
	"fmt"
	"sync"
)

func factorial(n int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	c <- result
}
func main() {
	var wg sync.WaitGroup
	c := make(chan int)

	n := 10
	fmt.Printf("Calculating factorial of %d using concurrency:\n", n)
	wg.Add(n)

	for i := 1; i <= n; i++ {
		go factorial(i, &wg, c)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	result := 1
	for res := range c {
		result *= res
	}

	fmt.Println("Result:", result)
}
