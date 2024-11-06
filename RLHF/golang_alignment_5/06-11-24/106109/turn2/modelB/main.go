package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Drone interface defines the basic operations that a drone should support
type Drone interface {
	SetAltitude(altitude float64)
	GetAltitude() float64
	AdjustAltitude(delta float64)
}

// BasicDrone is a simple implementation of the Drone interface
type BasicDrone struct {
	altitude float64
}

func (d *BasicDrone) SetAltitude(altitude float64) {
	d.altitude = altitude
}

func (d *BasicDrone) GetAltitude() float64 {
	return d.altitude
}

func (d *BasicDrone) AdjustAltitude(delta float64) {
	d.altitude += delta
}

// AltitudeControllerPlugin represents a custom plugin that can adjust the drone's altitude
type AltitudeControllerPlugin interface {
	Init(drone Drone)
	Execute()
}

// ExampleAltitudeControllerPlugin implements a simple altitude control algorithm with random adjustments
type ExampleAltitudeControllerPlugin struct {
	drone     Drone
	targetAlt float64
}

func (p *ExampleAltitudeControllerPlugin) Init(drone Drone) {
	p.drone = drone
	p.targetAlt = 10.0 // Initial target altitude
}

func (p *ExampleAltitudeControllerPlugin) Execute() {
	// Randomly adjust the target altitude within a range
	minAlt := p.targetAlt - 2.0
	maxAlt := p.targetAlt + 2.0
	p.targetAlt = rand.Float64()*(maxAlt-minAlt) + minAlt

	// Calculate the delta to adjust the altitude towards the target
	currentAlt := p.drone.GetAltitude()
	delta := (p.targetAlt - currentAlt) * 0.1 // 10% step size

	// Adjust the drone's altitude
	p.drone.AdjustAltitude(delta)

	fmt.Printf("Drone adjusting altitude towards %.2f meters. Current altitude: %.2f meters\n", p.targetAlt, currentAlt)
}

func main() {
	// Create a basic drone
	drone := &BasicDrone{}
	drone.SetAltitude(5.0) // Initial altitude

	// Add the altitude controller plugin
	plugin := &ExampleAltitudeControllerPlugin{}
	plugin.Init(drone)

	// Simulate real-time operations with the plugin
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		plugin.Execute() // Adjust the altitude based on the plugin's algorithm
	}

	fmt.Println("Drone altitude control simulation completed.")
}
