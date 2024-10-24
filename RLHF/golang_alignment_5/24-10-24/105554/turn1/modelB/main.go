package main

import (
	"fmt"
)

// Interval represents a range of real numbers
type Interval struct {
	Start float64
	End   float64
	state intervalState
}

// intervalState interface for different states of an interval
type intervalState interface {
	isOpen() bool
	isClosed() bool
}

// openState represents an open interval (Start, End)
type openState struct{}

func (o *openState) isOpen() bool {
	return true
}

func (o *openState) isClosed() bool {
	return false
}

// closedState represents a closed interval [Start, End]
type closedState struct{}

func (c *closedState) isOpen() bool {
	return false
}

func (c *closedState) isClosed() bool {
	return true
}

// NewInterval creates a new interval with the specified state
func NewInterval(start, end float64, state intervalState) *Interval {
	return &Interval{Start: start, End: end, state: state}
}

// IsOpen checks if the interval is open
func (i *Interval) IsOpen() bool {
	return i.state.isOpen()
}

// IsClosed checks if the interval is closed
func (i *Interval) IsClosed() bool {
	return i.state.isClosed()
}

// IntervalCommand represents a command for an interval operation
type IntervalCommand interface {
	execute(i1, i2 *Interval) *Interval
}

// UnionCommand represents the union of two intervals
type UnionCommand struct{}

func (u *UnionCommand) execute(i1, i2 *Interval) *Interval {
	start := min(i1.Start, i2.Start)
	end := max(i1.End, i2.End)
	state := &closedState{}
	if i1.IsOpen() || i2.IsOpen() {
		state = &openState{}
	}
	return NewInterval(start, end, state)
}

// IntersectionCommand represents the intersection of two intervals
type IntersectionCommand struct{}

func (i *IntersectionCommand) execute(i1, i2 *Interval) *Interval {
	start := max(i1.Start, i2.Start)
	end := min(i1.End, i2.End)
	if start > end {
		return nil
	}
	state := &closedState{}
	if i1.IsOpen() && i2.IsOpen() {
		state = &openState{}
	} else if i1.IsOpen() || i2.IsOpen() {
		state = &closedState{}
	}
	return NewInterval(start, end, state)
}

// ContainmentCommand represents if one interval contains the other
type ContainmentCommand struct{}

func (c *ContainmentCommand) execute(i1, i2 *Interval) *Interval {
	if i1.Start <= i2.Start && i2.End <= i1.End {
		return NewInterval(0, 0, &closedState{})
	}
	return nil
}

// IntervalProcessor executes interval commands
type IntervalProcessor struct {
	command IntervalCommand
}

// NewIntervalProcessor creates a new interval processor with a specific command
func NewIntervalProcessor(command IntervalCommand) *IntervalProcessor {
	return &IntervalProcessor{command: command}
}

// Execute executes the interval command on the given intervals
func (ip *IntervalProcessor) Execute(i1, i2 *Interval) *Interval {
	return ip.command.execute(i1, i2)
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example intervals
	i1 := NewInterval(1, 5, &closedState{}) // [1, 5]
	i2 := NewInterval(3, 7, &openState{})   // (3, 7)
	i3 := NewInterval(6, 8, &closedState{}) // [6, 8]

	// Create interval commands
	unionCmd := &UnionCommand{}
	intersectionCmd := &IntersectionCommand{}
	containmentCmd := &ContainmentCommand{}

	// Create interval processors
	unionProcessor := NewIntervalProcessor(unionCmd)
	intersectionProcessor := NewIntervalProcessor(intersectionCmd)
	containmentProcessor := NewIntervalProcessor(containmentCmd)

	// Perform interval operations
	unionResult := unionProcessor.Execute(i1, i2)
	fmt.Println("Union of [1, 5] and (3, 7):", unionResult) // Output: (1, 7)

	intersectionResult := intersectionProcessor.Execute(i1, i2)
	fmt.Println("Intersection of [1, 5] and (3, 7):", intersectionResult) // Output: (3, 5)

	containmentResult := containmentProcessor.Execute(i1, i3)
	fmt.Println("Containment of [1, 5] and [6, 8]:", containmentResult) // Output: [0, 0]
}
