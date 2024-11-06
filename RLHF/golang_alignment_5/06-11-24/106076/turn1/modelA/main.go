package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Represents a single thermostat in the climate control system
type Thermostat struct {
	id          int
	targetTemp  float64
	currentTemp float64
	mutex       sync.Mutex
}

// Simulate adjusting the temperature of a thermostat
func (t *Thermostat) adjustTemperature() {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	// Simulate temperature adjustment logic
	delta := rand.Float64()*0.5 - 0.25
	t.currentTemp += delta
}

// Global function to simulate controlling multiple thermostats
func controlThermostats(thermostats []*Thermostat) {
	for _, thermostat := range thermostats {
		go func(t *Thermostat) {
			for {
				t.adjustTemperature()
				time.Sleep(time.Second * 1) // Simulate adjustment interval
			}
		}(thermostat)
	}
}

func main() {
	const numThermostats = 10
	thermostats := make([]*Thermostat, 0, numThermostats)

	// Initialize thermostats
	for i := 0; i < numThermostats; i++ {
		thermostats = append(thermostats, &Thermostat{
			id:          i,
			targetTemp:  22.0,
			currentTemp: 20.0,
		})
	}

	// Control thermostats concurrently
	controlThermostats(thermostats)

	// Simulate system running for a while
	time.Sleep(time.Second * 5)

	// Print final temperatures
	for _, thermostat := range thermostats {
		fmt.Printf("Thermostat %d: Current Temp: %.2f\n", thermostat.id, thermostat.currentTemp)
	}
}
