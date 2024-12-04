package main

import "fmt"

func process(i int) int {
	fmt.Println("Hello")
	return i
}
func worker(id int, tasks <-chan int, results chan<- int) {
	for task := range tasks {
		results <- process(task) // Represents processing logic
	}
}

func main() {
	const numWorkers = 4
	tasks := make(chan int, 100)
	results := make(chan int, 100)

	for w := 0; w < numWorkers; w++ {
		go worker(w, tasks, results)
	}

	// Send tasks
	go func() {
		for i := 0; i < 100; i++ {
			tasks <- i
		}
		close(tasks)
	}()

	// Collect results
	for i := 0; i < 100; i++ {
		<-results
	}
}
