package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func task1(wg *sync.WaitGroup) {
	defer wg.Done() //signal that this goroutine is done
	fmt.Println("Starting task 1...")
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	fmt.Println("Task 1 completed.")
}

func task2(wg *sync.WaitGroup) {
	defer wg.Done() //signal that this goroutine is done
	fmt.Println("Starting task 2...")
	time.Sleep(time.Duration(rand.Intn(4)) * time.Second)
	fmt.Println("Task 2 completed.")
}

func main() {
	var wg sync.WaitGroup //Create a WaitGroup

	//Add two tasks to the WaitGroup
	wg.Add(2)

	//Run task1 and task2 as separate goroutines
	go task1(&wg)
	go task2(&wg)

	//Wait for both tasks to finish
	fmt.Println("Waiting for tasks to complete...")
	wg.Wait()

	fmt.Println("Both tasks completed.")
}
