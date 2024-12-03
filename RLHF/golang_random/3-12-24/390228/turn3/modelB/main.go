package main

import (
	"fmt"
	"sync"
	"time"
)

// Producer function to generate items and send them to the channel
func producer(id int, jobs chan<- int, count int) {
	for i := 0; i < count; i++ {
		fmt.Printf("Producer %d producing job %d\n", id, i)
		jobs <- i
		time.Sleep(time.Millisecond * 100) // Simulate production time
	}
}

// Consumer function to process items from the channel
func consumer(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Consumer %d processing job %d\n", id, job)
		time.Sleep(time.Millisecond * 150) // Simulate processing time
	}
}

func main() {
	const numProducers = 2
	const numConsumers = 3
	const numJobsPerProducer = 5

	jobs := make(chan int, 10) // Buffered channel to handle jobs
	var wg sync.WaitGroup

	// Start producers
	for p := 1; p <= numProducers; p++ {
		go producer(p, jobs, numJobsPerProducer)
	}

	// Start consumers
	for c := 1; c <= numConsumers; c++ {
		wg.Add(1)
		go consumer(c, jobs, &wg)
	}

	// Close the jobs channel after all producers are done
	go func() {
		time.Sleep(time.Second) // Adjust sleep to allow producers to finish
		close(jobs)
	}()

	// Wait for all consumers to finish
	wg.Wait()

	fmt.Println("All jobs processed.")
}
