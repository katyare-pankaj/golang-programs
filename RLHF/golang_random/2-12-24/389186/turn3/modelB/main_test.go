package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// FeatureToggle represents a simple feature toggle mechanism
type FeatureToggle struct {
	enabled bool
}

// NewFeatureToggle returns a new FeatureToggle instance with the specified initial state
func NewFeatureToggle(enabled bool) *FeatureToggle {
	return &FeatureToggle{enabled: enabled}
}

// IsEnabled checks if the feature toggle is enabled
func (ft *FeatureToggle) IsEnabled() bool {
	return ft.enabled
}

// SetEnabled sets the enabled state of the feature toggle
func (ft *FeatureToggle) SetEnabled(enabled bool) {
	ft.enabled = enabled
}

// A simple function to demonstrate feature toggle behavior
func FeatureBehavior(ft *FeatureToggle) string {
	if ft.IsEnabled() {
		return "Feature is enabled"
	}
	return "Feature is disabled"
}

func TestFeatureToggle(t *testing.T) {
	// Test that a new feature toggle is enabled when initialized with true
	t.Run("Initialize with enabled state", func(t *testing.T) {
		ft := NewFeatureToggle(true)
		assert.True(t, ft.IsEnabled(), "The feature toggle should be enabled")
		assert.Equal(t, "Feature is enabled", FeatureBehavior(ft), "Expected behavior when feature is enabled")
	})

	// Test that a new feature toggle is disabled when initialized with false
	t.Run("Initialize with disabled state", func(t *testing.T) {
		ft := NewFeatureToggle(false)
		assert.False(t, ft.IsEnabled(), "The feature toggle should be disabled")
		assert.Equal(t, "Feature is disabled", FeatureBehavior(ft), "Expected behavior when feature is disabled")
	})

	// Test toggling the feature from disabled to enabled
	t.Run("Toggle from disabled to enabled", func(t *testing.T) {
		ft := NewFeatureToggle(false)
		ft.SetEnabled(true)
		assert.True(t, ft.IsEnabled(), "The feature toggle should now be enabled")
		assert.Equal(t, "Feature is enabled", FeatureBehavior(ft), "Expected behavior when feature is enabled")
	})

	// Test toggling the feature from enabled to disabled
	t.Run("Toggle from enabled to disabled", func(t *testing.T) {
		ft := NewFeatureToggle(true)
		ft.SetEnabled(false)
		assert.False(t, ft.IsEnabled(), "The feature toggle should now be disabled")
		assert.Equal(t, "Feature is disabled", FeatureBehavior(ft), "Expected behavior when feature is disabled")
	})
}
