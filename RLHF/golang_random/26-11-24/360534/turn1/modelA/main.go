package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// Define a sample struct to be serialized and deserialized
type Data struct {
	Value     int    `json:"value"`
	Timestamp string `json:"timestamp"`
}

// Serialize a Data object into a JSON string
func serializeData(data Data) (string, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("error serializing data: %w", err)
	}
	return string(jsonBytes), nil
}

// Deserialize a JSON string into a Data object
func deserializeData(jsonString string) (*Data, error) {
	data := &Data{}
	err := json.Unmarshal([]byte(jsonString), data)
	if err != nil {
		return nil, fmt.Errorf("error deserializing data: %w", err)
	}
	return data, nil
}

// Goroutine to serialize data
func serializeGoroutine(serializeChan <-chan Data, resultChan chan<- string) {
	for {
		select {
		case data := <-serializeChan:
			jsonString, err := serializeData(data)
			if err != nil {
				resultChan <- fmt.Sprintf("Serialization error: %v", err)
			} else {
				resultChan <- jsonString
			}
		default:
			time.Sleep(1 * time.Millisecond)
		}
	}
}

// Goroutine to deserialize data
func deserializeGoroutine(deserializeChan <-chan string, resultChan chan<- Data) {
	for {
		select {
		case jsonString := <-deserializeChan:
			data, err := deserializeData(jsonString)
			if err != nil {
				resultChan <- Data{Value: -1, Timestamp: err.Error()}
			} else {
				resultChan <- *data
			}
		default:
			time.Sleep(1 * time.Millisecond)
		}
	}
}

func main() {
	// Create channels to synchronize serialization and deserialization
	serializeChan := make(chan Data)
	deserializeChan := make(chan string)
	resultChan := make(chan Data)

	// Start serialization and deserialization goroutines
	go serializeGoroutine(serializeChan, deserializeChan)
	go deserializeGoroutine(deserializeChan, resultChan)

	// Example usage
	for i := 0; i < 10; i++ {
		data := Data{Value: i, Timestamp: time.Now().String()}
		serializeChan <- data

		select {
		case result := <-resultChan:
			fmt.Printf("Deserialized Data: Value: %d, Timestamp: %s\n", result.Value, result.Timestamp)
		default:
			time.Sleep(1 * time.Millisecond)
		}
	}
}
