package main

import (
	"fmt"
)

// Interval represents a closed interval [min, max]
type Interval struct {
	Min float64
	Max float64
}

// Subject interface defines the methods for subjects (data streams)
type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObservers()
	GetInterval() Interval
}

// DataStream implements the Subject interface to represent a streaming data source
type DataStream struct {
	observers       []Observer
	currentInterval Interval
}

func (ds *DataStream) RegisterObserver(o Observer) {
	ds.observers = append(ds.observers, o)
}

func (ds *DataStream) RemoveObserver(o Observer) {
	// Implementation for removing observers
}

func (ds *DataStream) NotifyObservers() {
	for _, observer := range ds.observers {
		observer.Update(ds.currentInterval)
	}
}

func (ds *DataStream) GetInterval() Interval {
	return ds.currentInterval
}

// Observer interface defines the update method for observers
type Observer interface {
	Update(interval Interval)
}

// MovingAverageCalculator is an observer that calculates the moving average based on the data interval
type MovingAverageCalculator struct {
	windowSize     int
	currentAverage float64
}

func NewMovingAverageCalculator(windowSize int) *MovingAverageCalculator {
	return &MovingAverageCalculator{windowSize: windowSize}
}

func (mac *MovingAverageCalculator) Update(interval Interval) {
	// Calculate the moving average using the interval and window size
	minAverage := interval.Min / float64(mac.windowSize)
	maxAverage := interval.Max / float64(mac.windowSize)
	mac.currentAverage = (minAverage + maxAverage) / 2
}

func (mac *MovingAverageCalculator) GetCurrentAverage() float64 {
	return mac.currentAverage
}

func main() {
	// Set up the data stream and observers
	dataStream := &DataStream{}
	movingAverageCalculator := NewMovingAverageCalculator(5)

	dataStream.RegisterObserver(movingAverageCalculator)

	// Simulate streaming data and update the interval
	dataStream.currentInterval = Interval{Min: 20.0, Max: 80.0}
	dataStream.NotifyObservers()

	fmt.Printf("Current Moving Average: %.2f\n", movingAverageCalculator.GetCurrentAverage())

	// Simulate updating the data stream interval
	dataStream.currentInterval = Interval{Min: 40.0, Max: 100.0}
	dataStream.NotifyObservers()
	fmt.Printf("Current Moving Average: %.2f\n", movingAverageCalculator.GetCurrentAverage())
}
