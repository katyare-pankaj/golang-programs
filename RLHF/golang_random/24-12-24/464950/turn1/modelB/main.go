package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Generate some random data for testing
	data := make([]int, 1000000)
	for i := range data {
		data[i] = i
	}

	// Profile the execution of the function with defer
	start := time.Now()
	funcWithDefer(data)
	elapsed := time.Since(start)
	fmt.Println("Execution time with defer:", elapsed)

	// Profile the execution of the function without defer
	start = time.Now()
	funcWithoutDefer(data)
	elapsed = time.Since(start)
	fmt.Println("Execution time without defer:", elapsed)
}

func funcWithDefer(data []int) {
	var sum int

	// Open a file for writing results
	file, err := os.Create("results.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for _, value := range data {
		sum += value
	}

	// Write the sum to the file
	_, err = file.WriteString(fmt.Sprintf("Sum: %d\n", sum))
	if err != nil {
		panic(err)
	}
}

func funcWithoutDefer(data []int) {
	var sum int

	// Open a file for writing results
	file, err := os.Create("results.txt")
	if err != nil {
		panic(err)
	}

	defer func() {
		// Close the file at the end of the function
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	for _, value := range data {
		sum += value
	}

	// Write the sum to the file
	_, err = file.WriteString(fmt.Sprintf("Sum: %d\n", sum))
	if err != nil {
		panic(err)
	}
}
