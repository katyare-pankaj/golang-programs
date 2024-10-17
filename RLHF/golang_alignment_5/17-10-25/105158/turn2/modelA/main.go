package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Define an error type
type AppError struct {
	msg string
}

func (e *AppError) Error() string {
	return e.msg
}

// Continuation passing style function to handle sensor data processing
func processSensorDataCPS(data []byte, cont func(error)) {
	// Simulate some work to process the data
	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		// Check for data validity
		if len(data) == 0 {
			cont(&AppError{msg: "Invalid data received"})
			return
		}
		fmt.Println("Processed sensor data:", string(data))
		cont(nil)
	}()
}

// Using sync.Pool for byte slice reusing
var dataPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 1024) // Adjust size as needed for typical sensor data packet
	},
}

func main() {
	// Asynchronous usage example with CPS for real-time sensor data handling
	for {
		// Get a buffer from the pool
		buf := dataPool.Get().([]byte)

		// Simulate reading sensor data
		n := rand.Intn(len(buf))
		for i := 0; i < n; i++ {
			buf[i] = byte(rand.Intn(256))
		}
		sensorData := buf[:n]

		processSensorDataCPS(sensorData, func(err error) {
			if err != nil {
				fmt.Println("Error processing data:", err)
			}
			// Return the buffer to the pool for reuse
			dataPool.Put(buf)
		})

		// Simulate real-time processing with a short delay
		time.Sleep(time.Second)
	}
}
