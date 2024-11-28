package main

import (
 "fmt"
 "math/rand"
 "sync"
 "time"
)

const (
 numProducers  = 3
 numConsumers  = 3
 bufferSize = 100   // Channel buffer size
)

func producer(id int, jobs chan<- int, wg *sync.WaitGroup) {
 defer wg.Done()
 for {
  job := rand.Intn(100)
  jobs <- job
  fmt.Printf("Producer %d: Produced job %d\n", id, job)
  time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
 }
}

func consumer(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
 defer wg.Done()
 for job := range jobs {
  result := job * job
  results <- result
  fmt.Printf("Consumer %d: Processed job %d, Result: %d\n", id, job, result)
  time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
 }
}

func main() {
 // Create channels for communication
 jobs := make(chan int, bufferSize)
 results := make(chan int)

 // Create WaitGroups
 var wgProducers sync.WaitGroup
 var wgConsumers sync.WaitGroup
 var wgResult sync.WaitGroup

 wgProducers.Add(numProducers)
 wgConsumers.Add(numConsumers)
 wgResult.Add(1)

 // Start producers
 for i := 0; i < numProducers; i++ {
  go producer(i, jobs, &wgProducers)
 }

 // Start consumers
 for i := 0; i < numConsumers; i++ {
  go consumer(i, jobs, results, &wgConsumers)
 }

 // Start a routine to read results
 go func() {
  defer wgResult.Done()
  for result := range results {
   fmt.Printf("Main: Result: %d\n", result)
  }
 }()