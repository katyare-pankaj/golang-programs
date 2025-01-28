package main

import (
	"fmt"
	"log"
	"time"
)

// logExecution logs the execution time of a function
func logExecution(fn func()) func() {
	return func() {
		start := time.Now()
		fn()
		log.Printf("Execution time: %s\n", time.Since(start))
	}
}

func mainTask() {
	time.Sleep(2 * time.Second) // Simulate lengthy task
	fmt.Println("Main task completed")
}

func main() {
	loggedTask := logExecution(mainTask)
	loggedTask()
}
