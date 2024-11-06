package main

import (
	"fmt"
	"math/rand"
)

// DroneNavigationSystem interface defines the basic navigation operations
type DroneNavigationSystem interface {
	PlanPath(start, end Point) []Point
	GetCurrentPosition() Point
}

// BasicNavigationSystem is a simple implementation of the DroneNavigationSystem interface
type BasicNavigationSystem struct {
	currentPosition Point
}

func (n *BasicNavigationSystem) PlanPath(start, end Point) []Point {
	// Simple linear path planning
	path := []Point{}
	for i := 0; i < 10; i++ {
		x := start.X + (end.X-start.X)*float64(i)/10
		y := start.Y + (end.Y-start.Y)*float64(i)/10
		path = append(path, Point{X: x, Y: y})
	}
	return path
}

func (n *BasicNavigationSystem) GetCurrentPosition() Point {
	return n.currentPosition
}

// Point represents a 2D position
type Point struct {
	X, Y float64
}

// PathPlannerPlugin represents a custom plugin that can extend path planning
type PathPlannerPlugin interface {
	Init(navSystem DroneNavigationSystem)
	PlanPath(start, end Point) []Point
}

// AStarPlugin implements A* path planning algorithm
type AStarPlugin struct {
	navSystem DroneNavigationSystem
}

func (p *AStarPlugin) Init(navSystem DroneNavigationSystem) {
	p.navSystem = navSystem
}

func (p *AStarPlugin) PlanPath(start, end Point) []Point {
	// Perform A* path planning here (for demonstration, we'll just use a random path)
	path := []Point{}
	for i := 0; i < 10; i++ {
		x := start.X + rand.Float64()*10
		y := start.Y + rand.Float64()*10
		path = append(path, Point{X: x, Y: y})
	}
	fmt.Println("Using A* Path Planner")
	return path
}

func main() {
	// Create a basic navigation system
	navSystem := &BasicNavigationSystem{}

	// Add an A* path planning plugin
	aStarPlugin := &AStarPlugin{}
	aStarPlugin.Init(navSystem)

	// Extensible navigation system
	start := Point{X: 0, Y: 0}
	end := Point{X: 100, Y: 100}

	fmt.Println("Basic Path Planning:")
	path := navSystem.PlanPath(start, end)
	for _, p := range path {
		fmt.Println(p)
	}

	fmt.Println("\nA* Path Planning using plugin:")
	// Use the A* path planner plugin
	path = aStarPlugin.PlanPath(start, end)
	for _, p := range path {
		fmt.Println(p)
	}
}
