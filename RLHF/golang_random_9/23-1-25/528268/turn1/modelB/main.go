package main

import (
	"fmt"
	"reflect"
	"sync"
	"unsafe"
)

type DataPipeline struct {
	input  chan []byte
	output chan []byte
	wg     sync.WaitGroup
	bw     []byte
}

const (
	bufSize = 1 << 20 // 1 MB
)

func NewDataPipeline(workers int) *DataPipeline {
	p := &DataPipeline{
		input:  make(chan []byte),
		output: make(chan []byte),
		bw:     make([]byte, bufSize),
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
		// Slice the data using unsafe functions to avoid copying bytes
		// Assuming that `p.bw` is always large enough to hold the data
		start := uintptr(unsafe.Pointer(&p.bw[0]))
		end := start + uintptr(len(data))
		buf := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
			Data: start,
			Len:  int(end - start),
			Cap:  cap(p.bw),
		}))
		copy(buf, data)

		// Modify the data in place here
		// For simplicity, we just append some text
		buf = append(buf, []byte(" processed")...)

		// Write the modified data back to the output
		p.output <- buf
	}
}

func main() {
	pipeline := NewDataPipeline(4)

	// Data generator
	go func() {
		for i := 0; i < 100; i++ {
			data := fmt.Sprintf("Some data %d", i)
			pipeline.Add([]byte(data))
		}
	}()

	// Data reader
	for i := 0; i < 100; i++ {
		data := pipeline.Read()
		fmt.Println(string(data))
	}

	pipeline.Close()
}
