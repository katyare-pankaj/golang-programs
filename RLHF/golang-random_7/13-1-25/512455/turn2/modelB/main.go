package main

import (
	"fmt"
	"sync"
)

func printSquare(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Square of %d is: %d\n", num, num*num)
}

func main() {
	var wg sync.WaitGroup
	numbers := []int{1, 2, 3, 4, 5}
	wg.Add(len(numbers))

	for _, num := range numbers {
		go printSquare(num, &wg)
	}

	wg.Wait()
	fmt.Println("All computations are finished.")
}
