// Global state (shared data)

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	currentTemperature float64
	targetTemperature  float64
	heatingEnabled     bool
)

// Temperature Controller component
func temperatureController() {
	for {
		// Receive temperature reading from sensor
		temperatureReading := readTemperatureSensor()

		// Check if the temperature needs to be adjusted
		if temperatureReading != targetTemperature {
			// Adjust the heating/cooling unit
			adjustHeatingCoolingUnit(temperatureReading)
		}

		// Wait for a short interval before re-checking the temperature
		time.Sleep(time.Second)
	}
}

// Heating/Cooling Unit component
func heatingCoolingUnit() {
	for {
		// Check if heating is enabled
		if heatingEnabled {
			// Turn on the heating element
			turnOnHeating()
		} else {
			// Turn off the heating element
			turnOffHeating()
		}

		// Wait for a short interval before re-checking the status
		time.Sleep(time.Second)
	}
}

// Function to read temperature from the sensor
func readTemperatureSensor() float64 {
	// Simulate temperature reading
	return rand.Float64()*30 + 15
}

// Function to adjust the heating/cooling unit based on the current temperature
func adjustHeatingCoolingUnit(temperatureReading float64) {
	// Adjust the heatingEnabled flag based on the target temperature
	heatingEnabled = temperatureReading < targetTemperature
}

// Function to turn on the heating element
func turnOnHeating() {
	fmt.Println("Heating turned on.")
}

// Function to turn off the heating element
func turnOffHeating() {
	fmt.Println("Heating turned off.")
}

func main() {
	// Initialize the target temperature
	targetTemperature = 22.0

	// Start the temperature controller and heating/cooling unit as goroutines
	go temperatureController()
	go heatingCoolingUnit()

	// Main thread waits for user input to exit
	fmt.Scanln()
}
