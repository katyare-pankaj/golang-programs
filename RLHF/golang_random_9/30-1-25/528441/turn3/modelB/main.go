package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

const (
    numWorkers    = 4      // Initial number of worker goroutines
    bufferSize    = 1000    // Buffer size for worker channels
    maxLoadFactor = 0.8    // Maximum load factor before spawning a new worker
    minLoadFactor = 0.2    // Minimum load factor before stopping a worker
    restartDelay  = 5 * time.Second // Delay before restarting a failed worker
)

type Worker struct {
    id         int
    taskChannel chan int
    wg         *sync.WaitGroup
    quit       chan struct{}
    errChannel chan error // Channel to report worker errors
}

func newWorker(id int, wg *sync.WaitGroup, errChannel chan error) *Worker {
    return &Worker{
        id:         id,
        taskChannel: make(chan int, bufferSize),
        wg:         wg,
        quit:       make(chan struct{}),
        errChannel: errChannel,
    }
}

func (w *Worker) start() {
    w.wg.Add(1)
    go func() {
        defer w.wg.Done()
        for {
            select {
            case task, ok := <-w.taskChannel:
                if !ok {
                    return // Worker channel is closed, exit
                }
                // Simulate processing data
                // Introduce a random error to simulate worker failures
                if rand.Float32() < 0.05 {
                    w.errChannel <- fmt.Errorf("worker %d: simulation error", w.id)
                    return
                }
                time.Sleep(time.Duration(task) * time.Millisecond)
                fmt.Printf("Worker %d processed task %d\n", w.id, task)
            case <-w.quit:
                return // Worker is quitting
            }
        }
    }()
}

func (w *Worker) stop() {
    close(w.taskChannel) // Close the worker channel to signal termination
    <-w.quit             // Wait for the worker to finish
}

type LoadBalancer struct {
    workers     []*Worker
    nextWorker  int
    loadFactor  float32
    wg         *sync.WaitGroup
    taskChannel chan int
    mu          sync.Mutex
    errChannel chan error
    quit        chan struct{}
}

func newLoadBalancer(wg *sync.WaitGroup) *LoadBalancer {
    lb := &LoadBalancer{
        workers:     make([]*Worker, 0, numWorkers),
        nextWorker:  0,
        loadFactor:  0,
        wg:         wg,
        taskChannel: make(chan int, bufferSize),
        errChannel: make(chan error),
        quit:        make(chan struct{}),
    }

    // Start the initial set of workers
    for i := 0; i < numWorkers; i++ {
        w := newWorker(i, lb.wg, lb.errChannel)
        w.start()
        lb.workers = append(lb.workers, w)
    }

    return lb
}

func (lb *LoadBalancer) start() {
    lb.wg.Add(1)
    go func() {
        defer lb.wg.Done()
        for {
            select {
            case task, ok := <-lb.taskChannel:
                if !ok {
                    return // Load balancer channel is closed, exit
                }

                // Distribute tasks to workers using round-robin
                worker := lb.workers[lb.nextWorker]
                lb.nextWorker = (lb.nextWorker + 1) % len(lb.workers)