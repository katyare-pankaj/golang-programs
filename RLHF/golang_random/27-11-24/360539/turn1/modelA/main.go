package main

import (
	"fmt"
	"reflect"
)

// Define a simple TCP frame structure
type TCPFrame struct {
	SourcePort           uint16
	DestinationPort      uint16
	SequenceNumber       uint32
	AcknowledgmentNumber uint32
	DataOffset           uint8
	Flags                uint8
	WindowSize           uint16
	Checksum             uint16
	UrgentPointer        uint16
}

func main() {
	// Create an instance of TCPFrame
	frame := TCPFrame{
		SourcePort:           12345,
		DestinationPort:      80,
		SequenceNumber:       1,
		AcknowledgmentNumber: 0,
		DataOffset:           5,
		Flags:                0x18, // SYN + ACK
		WindowSize:           65535,
		Checksum:             0,
		UrgentPointer:        0,
	}

	// Use reflection to inspect the TCPFrame
	reflectValue := reflect.ValueOf(frame)

	// Get the number of fields in the structure
	numFields := reflectValue.NumField()
	fmt.Printf("TCPFrame has %d fields.\n", numFields)

	// Iterate over each field
	for i := 0; i < numFields; i++ {
		field := reflectValue.Field(i)
		fieldName := reflectValue.Type().Field(i).Name
		fieldValue := field.Interface()

		// Print the field name and value
		fmt.Printf("Field %d: %s = %v\n", i, fieldName, fieldValue)
	}

	// Manipulate a field using reflection
	// Let's change the SourcePort to 8080
	sourcePortField := reflectValue.FieldByName("SourcePort")
	if sourcePortField.IsValid() {
		sourcePortField.SetUint(8080)
		fmt.Println("SourcePort changed to 8080.")
	} else {
		fmt.Println("SourcePort field not found.")
	}

	// Print the updated frame
	fmt.Println("Updated TCPFrame:")
	reflectValue = reflect.ValueOf(frame)
	numFields = reflectValue.NumField()
	for i := 0; i < numFields; i++ {
		field := reflectValue.Field(i)
		fieldName := reflectValue.Type().Field(i).Name
		fieldValue := field.Interface()
		fmt.Printf("Field %d: %s = %v\n", i, fieldName, fieldValue)
	}
}
