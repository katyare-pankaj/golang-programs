package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// FeatureToggle represents a feature toggle
type FeatureToggle struct {
	enabled bool
}

// NewFeatureToggle creates a new feature toggle with the given enabled status
func NewFeatureToggle(enabled bool) *FeatureToggle {
	return &FeatureToggle{enabled: enabled}
}

// IsEnabled returns true if the feature is enabled, false otherwise
func (ft *FeatureToggle) IsEnabled() bool {
	return ft.enabled
}

// ToggleEnabled toggles the enabled status of the feature
func (ft *FeatureToggle) ToggleEnabled() {
	ft.enabled = !ft.enabled
}

func TestFeatureToggle(t *testing.T) {
	t.Run("Create a new toggle", func(t *testing.T) {
		ft := NewFeatureToggle(true)
		assert.True(t, ft.IsEnabled(), "Toggle should be enabled initially")
	})

	t.Run("Toggle enabled state", func(t *testing.T) {
		ft := NewFeatureToggle(true)
		assert.True(t, ft.IsEnabled(), "Toggle should be enabled initially")
		ft.ToggleEnabled()
		assert.False(t, ft.IsEnabled(), "Toggle should be disabled after toggling")
		ft.ToggleEnabled()
		assert.True(t, ft.IsEnabled(), "Toggle should be enabled after toggling again")
	})

	t.Run("Toggle disabled state", func(t *testing.T) {
		ft := NewFeatureToggle(false)
		assert.False(t, ft.IsEnabled(), "Toggle should be disabled initially")
		ft.ToggleEnabled()
		assert.True(t, ft.IsEnabled(), "Toggle should be enabled after toggling")
		ft.ToggleEnabled()
		assert.False(t, ft.IsEnabled(), "Toggle should be disabled after toggling again")
	})
}
