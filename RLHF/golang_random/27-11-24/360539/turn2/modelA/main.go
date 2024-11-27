package main

import (
	"encoding/binary"
	"fmt"
	"reflect"
)

type TCPFrame struct {
	SourcePort           uint16
	DestinationPort      uint16
	SequenceNumber       uint32
	AcknowledgmentNumber uint32
	DataOffset           uint8
	Reserved             uint8
	Flags                uint8
	WindowSize           uint16
	Checksum             uint16
	UrgentPointer        uint16
	Options              []byte
	Data                 []byte
}

func serializeTCPFrame(frame interface{}) []byte {
	v := reflect.ValueOf(frame)
	t := reflect.TypeOf(frame)
	numFields := t.NumField()
	var data []byte

	for i := 0; i < numFields; i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		switch fieldType.Kind {
		case reflect.Uint8:
			data = append(data, byte(field.Uint()))
		case reflect.Uint16:
			data = append(data, binary.LittleEndian.Uint16([]byte(field.Uint())))
		case reflect.Uint32:
			data = append(data, binary.LittleEndian.Uint32([]byte(field.Uint())))
		case reflect.Slice:
			switch fieldType.Elem.Kind {
			case reflect.Uint8:
				data = append(data, field.Bytes()...)
			default:
				panic(fmt.Sprintf("Unsupported slice element type: %v", fieldType.Elem.Kind))
			}
		default:
			panic(fmt.Sprintf("Unsupported field type: %v", fieldType.Kind))
		}
	}

	return data
}
func inspectTCPFrame(frame interface{}) {
	v := reflect.ValueOf(frame)
	t := reflect.TypeOf(frame)
	numFields := t.NumField()

	fmt.Printf("TCP Frame Content:\n")
	for i := 0; i < numFields; i++ {
		field := v.Field(i)
		fieldType := t.Field(i)
		fmt.Printf("  %s: %v (%v)\n", fieldType.Name, field.Interface(), field.Kind())
	}
}
func main() {
	frame := TCPFrame{
		SourcePort:           1234,
		DestinationPort:      5678,
		SequenceNumber:       1000,
		AcknowledgmentNumber: 2000,
		DataOffset:           5,
		Reserved:             0,
		Flags:                0x02, // SYN flag
		WindowSize:           8192,
		Checksum:             0, // Initialize with 0 for calculation
		UrgentPointer:        0,
		Options:              []byte{0x01, 0x03, 0x03, 0x07}, // MSS option
		Data:                 []byte("Hello, World!"),
	}

	// Inspect the frame
	inspectTCPFrame(frame)

	// Serialize the frame into a byte slice
	serializedData := serializeTCPFrame(frame)
	fmt.Printf("Serialized TCP Frame (bytes): %v\n", serializedData)

	// You can now transmit the serializedData over the network
}
