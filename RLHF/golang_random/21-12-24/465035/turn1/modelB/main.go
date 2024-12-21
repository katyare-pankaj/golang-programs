package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// IoT Agent structure
type iotAgent struct {
	id         int
	sensorData float32
	wg         *sync.WaitGroup
}

// Function to simulate data reading from sensor
func (a *iotAgent) readSensorData() {
	defer a.wg.Done()
	// Simulate sensor data reading with a random delay
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	a.sensorData = rand.Float32() * 100
	fmt.Printf("Agent %d: Read sensor data = %.2f\n", a.id, a.sensorData)
}

// Function to process the sensor data
func (a *iotAgent) processData() {
	defer a.wg.Done()
	// Simulate data processing with a random delay
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	processedData := a.sensorData * 2
	fmt.Printf("Agent %d: Processed data = %.2f\n", a.id, processedData)
}

// Function to send the processed data to a central server
func (a *iotAgent) sendData() {
	defer a.wg.Done()
	// Simulate sending data with a random delay
	time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
	fmt.Printf("Agent %d: Data sent to the central server\n", a.id)
}

func main() {
	numAgents := 5
	wg := &sync.WaitGroup{}

	// Create a slice of IoT Agents
	agents := make([]*iotAgent, numAgents)

	for i := 0; i < numAgents; i++ {
		agents[i] = &iotAgent{
			id: i + 1,
			wg: wg,
		}
	}

	// Add number of goroutines to the WaitGroup
	wg.Add(3 * numAgents)

	// Start data reading, processing, and sending goroutines for each agent
	for _, agent := range agents {
		go agent.readSensorData()
		go agent.processData()
		go agent.sendData()
	}

	// Wait for all goroutines to complete
	wg.Wait()
	fmt.Println("All tasks completed. IoT agents are idle.")
}
