package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

type Data struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func main() {
	var wg sync.WaitGroup
	dataChan := make(chan Data, 10)
	serializedChan := make(chan string, 10)
	deserializedChan := make(chan Data, 10)

	// Producer goroutine to generate Data and send it to dataChan
	go func() {
		for i := 0; i < 10; i++ {
			data := Data{
				Name:  fmt.Sprintf("Data %d", i),
				Value: i * i,
			}
			dataChan <- data
		}
		close(dataChan)
	}()

	// Serializer goroutine to receive Data from dataChan, serialize it, and send it to serializedChan
	go func() {
		for data := range dataChan {
			serializedData, _ := json.Marshal(data)
			serializedChan <- string(serializedData)
		}
		close(serializedChan)
	}()

	// Deserializer goroutine to receive serialized data from serializedChan, deserialize it, and send it to deserializedChan
	go func() {
		for serializedData := range serializedChan {
			var deserializedData Data
			json.Unmarshal([]byte(serializedData), &deserializedData)
			deserializedChan <- deserializedData
		}
		close(deserializedChan)
	}()

	// Consumer goroutine to receive deserialized Data from deserializedChan and print it
	go func() {
		for deserializedData := range deserializedChan {
			fmt.Printf("Deserialized Data: %v\n", deserializedData)
		}
		wg.Done()
	}()

	wg.Add(1)
	wg.Wait()
}
