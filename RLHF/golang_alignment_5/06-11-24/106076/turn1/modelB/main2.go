// TemperatureController component
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

type TemperatureController struct {
	currentTemperature float64
	targetTemperature  float64
	heatingEnabled     bool
}

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

func (tc *TemperatureController) run() {
	for {
		// Receive temperature reading from sensor
		temperatureReading := readTemperatureSensor()

		// Update internal state
		tc.updateTemperature(temperatureReading)

		// Check if the temperature needs to be adjusted
		if tc.currentTemperature != tc.targetTemperature {
			// Adjust the heating/cooling unit
			adjustHeatingCoolingUnit(tc.targetTemperature)
		}

		// Wait for a short interval before re-checking the temperature
		time.Sleep(time.Second)
	}
}

func (tc *TemperatureController) updateTemperature(temperature float64) {
	tc.currentTemperature = temperature
	tc.heatingEnabled = temperature < tc.targetTemperature
}

// Heating/Cooling Unit component
func heatingCoolingUnit(heatingEnabled *bool) {
	for {
		// Check if heating is enabled
		if *heatingEnabled {
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

func main() {
	// Initialize the target temperature
	targetTemperature := 22.0

	// Create a TemperatureController instance
	tc := &TemperatureController{targetTemperature: targetTemperature}

	// Start the temperature controller and heating/cooling unit as goroutines
	go tc.run()
	go heatingCoolingUnit(&tc.heatingEnabled)

	// Main thread waits for user input to exit
	fmt.Scanln()
}

// Function to turn on the heating element
func turnOnHeating() {
	fmt.Println("Heating turned on.")
}

// Function to turn off the heating element
func turnOffHeating() {
	fmt.Println("Heating turned off.")
}
