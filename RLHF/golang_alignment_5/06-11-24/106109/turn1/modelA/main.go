package main

import (
	"fmt"
	"math/rand"
)

// Drone interface defines the basic operations that a drone should support
type Drone interface {
	TakeOff()
	Land()
	Fly(distance float64)
	GetPosition() (float64, float64)
}

// BasicDrone is a simple implementation of the Drone interface
type BasicDrone struct {
	positionX, positionY float64
}

func (d *BasicDrone) TakeOff() {
	fmt.Println("Drone taking off...")
}

func (d *BasicDrone) Land() {
	fmt.Println("Drone landing...")
}

func (d *BasicDrone) Fly(distance float64) {
	d.positionX += distance
	fmt.Printf("Drone flying %.2f meters forward. New position: (%.2f, %.2f)\n", distance, d.positionX, d.positionY)
}

func (d *BasicDrone) GetPosition() (float64, float64) {
	return d.positionX, d.positionY
}

// DronePlugin represents a custom plugin that can extend the drone's behavior
type DronePlugin interface {
	Init(drone Drone)
	Execute()
}

// ExamplePlugin is a sample plugin that adds the ability to perform random turns
type ExamplePlugin struct {
	drone Drone
}

func (p *ExamplePlugin) Init(drone Drone) {
	p.drone = drone
}

func (p *ExamplePlugin) Execute() {
	turnAngle := rand.Float64() * 360
	fmt.Printf("Drone executing a random turn of %.2f degrees\n", turnAngle)
}

func main() {
	// Create a basic drone
	drone := &BasicDrone{}

	// Add a custom plugin
	plugin := &ExamplePlugin{}
	plugin.Init(drone)

	// Extensible drone simulation
	drone.TakeOff()
	drone.Fly(10.0)
	plugin.Execute() // Perform random turn using the plugin
	drone.Fly(20.0)
	drone.Land()

	fmt.Println("Drone mission completed.")
}
