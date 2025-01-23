package main

import (
	"fmt"
)

type Dataset struct {
	values           []float64
	normalizedValues []float64
}

func (d *Dataset) Normalize() {
	min := float64(1.0)
	max := float64(0.0)
	for _, value := range d.values {
		if value < min {
			min = value
		} else if value > max {
			max = value
		}
	}
	if max-min > 0 {
		for i, value := range d.values {
			d.normalizedValues[i] = (value - min) / (max - min)
		}
	}
}
func (d *Dataset) PrintNormalizedValues() {
	for _, value := range d.normalizedValues {
		fmt.Printf("%f ", value)
	}
	fmt.Println()
}
func main() {
	// Create a Dataset
	dataset := Dataset{}

	// Add some data to the dataset
	dataset.values = []float64{10, 20, 30, 40, 50}

	// Normalize the dataset
	dataset.Normalize()

	// Print the normalized values
	dataset.PrintNormalizedValues() // Output will be 0.000000 0.250000 0.500000 0.750000 1.000000

	// Free up memory (lifecycle management)
	dataset.normalizedValues = nil
}
