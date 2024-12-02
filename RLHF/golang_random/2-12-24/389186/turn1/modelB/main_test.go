package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockFeatureSwitch simulates A/B testing feature toggling
func MockFeatureSwitch(feature string) string {
	if feature == "FeatureA" {
		return "A"
	}
	if feature == "FeatureB" {
		return "B"
	}
	return "default"
}

func TestFeatureSwitch(t *testing.T) {
	tests := []struct {
		name     string
		feature  string
		expected string
	}{
		{"Test A feature switch", "FeatureA", "A"},
		{"Test B feature switch", "FeatureB", "B"},
		{"Test default feature switch", "UnknownFeature", "default"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := MockFeatureSwitch(test.feature)
			assert.Equal(t, test.expected, actual, "The feature switch should return the expected outcome")
		})
	}
}
