package main

import (
	"fmt"
	"sync"
)

const bufSize = 1 << 20 // 1 MB

type DataPipeline struct {
	input  chan []byte
	output chan []byte
	bw     chan []byte
	wg     sync.WaitGroup
}

func NewDataPipeline(workers int) *DataPipeline {
	p := &DataPipeline{
		input:  make(chan []byte),
		output: make(chan []byte),
		bw:     make(chan []byte, workers),
	}
	p.wg.Add(workers)
	for i := 0; i < workers; i++ {
		go p.process()
	}
	return p
}

func (p *DataPipeline) Add(data []byte) {
	p.input <- data
}

func (p *DataPipeline) Read() []byte {
	return <-p.output
}

func (p *DataPipeline) Close() {
	close(p.input)
	p.wg.Wait()
	close(p.output)
}

func (p *DataPipeline) process() {
	defer p.wg.Done()
	for data := range p.input {
		// Wait for a free buffer from the pool
		buf := <-p.bw

		// Modify the data in place in the buffer
		copy(buf, data)
		buf = append(buf, []byte(" processed")...)

		// Write the modified data back to the output channel
		p.output <- buf

		// Return the buffer to the pool
		p.bw <- buf
	}
}

func main() {
	pipeline := NewDataPipeline(4)

	// Data generator
	go func() {
		for i := 0; i < 1000; i++ {
			data := fmt.Sprintf("Some data %d", i)
			pipeline.Add([]byte(data))
		}
		close(pipeline.input)
	}()

	// Data reader
	go func() {
		for data := range pipeline.output {
			// Process the data
			//fmt.Println(string(data))

			// Return the buffer to the pool
			pipeline.bw <- data
		}
	}()

	pipeline.Close()
}
