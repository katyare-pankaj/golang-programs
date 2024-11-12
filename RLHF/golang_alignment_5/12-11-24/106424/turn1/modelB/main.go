package main

import (
	"fmt"
	"unsafe"
)

// Simple struct to serialize
type Data struct {
	A uint8
	B uint16
	C uint32
}

// Serializes data into a byte array using loop unrolling
func serializeData(data *Data, buf []byte) {
	// Unsafe pointer to the data struct
	dataPtr := unsafe.Pointer(data)

	// Unsafe pointer to the byte array
	bufPtr := unsafe.Pointer(&buf[0])

	// Loop unrolling factor (adjust this based on your architecture)
	const unrollFactor = 4

	for i := 0; i < len(buf); i += unrollFactor * 8 {
		// Load data from the struct into registers
		d0 := *(*uint64)(unsafe.Add(dataPtr, uintptr(i)))
		d1 := *(*uint64)(unsafe.Add(dataPtr, uintptr(i+8)))
		d2 := *(*uint64)(unsafe.Add(dataPtr, uintptr(i+16)))
		d3 := *(*uint64)(unsafe.Add(dataPtr, uintptr(i+24)))

		// Store data into the byte array
		*(*uint64)(unsafe.Add(bufPtr, uintptr(i))) = d0
		*(*uint64)(unsafe.Add(bufPtr, uintptr(i+8))) = d1
		*(*uint64)(unsafe.Add(bufPtr, uintptr(i+16))) = d2
		*(*uint64)(unsafe.Add(bufPtr, uintptr(i+24))) = d3
	}
}

func main() {
	data := &Data{A: 1, B: 2, C: 3}
	buf := make([]byte, 8) // sizeof(Data) = 8 bytes

	// Serialize the data
	serializeData(data, buf)

	// Display the serialized data
	fmt.Printf("%x %x %x %x\n", buf[0], buf[1], buf[2], buf[3])
	fmt.Printf("%x %x %x %x\n", buf[4], buf[5], buf[6], buf[7])
}
