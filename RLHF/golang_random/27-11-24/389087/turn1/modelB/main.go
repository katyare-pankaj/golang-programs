package main

import (
	"fmt"
	"time"
)

// Define an Observable type as a channel of integers
type Observable <-chan int

// A function that generates a stream of integers
func generateIntegerStream() Observable {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
			time.Sleep(time.Second)
		}
		close(out)
	}()
	return out
}

// A function that prints the integers from the stream
func printIntegers(stream Observable) {
	for num := range stream {
		fmt.Println(num)
	}
}

// A function that filters even integers from the stream
func filterEvenIntegers(stream Observable) Observable {
	out := make(chan int)
	go func() {
		for num := range stream {
			if num%2 == 0 {
				out <- num
			}
		}
		close(out)
	}()
	return out
}

func main() {
	stream := generateIntegerStream()
	// Print all integers from the stream
	printIntegers(stream)
	// Print only even integers from the stream
	evenStream := filterEvenIntegers(stream)
	printIntegers(evenStream)
}
