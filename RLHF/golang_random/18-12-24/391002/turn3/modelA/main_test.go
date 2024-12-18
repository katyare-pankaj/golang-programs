package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type result struct {
	task  int
	error error
	value int
}

func processTask(id int, workCh <-chan int, resultCh chan<- result, wg *sync.WaitGroup) {
	defer wg.Done()

	select {
	case task, ok := <-workCh:
		if !ok {
			return
		}

		// Simulate processing
		time.Sleep(time.Millisecond * 50)

		if task%2 == 0 {
			resultCh <- result{task: task, error: fmt.Errorf("even task failure: %d", task)}
		} else {
			resultCh <- result{task: task, error: nil, value: task * task}
		}

	default:
		// If workCh is closed, simply return
		return
	}
}

func processWork(numTasks int) []result {
	var wg sync.WaitGroup
	workCh := make(chan int, numTasks)
	resultCh := make(chan result)

	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		go processTask(i, workCh, resultCh, &wg)
	}

	for i := 0; i < numTasks; i++ {
		workCh <- i
	}

	close(workCh)

	var results []result
	wg.Wait()

	close(resultCh)

	for r := range resultCh {
		results = append(results, r)
	}

	return results
}

func TestProcessWork(t *testing.T) {
	numTasks := 10
	results := processWork(numTasks)

	var errorCount int
	var successCount int

	for _, result := range results {
		if result.error != nil {
			errorCount++
		} else {
			successCount++
		}
	}

	if errorCount != numTasks/2 {
		t.Errorf("expected %d errors, but got %d", numTasks/2, errorCount)
	}

	if successCount != numTasks/2 {
		t.Errorf("expected %d successes, but got %d", numTasks/2, successCount)
	}

	for _, result := range results {
		if result.error != nil {
			expectedError := fmt.Errorf("even task failure: %d", result.task)
			if result.error.Error() != expectedError.Error() {
				t.Errorf("unexpected error: expected %v, got %v", expectedError, result.error)
			}
		} else {
			if result.value != result.task*result.task {
				t.Errorf("expected square of task %d, but got %d", result.task, result.value)
			}
		}
	}
}

func main() {
	// For demonstration purposes, run the test
	t := &testing.T{}
	TestProcessWork(t)
	fmt.Println("All tests passed.")
}
