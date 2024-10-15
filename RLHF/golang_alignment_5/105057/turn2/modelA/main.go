package main

import (
	"fmt"
	"sync/atomic"
	"time"
	"unsafe"
)

// HighPerformanceRingBuffer is a high-performance ring buffer designed for concurrent use.
type HighPerformanceRingBuffer struct {
	data     []byte // Use byte slice for raw memory access
	writePos uint64
	readPos  uint64
	mask     uint64
	size     int
}

// NewHighPerformanceRingBuffer creates a new high-performance ring buffer.
func NewHighPerformanceRingBuffer(size int) *HighPerformanceRingBuffer {
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

	rb := &HighPerformanceRingBuffer{
		data: make([]byte, size),
		mask: uint64(size - 1),
		size: size,
	}

	return rb
}

// Capacity returns the maximum number of events the ring buffer can hold.
func (rb *HighPerformanceRingBuffer) Capacity() int {
	return rb.size
}

// WriteEvent writes an event into the ring buffer. Returns true if successful, false if the buffer is full.
func (rb *HighPerformanceRingBuffer) WriteEvent(event []byte) bool {
	eventSize := len(event)
	if eventSize == 0 {
		return false // Ignore empty events
	}

	nextWritePos := (rb.writePos + uint64(eventSize)) & rb.mask
	if nextWritePos <= rb.readPos {
		// Buffer is full if next write position would overwrite read position
		return false
	}

	// Copy the event size into the buffer at the write position
	writeIndex := int(rb.writePos) % rb.size
	if writeIndex+4 <= rb.size {
		copy(rb.data[writeIndex:], event[:eventSize])
		*(*uint32)(unsafe.Pointer(&rb.data[writeIndex])) = uint32(eventSize) // Store event size
	} else {
		// Handle wrap around
		copy(rb.data[writeIndex:], event[:rb.size-writeIndex])
		copy(rb.data[:eventSize-(rb.size-writeIndex)], event[rb.size-writeIndex:])
		*(*uint32)(unsafe.Pointer(&rb.data[writeIndex])) = uint32(eventSize) // Store event size
	}

	atomic.StoreUint64(&rb.writePos, nextWritePos)
	return true
}

// ReadEvent reads an event from the ring buffer. Returns the event data and true if an event was read, false if the buffer is empty.
func (rb *HighPerformanceRingBuffer) ReadEvent() ([]byte, bool) {
	if atomic.LoadUint64(&rb.writePos) == rb.readPos {
		// Buffer is empty
		return nil, false
	}

	// Read the event size first
	readIndex := int(rb.readPos) % rb.size
	eventSize := int(*(*uint32)(unsafe.Pointer(&rb.data[readIndex])))
	if eventSize <= 0 {
		return nil, false // Corrupted data, return false
	}

	nextReadPos := (rb.readPos + uint64(eventSize)) & rb.mask

	// Copy the event data into a slice for return
	eventData := make([]byte, eventSize)
	if readIndex+4 <= rb.size {
		copy(eventData, rb.data[readIndex:readIndex+eventSize])
	} else {
		// Handle wrap around
		copy(eventData, rb.data[readIndex:])
		copy(eventData[eventSize-readIndex:], rb.data[:eventSize-(rb.size-readIndex)])
	}

	atomic.StoreUint64(&rb.readPos, nextReadPos)
	return eventData, true
}

func main() {
	const numEvents = 1000
	eventSizes := []int{10, 20, 30, 40, 50} // Variable event sizes for simulation
	events := make([][]byte, numEvents)
	for i := range events {
		eventSize := eventSizes[i%len(eventSizes)]
		events[i] = make([]byte, eventSize+4) // +4 for storing event size
		*(*uint32)(unsafe.Pointer(&events[i][0])) = uint32(eventSize)
	}

	bufferSize := 1024 * 1024 // 1MB buffer
	ringBuffer := NewHighPerformanceRingBuffer(bufferSize)

	// Producer goroutine
	go func() {
		for i := 0; i < numEvents; i++ {
			event := events[i]
			for !ringBuffer.WriteEvent(event) {
				// Buffer is full, wait for some space
				time.Sleep(time.Microsecond)
			}
		}
	}()

	// Consumer goroutine
	go func() {
		for i := 0; i < numEvents; i++ {
			eventData, ok := ringBuffer.ReadEvent()
			if !ok {
				// Buffer is empty, wait for some events
				time.Sleep(time.Microsecond)
				continue
			}
			// Process the event data
			_ = eventData
		}
	}()

	// Let the goroutines run for a while to process all events
	time.Sleep(time.Second)
	fmt.Println("Ring buffer test completed.")
}
