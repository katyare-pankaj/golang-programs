package main

import (
	"fmt"
)

func sumGoroutine(numbers chan int, result chan int) {
	total := 0
	for num := range numbers {
		total += num
	}
	result <- total
	close(result) // Indicate that the result has been calculated
}

func main() {
	numbers := make(chan int, 5)
	result := make(chan int)

	go sumGoroutine(numbers, result)

	// Send some numbers to the channel
	numbers <- 1
	numbers <- 2
	numbers <- 3
	numbers <- 4
	numbers <- 5

	close(numbers) // Close the numbers channel

	// Receive the result
	fmt.Println("The sum is:", <-result)
}