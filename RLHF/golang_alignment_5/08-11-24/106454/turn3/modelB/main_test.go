package main

import (
	"math"
	"testing"
)

// AreaCalculator takes a radius and returns the area
type AreaCalculator func(float64) float64

// TestAreaFunction tests the area calculator function with a custom assertion function
func TestAreaFunction(t *testing.T, areaCalculator AreaCalculator, assert func(float64, float64, string)) {
	tests := []struct {
		radius   float64
		expected float64
		testName string
	}{
		{radius: 5.0, expected: 78.54, testName: "Test Circle with Radius 5"},
		{radius: 0.0, expected: 0.0, testName: "Test Circle with Radius 0"},
		// Add more test cases
	}

	for _, test := range tests {
		actual := areaCalculator(test.radius)
		assert(actual, test.expected, test.testName)
	}
}

// Function to calculate the area of a circle (pure function)
func calculateArea(radius float64) float64 {
	return math.Pi * radius * radius
}

func TestCalculateArea(t *testing.T) {
	// Using the TestAreaFunction with an assertion function
	TestAreaFunction(t, calculateArea, func(actual, expected, testName string) {
		if actual != expected {
			t.Errorf("%s: got %f, expected %f", testName, actual, expected)
		}
	})
}
