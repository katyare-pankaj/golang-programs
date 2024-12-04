package main

import (
	"fmt"
	"log"
)

func panicIfInvalid(x int) {
	if x <= 0 {
		panic(fmt.Sprintf("x must be greater than zero: %d", x))
	}
}

func criticalOperation(x int) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Critical operation recovered from panic: %v\n", r)
		}
	}()

	fmt.Println("Starting critical operation...")

	panicIfInvalid(x)

	fmt.Println("Critical operation completed successfully.")
}

func main() {
	// Example usage with invalid input that causes a panic
	criticalOperation(-1)

	fmt.Println("Program continues execution after panic is recovered.")
}
