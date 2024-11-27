package main

import (
	"fmt"
	"hash/crc32"
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

func calculateChecksum(frame interface{}) uint16 {
	v := reflect.ValueOf(frame)
	t := reflect.TypeOf(frame)
	numFields := t.NumField()
	var data []byte

	// Iterate through all the fields of the struct and append their values to the data slice
	for i := 0; i < numFields; i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		// Skip the Options and Data fields as they are byte slices and we need to handle them differently
		if fieldType.Name == "Options" || fieldType.Name == "Data" {
			continue
		}

		// Convert the field value to bytes and append it to the data slice
		switch field.Kind() {
		case reflect.Uint8, reflect.Uint16, reflect.Uint32:
			value := uint32(field.Uint())
			data = append(data, byte(value>>8), byte(value))
		default:
			panic(fmt.Sprintf("Unsupported field type: %v", field.Kind()))
		}
	}

	// Handle the Options and Data fields separately
	optionsField := v.FieldByName("Options")
	if optionsField.IsValid() {
		data = append(data, optionsField.Bytes()...)
	}

	dataField := v.FieldByName("Data")
	if dataField.IsValid() {
		data = append(data, dataField.Bytes()...)
	}

	// Calculate the checksum using the crc32 package
	checksum := crc32.ChecksumIEEE(data)

	// Return the checksum as a uint16
	return uint16(checksum)
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

	// Calculate the checksum
	checksum := calculateChecksum(frame)

	// Update the Checksum field in the frame
	frame.Checksum = checksum

	// Inspect the frame using reflection to verify the checksum
	v := reflect.ValueOf(frame)
	t := reflect.TypeOf(frame)
	numFields := t.NumField()
	for i := 0; i < numFields; i++ {
		field := v.Field(i)
		fieldType := t.Field(i)
		fmt.Printf("%s: %v\n", fieldType.Name, field.Interface())
	}
}
