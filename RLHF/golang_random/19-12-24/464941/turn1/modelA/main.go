package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, taskChannel <-chan string, resultChannel chan<- string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker terminated due to context cancellation.")
			return
		case task := <-taskChannel:
			fmt.Println("Processing task:", task)
			time.Sleep(time.Second) // Simulate work
			result := "Processed: " + task
			resultChannel <- result
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	taskChannel := make(chan string)
	resultChannel := make(chan string)

	go worker(ctx, taskChannel, resultChannel)

	tasks := []string{"Task 1", "Task 2", "Task 3"}
	for _, task := range tasks {
		taskChannel <- task
	}

	close(taskChannel)

	for result := range resultChannel {
		fmt.Println(result)
	}
}
