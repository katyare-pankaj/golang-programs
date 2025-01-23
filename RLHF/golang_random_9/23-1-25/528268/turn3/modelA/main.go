package main

import (
	"fmt"
	"sync"
)

type BufferPool struct {
	pool sync.Pool
}

func NewBufferPool(bufferSize int) *BufferPool {
	return &BufferPool{
		pool: sync.Pool{
			New: func() interface{} {
				return make([]byte, 0, bufferSize)
			},
		},
	}
}

func (bp *BufferPool) Get() []byte {
	return bp.pool.Get().([]byte)
}

func (bp *BufferPool) Put(buf []byte) {
	bp.pool.Put(buf[:0])
}

func worker(pool *BufferPool, input <-chan []byte, output chan<- []byte, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range input {
		buf := pool.Get()
		buf = append(buf, data...)
		buf = append(buf, []byte(" processed")...)
		output <- buf
		pool.Put(buf)
	}
}

func main() {
	const (
		numWorkers = 4
		numJobs    = 10
	)

	// Initialize buffer pool with a typical buffer size
	bufferSize := 64
	bufferPool := NewBufferPool(bufferSize)

	input := make(chan []byte, numJobs)
	output := make(chan []byte, numJobs)
	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(bufferPool, input, output, &wg)
	}

	// Start a goroutine to close the output channel once all workers are done
	go func() {
		wg.Wait()
		close(output)
	}()

	// Data generator
	go func() {
		for i := 0; i < numJobs; i++ {
			data := []byte(fmt.Sprintf("Some data %d", i))
			input <- data
		}
		close(input) // Close the input channel to signal no more data
	}()

	// Data reader
	for processedData := range output {
		fmt.Println(string(processedData))
	}
}
