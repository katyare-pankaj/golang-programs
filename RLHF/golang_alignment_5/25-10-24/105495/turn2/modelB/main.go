package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// TemperatureController represents a safety-critical temperature control system.
type TemperatureController struct {
	currentTemperature float64
	targetTemperature  float64
	hysteresis         float64
}

// NewTemperatureController creates a new TemperatureController with the given parameters.
func NewTemperatureController(currentTemperature, targetTemperature, hysteresis float64) *TemperatureController {
	return &TemperatureController{
		currentTemperature: currentTemperature,
		targetTemperature:  targetTemperature,
		hysteresis:         hysteresis,
	}
}

// Control updates the temperature based on the current and target temperatures.
func (c *TemperatureController) Control() {
	for {
		// Simulate random temperature changes in the room
		c.currentTemperature += math.Round(rand.NormFloat64() * 0.5)

		// Calculate the temperature deviation from the target
		temperatureDeviation := c.currentTemperature - c.targetTemperature

		// Branch fold: Perform a single conditional check and handle all cases together
		action := ""
		if temperatureDeviation > c.hysteresis {
			action = "Turning off the heater."
		} else if temperatureDeviation < -c.hysteresis {
			action = "Turning on the heater."
		} else {
			action = "Doing nothing."
		}
		fmt.Println(action)
	}
}

func main() {
	// Initialize the random number generator
	rand.Seed(time.Now().UnixNano())

	// Create a new TemperatureController with the given parameters
	controller := NewTemperatureController(22.0, 23.0, 0.5)

	// Start the control loop
	go controller.Control()

	// Simulate external events
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		fmt.Println("External event: Temperature changed by", rand.NormFloat64()*2.0)
	}
}
