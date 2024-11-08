package main

import (
	"math"
	"testing"
)

// AreaCalculator takes a radius and returns the area
type AreaCalculator func(float64) float64

// TestCalculateArea is the test function that is recognized by the testing framework
func TestCalculateArea(t *testing.T) {
	tests := []struct {
		radius   float64
		expected float64
		testName string
	}{
		{radius: 5.0, expected: 78.54, testName: "Test Circle with Radius 5"},
		{radius: 0.0, expected: 0.0, testName: "Test Circle with Radius 0"},
		// Add more test cases as needed
	}

	// Function to calculate the area of a circle (pure function)
	calculateArea := func(radius float64) float64 {
		return math.Pi * radius * radius
	}

	// Loop through each test case
	for _, test := range tests {
		actual := calculateArea(test.radius)

		// Use a small delta when comparing floating point numbers
		if math.Abs(actual-test.expected) > 0.01 {
			t.Errorf("%s: got %f, expected %f", test.testName, actual, test.expected)
		}
	}
}
