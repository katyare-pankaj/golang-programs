package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	initialWorkers = 4
	bufferSize     = 100
	maxLoadFactor  = 0.8
	minLoadFactor  = 0.2
	maxRetries     = 3
)

type Task struct {
	id      int
	retries int
}

type Worker struct {
	id       int
	taskChan chan Task
	done     chan struct{}
}

func newWorker(id int) *Worker {
	return &Worker{
		id:       id,
		taskChan: make(chan Task),
		done:     make(chan struct{}),
	}
}

func (w *Worker) start(wg *sync.WaitGroup, taskChan chan Task) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case task, ok := <-w.taskChan:
				if !ok {
					return
				}
				if err := processTask(task); err != nil {
					fmt.Printf("Worker %d failed task %d, retrying...\n", w.id, task.id)
					task.retries++
					if task.retries <= maxRetries {
						taskChan <- task
					} else {
						fmt.Printf("Task %d failed after %d retries\n", task.id, maxRetries)
					}
				} else {
					fmt.Printf("Worker %d completed task %d\n", w.id, task.id)
				}
			case <-w.done:
				return
			}
		}
	}()
}

func (w *Worker) stop() {
	close(w.done)
	close(w.taskChan)
}

func processTask(task Task) error {
	// Simulate random task failure
	if rand.Intn(10) < 2 {
		return fmt.Errorf("task %d failed", task.id)
	}
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	return nil
}

type LoadBalancer struct {
	taskQueue chan Task
	workers   []*Worker
	mu        sync.Mutex
}

func newLoadBalancer() *LoadBalancer {
	return &LoadBalancer{
		taskQueue: make(chan Task, bufferSize),
		workers:   make([]*Worker, 0, initialWorkers),
	}
}

func (lb *LoadBalancer) start(wg *sync.WaitGroup) {
	for i := 0; i < initialWorkers; i++ {
		worker := newWorker(i)
		worker.start(wg, lb.taskQueue)
		lb.workers = append(lb.workers, worker)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for task := range lb.taskQueue {
			lb.mu.Lock()
			if lb.shouldScaleUp() {
				lb.addWorker(wg)
			}
			worker := lb.selectWorker()
			worker.taskChan <- task
			lb.mu.Unlock()
		}
	}()
}

func (lb *LoadBalancer) stop() {
	for _, worker := range lb.workers {
		worker.stop()
	}
	close(lb.taskQueue)
}

func (lb *LoadBalancer) addWorker(wg *sync.WaitGroup) {
	worker := newWorker(len(lb.workers))
	worker.start(wg, lb.taskQueue)
	lb.workers = append(lb.workers, worker)
}

func (lb *LoadBalancer) selectWorker() *Worker {
	return lb.workers[rand.Intn(len(lb.workers))]
}

func (lb *LoadBalancer) shouldScaleUp() bool {
	totalTasks := len(lb.taskQueue)
	currentLoadFactor := float64(totalTasks) / float64(len(lb.workers)*bufferSize)
	return currentLoadFactor > maxLoadFactor
}

func main() {
	var wg sync.WaitGroup
	lb := newLoadBalancer()
	lb.start(&wg)

	// Simulate task generation
	go func() {
		for i := 0; i < 1000; i++ {
			lb.taskQueue <- Task{id: i}
		}
		close(lb.taskQueue)
	}()

	wg.Wait()
}
