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
	dataSize    = 1024 * 1024 // 1 MB of data
	unrollFactor = 8         // Adjust this based on your architecture
)

// Serializes data into a byte array using loop unrolling
func serializeData(data []Data, buf []byte) {
	dataPtr := unsafe.Pointer(&data[0])
	bufPtr := unsafe.Pointer(&buf[0])

	for i := 0; i < len(buf); i += unrollFactor * 8 {
		d0 := *(*uint64)(unsafe.Add(dataPtr, uintptr(i)))
		d1 := *(*uint64)(unsafe.Add(dataPtr, uintptr(i+8)))
		d2 := *(*uint64)(unsafe.Add(dataPtr, uintptr(i+16)))
		d3 := *(*uint