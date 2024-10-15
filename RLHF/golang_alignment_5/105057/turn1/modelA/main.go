package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// RingBuffer is a fixed-size lock-free ring buffer for asynchronous events.
type RingBuffer struct {
	data     []interface{}
	writePos uint64
	readPos  uint64
	mask     uint64
	mu       sync.Mutex // For resizing
}

// NewRingBuffer creates a new RingBuffer with the specified size.
func NewRingBuffer(size int) *RingBuffer {
	if size <= 0 {
		panic("Ring buffer size must be positive")
	}

	// Calculate the next power of 2 for efficient masking
	size--
	size |= size >> 1
	size |= size >> 2
	size |= size >> 4
	size |= size >> 8
	size |= size >> 16
	size |= size >> 32
	size++

	rb := &RingBuffer{
		data: make([]interface{}, size),
		mask: uint64(size - 1),
	}
	return rb
}

// Capacity returns the maximum number of events the ring buffer can hold.
func (rb *RingBuffer) Capacity() int {
	return len(rb.data)
}

// WriteEvent writes an event into the ring buffer. Returns true if successful, false if the buffer is full.
func (rb *RingBuffer) WriteEvent(event interface{}) bool {
	nextWritePos := (rb.writePos + 1) & rb.mask
	if nextWritePos == atomic.LoadUint64(&rb.readPos) {
		// Buffer is full
		return false
	}

	rb.data[rb.writePos&rb.mask] = event
	atomic.StoreUint64(&rb.writePos, nextWritePos)
	return true
}

// ReadEvent reads an event from the ring buffer. Returns true if an event was read, false if the buffer is empty.
func (rb *RingBuffer) ReadEvent() (interface{}, bool) {
	if atomic.LoadUint64(&rb.writePos) == rb.readPos {
		// Buffer is empty
		return nil, false
	}

	event := rb.data[rb.readPos&rb.mask]
	atomic.StoreUint64(&rb.readPos, (rb.readPos+1)&rb.mask)
	return event, true
}

// Resize allows resizing the ring buffer (not thread-safe while reading/writing).
func (rb *RingBuffer) Resize(newSize int) {
	rb.mu.Lock()
	defer rb.mu.Unlock()

	if newSize <= 0 {
		panic("Ring buffer size must be positive")
	}

	newData := make([]interface{}, newSize)

	// Copy existing data up to the read pointer
	copyLength := int(rb.readPos) % len(rb.data)
	copy(newData, rb.data[:copyLength])

	// Copy data from the start of the buffer if read pointer has wrapped
	if copyLength < len(rb.data) {
		copy(newData[copyLength:], rb.data[copyLength:])
	}

	rb.data = newData
	rb.mask = uint64(newSize - 1)
	rb.writePos = uint64(copyLength)
}

func main() {
	bufferSize := 10
	ringBuffer := NewRingBuffer(bufferSize)

	// Producer goroutine: Simulate asynchronous events writing to the ring buffer
	go func() {
		for i := 0; i < 20; i++ {
			event := fmt.Sprintf("Event %d", i)
			for !ringBuffer.WriteEvent(event) {
				// Buffer is full, wait for some space
			}
			fmt.Println("Produced:", event)
		}
	}()

	// Consumer goroutine: Simulate processing asynchronous events from the ring buffer
	go func() {
		for {
			event, ok := ringBuffer.ReadEvent()
			if !ok {
				// Buffer is empty, wait for some events
				continue
			}
			fmt.Println("Consumed:", event)
		}
	}()

	// Let the goroutines run for a while
	select {}
}
