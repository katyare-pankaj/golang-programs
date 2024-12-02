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
	t.Run("Initial Enabled State", func(t *testing.T) {
		ft := NewFeatureToggle(true)
		assert.True(t, ft.IsEnabled(), "Feature should be enabled initially")
	})

	t.Run("Initial Disabled State", func(t *testing.T) {
		ft := NewFeatureToggle(false)
		assert.False(t, ft.IsEnabled(), "Feature should be disabled initially")
	})

	t.Run("Toggle Between States", func(t *testing.T) {
		ft := NewFeatureToggle(true)
		assert.True(t, ft.IsEnabled(), "Feature should be enabled initially")
		ft.ToggleEnabled()
		assert.False(t, ft.IsEnabled(), "Feature should be disabled after toggling")
		ft.ToggleEnabled()
		assert.True(t, ft.IsEnabled(), "Feature should be enabled after toggling again")
	})

	t.Run("Multiple Toggles", func(t *testing.T) {
		ft := NewFeatureToggle(false)
		assert.False(t, ft.IsEnabled(), "Feature should be disabled initially")
		for i := 0; i < 5; i++ {
			ft.ToggleEnabled()
			if i%2 == 0 {
				assert.True(t, ft.IsEnabled(), "Feature should be enabled after even toggles")
			} else {
				assert.False(t, ft.IsEnabled(), "Feature should be disabled after odd toggles")
			}
		}
	})
}
