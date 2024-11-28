package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func task1() {
	defer wg.Done()
	fmt.Println("Started task 1")
	time.Sleep(1 * time.Second)
	fmt.Println("Finished task 1")
}

func task2() {
	defer wg.Done()
	fmt.Println("Started task 2")
	time.Sleep(1 * time.Second)
	fmt.Println("Finished task 2")
}

func task3() {
	defer wg.Done()
	fmt.Println("Started task 3")
	time.Sleep(1 * time.Second)
	fmt.Println("Finished task 3")
}

func task4() {
	defer wg.Done()
	fmt.Println("Started task 4")
	time.Sleep(1 * time.Second)
	fmt.Println("Finished task 4")
}

func main() {
	wg.Add(4) // Add the number of tasks to the wait group
	go task1()
	go task2()
	go task3()
	go task4()

	// Wait for all tasks to complete
	wg.Wait()
	fmt.Println("All tasks have completed")
}
