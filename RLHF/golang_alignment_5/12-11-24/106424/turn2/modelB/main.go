package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
	"unsafe"
)

// Simple struct to serialize
type Data struct {
	A uint8
	B uint16
	C uint32
}

const (
	dataSize     = 1024 * 1024 // 1 MB of data
	unrollFactor = 8           // Adjust this based on your architecture
)

// Serializes data into a byte array using loop unrolling
func serializeData(data []Data, buf []byte) {
	dataPtr := unsafe.Pointer(&data[0])
	bufPtr := unsafe.Pointer(&buf[0])

	for i := 0; i < len(buf); i += unrollFactor * 8 {
		d0 := *(*uint64)(unsafe.Add(dataPtr, uintptr(i)))
		d1 := *(*uint64)(unsafe.Add(dataPtr, uintptr(i+8)))
		d2 := *(*uint64)(unsafe.Add(dataPtr, uintptr(i+16)))
		d3 := *(*uint64)(unsafe.Add(dataPtr, uintptr(i+24)))
		d4 := *(*uint64)(unsafe.Add(dataPtr, uintptr(i+32)))
		d5 := *(*uint64)(unsafe.Add(dataPtr, uintptr(i+40)))
		d6 := *(*uint64)(unsafe.Add(dataPtr, uintptr(i+48)))
		d7 := *(*uint64)(unsafe.Add(dataPtr, uintptr(i+56)))

		*(*uint64)(unsafe.Add(bufPtr, uintptr(i))) = d0
		*(*uint64)(unsafe.Add(bufPtr, uintptr(i+8))) = d1
		*(*uint64)(unsafe.Add(bufPtr, uintptr(i+16))) = d2
		*(*uint64)(unsafe.Add(bufPtr, uintptr(i+24))) = d3
		*(*uint64)(unsafe.Add(bufPtr, uintptr(i+32))) = d4
		*(*uint64)(unsafe.Add(bufPtr, uintptr(i+40))) = d5
		*(*uint64)(unsafe.Add(bufPtr, uintptr(i+48))) = d6
		*(*uint64)(unsafe.Add(bufPtr, uintptr(i+56))) = d7
	}
}

// Deserializes data from a byte array into a slice of Data using loop unrolling
func deserializeData(buf []byte, data []Data) {
	dataPtr := unsafe.Pointer(&data[0])
	bufPtr := unsafe.Pointer(&buf[0])

	for i := 0; i < len(buf); i += unrollFactor * 8 {
		d0 := *(*uint64)(unsafe.Add(bufPtr, uintptr(i)))
		d1 := *(*uint64)(unsafe.Add(bufPtr, uintptr(i+8)))
		d2 := *(*uint64)(unsafe.Add(bufPtr, uintptr(i+16)))
		d3 := *(*uint64)(unsafe.Add(bufPtr, uintptr(i+24)))
		d4 := *(*uint64)(unsafe.Add(bufPtr, uintptr(i+32)))
		d5 := *(*uint64)(unsafe.Add(bufPtr, uintptr(i+40)))
		d6 := *(*uint64)(unsafe.Add(bufPtr, uintptr(i+48)))
		d7 := *(*uint64)(unsafe.Add(bufPtr, uintptr(i+56)))

		*(*uint64)(unsafe.Add(dataPtr, uintptr(i))) = d0
		*(*uint64)(unsafe.Add(dataPtr, uintptr(i+8))) = d1
		*(*uint64)(unsafe.Add(dataPtr, uintptr(i+16))) = d2
		*(*uint64)(unsafe.Add(dataPtr, uintptr(i+24))) = d3
		*(*uint64)(unsafe.Add(dataPtr, uintptr(i+32))) = d4
		*(*uint64)(unsafe.Add(dataPtr, uintptr(i+40))) = d5
		*(*uint64)(unsafe.Add(dataPtr, uintptr(i+48))) = d6
		*(*uint64)(unsafe.Add(dataPtr, uintptr(i+56))) = d7
	}
}

func main() {
	// Increase GOMAXPROCS to utilize all cores
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Create large data and buffer slices
	data := make([]Data, dataSize)
	buf := make([]byte, dataSize*8) // sizeof(Data) = 8 bytes

	// Initialize data with some value
	for i := range data {
		data[i] = Data{A: uint8(i), B: uint16(i * 2), C: uint32(i * 3)}
	}

	// Benchmark serialization and deserialization
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		start := time.Now()
		serializeData(data, buf)
		fmt.Println("Serialization time:", time.Since(start))
	}()

	go func() {
		defer wg.Done()
		start := time.Now()
		deserializeData(buf, data)
		fmt.Println("Deserialization time:", time.Since(start))
	}()

	wg.Wait()
}
