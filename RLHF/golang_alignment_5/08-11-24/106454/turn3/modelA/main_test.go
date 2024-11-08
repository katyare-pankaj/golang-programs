package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateArea(t *testing.T) {
	assert.Equal(t, 3.14, calculateArea(1.0))
	assert.Equal(t, 12.57, calculateArea(2.0))
}

func TestMapFloat64(t *testing.T) {
	numbers := []float64{1, 2, 3, 4, 5}

	doubled := mapFloat64(numbers, func(x float64) float64 {
		return x * 2
	})
	assert.Equal(t, []float64{2, 4, 6, 8, 10}, doubled)

	squared := mapFloat64(numbers, func(x float64) float64 {
		return x * x
	})
	assert.Equal(t, []float64{1, 4, 9, 16, 25}, squared)
}

func TestCalculateAreas(t *testing.T) {
	radii := []float64{1.0, 2.0}
	expectedAreas := []float64{3.14, 12.57}
	assert.Equal(t, expectedAreas, calculateAreas(radii))
}
