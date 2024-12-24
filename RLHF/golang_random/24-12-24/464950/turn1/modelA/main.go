package main

import (
	"fmt"
	"math/rand"
	"time"
)

func slowOperation() {
	time.Sleep(time.Millisecond * 10)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Starting...")

	for i := 0; i < 10000; i++ {
		defer slowOperation()
		fmt.Println("Iteration", i)
	}

	fmt.Println("Finished.")
}
