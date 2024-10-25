package main

import (
	"fmt"
	"math/rand"
	"time"
)

// TemperatureController represents a control system that regulates the temperature of a room.
type TemperatureController struct {
	currentTemperature float64
	targetTemperature  float64
	hysteresis         float64
}

// NewTemperatureController creates a new TemperatureController with the given initial temperature, target temperature, and hysteresis.
func NewTemperatureController(currentTemperature, targetTemperature, hysteresis float64) *TemperatureController {
	return &TemperatureController{
		currentTemperature: currentTemperature,
		targetTemperature:  targetTemperature,
		hysteresis:         hysteresis,
	}
}

// Control updates the temperature of the room based on the current temperature and the target temperature.
func (c *TemperatureController) Control() {
	for {
		// Simulate random temperature changes in the room
		c.currentTemperature += rand.NormFloat64() * 0.5

		// Check if the temperature is too high or too low
		if c.currentTemperature > c.targetTemperature+c.hysteresis {
			// Temperature is too high, turn off the heater
			fmt.Println("Temperature is too high, turning off the heater.")
		} else if c.currentTemperature < c.targetTemperature-c.hysteresis {
			// Temperature is too low, turn on the heater
			fmt.Println("Temperature is too low, turning on the heater.")
		} else {
			// Temperature is within the hysteresis range, do nothing
			fmt.Println("Temperature is within the hysteresis range, doing nothing.")
		}

		// Simulate a control cycle delay
		time.Sleep(time.Second)
	}
}

func main() {
	// Initialize the random number generator
	rand.Seed(time.Now().UnixNano())

	// Create a new TemperatureController with the given parameters
	controller := NewTemperatureController(22.0, 23.0, 0.5)

	// Start the control loop
	go controller.Control()

	// Simulate some external events that may affect the temperature
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		fmt.Println("External event: Temperature changed by", rand.NormFloat64()*2.0)
	}
}
