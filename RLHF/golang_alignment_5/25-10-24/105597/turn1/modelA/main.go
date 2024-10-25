package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
)

// BufferPool is a simple object pool for bytes.Buffers
type BufferPool struct {
	sync.Pool
}

// NewBufferPool creates a new BufferPool
func NewBufferPool() *BufferPool {
	return &BufferPool{
		Pool: sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	}
}

// GetBuffer retrieves a bytes.Buffer from the pool.
func (p *BufferPool) GetBuffer() *bytes.Buffer {
	return p.Get().(*bytes.Buffer)
}

// PutBuffer returns a bytes.Buffer to the pool.
func (p *BufferPool) PutBuffer(b *bytes.Buffer) {
	b.Reset()
	p.Put(b)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Get a buffer from the pool
	buf := bufferPool.GetBuffer()
	defer bufferPool.PutBuffer(buf)

	_, _ = buf.WriteString("Hello, World!")
	_, _ = w.Write(buf.Bytes())
}

var bufferPool *BufferPool

func main() {
	bufferPool = NewBufferPool()
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
