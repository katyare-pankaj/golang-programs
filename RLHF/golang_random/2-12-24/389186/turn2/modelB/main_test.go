package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// FeatureToggle is a simulated feature toggle system.
type FeatureToggle struct {
	toggles map[string]bool
}

// IsFeatureEnabled checks if a given feature is enabled.
func (f *FeatureToggle) IsFeatureEnabled(feature string) bool {
	enabled, exists := f.toggles[feature]
	return exists && enabled
}

// TestFeatureToggle verifies the behavior of the feature toggles.
func TestFeatureToggle(t *testing.T) {
	// Initializing the feature toggle system
	toggles := &FeatureToggle{
		toggles: map[string]bool{
			"FeatureA": true,
			"FeatureB": false,
		},
	}

	// Test cases
	tests := []struct {
		name     string
		feature  string
		expected bool
	}{
		{"Enabled feature", "FeatureA", true},
		{"Disabled feature", "FeatureB", false},
		{"Unknown feature", "FeatureC", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := toggles.IsFeatureEnabled(test.feature)
			assert.Equal(t, test.expected, result, "Feature toggle state should match expected value")
		})
	}
}
