package main

import (
	"fmt"
	"unsafe"
)

// Data struct to serialize/deserialize
type Data struct {
	A uint8
	B uint16
	C uint32
}

const (
	dataSize     = int(unsafe.Sizeof(Data{})) // Size of the Data struct in bytes
	unrollFactor = 4                          // Adjust this based on your architecture
)

// Serializes data into a byte array using loop unrolling
func serializeData(data []Data, buf []byte) {
	for i := 0; i < len(data); i += unrollFactor {
		for j := 0; j < unrollFactor; j++ {
			idx := i + j
			if idx >= len(data) {
				break
			}
			d := data[idx]
			buf[idx*dataSize] = byte(d.A)
			buf[idx*dataSize+1] = byte(d.B >> 8)
			buf[idx*dataSize+2] = byte(d.B)
			buf[idx*dataSize+3] = byte(d.C >> 24)
			buf[idx*dataSize+4] = byte((d.C >> 16) & 0xFF)
			buf[idx*dataSize+5] = byte((d.C >> 8) & 0xFF)
			buf[idx*dataSize+6] = byte(d.C)
		}
	}
}

// Deserializes data from a byte array into a Data slice using loop unrolling
func deserializeData(buf []byte, data []Data) {
	for i := 0; i < len(data); i += unrollFactor {
		for j := 0; j < unrollFactor; j++ {
			idx := i + j
			if idx >= len(data) {
				break
			}
			d := &data[idx]
			d.A = uint8(buf[idx*dataSize])
			d.B = uint16(buf[idx*dataSize+1])<<8 | uint16(buf[idx*dataSize+2])
			d.C = uint32(buf[idx*dataSize+3])<<24 | uint32(buf[idx*dataSize+4])<<16 | uint32(buf[idx*dataSize+5])<<8 | uint32(buf[idx*dataSize+6])
		}
	}
}

func main() {
	const numDataElements = 1000
	data := make([]Data, numDataElements)
	buf := make([]byte, numDataElements*dataSize)

	// Initialize data for serialization
	for i := range data {
		data[i].A = uint8(i)
		data[i].B = uint16(i * 2)
		data[i].C = uint32(i * 3)
	}

	// Serialize the data
	serializeData(data, buf)

	// Deserialize data back into the slice
	deserializeData(buf, data)

	// Verify deserialization
	for i := range data {
		if data[i].A != uint8(i) || data[i].B != uint16(i*2) || data[i].C != uint32(i*3) {
			fmt.Println("Deserialization failed!")
			return
		}
	}
	fmt.Println("Serialization/Deserialization successful!")
}
