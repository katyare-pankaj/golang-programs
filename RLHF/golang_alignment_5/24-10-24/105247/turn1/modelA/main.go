package main

import (
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

// FeatureToggle represents a toggle for a specific feature
type FeatureToggle struct {
	Name    string
	Enabled bool
}

// DataRetentionPolicy defines the data retention period for a feature
type DataRetentionPolicy struct {
	Feature string
	Period  time.Duration
}

var (
	featureToggles = []FeatureToggle{
		{Name: "NewAwesomeFeature", Enabled: false},
		{Name: "CoreUserTracking", Enabled: true},
	}

	dataRetentionPolicies = []DataRetentionPolicy{
		{Feature: "NewAwesomeFeature", Period: 24 * time.Hour},
		{Feature: "CoreUserTracking", Period: 30 * 24 * time.Hour},
	}
)

func isFeatureEnabled(featureName string) bool {
	for _, toggle := range featureToggles {
		if toggle.Name == featureName {
			return toggle.Enabled
		}
	}
	return false
}

func getDataRetentionPeriod(featureName string) time.Duration {
	for _, policy := range dataRetentionPolicies {
		if policy.Feature == featureName {
			return policy.Period
		}
	}
	return 0
}

func handleUserData(userID string, featureName string, data interface{}) {
	if !isFeatureEnabled(featureName) {
		return // Feature is disabled, no data retention
	}

	retentionPeriod := getDataRetentionPeriod(featureName)
	if retentionPeriod == 0 {
		fmt.Printf("Retention period not defined for feature: %s\n", featureName)
		return
	}

	// Simulate data storage with retention period
	fmt.Printf("Storing data for user %s, feature %s: %v (Retained for: %s)\n", userID, featureName, data, retentionPeriod)
	time.AfterFunc(retentionPeriod, func() {
		fmt.Printf("Data for user %s, feature %s expired.\n", userID, featureName)
	})
}

func main() {
	app := &cli.App{
		Name:  "VR Data Retention Manager",
		Usage: "Manages data retention for VR features using feature toggles",
		Commands: []*cli.Command{
			{
				Name:   "toggle",
				Usage:  "Toggle the status of a feature",
				Action: toggleFeature,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "feature",
						Aliases:  []string{"f"},
						Usage:    "Name of the feature to toggle",
						Required: true,
					},
				},
			},
			{
				Name:   "store",
				Usage:  "Store user data for a feature",
				Action: storeUserData,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "user",
						Aliases:  []string{"u"},
						Usage:    "User ID",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "feature",
						Aliases:  []string{"f"},
						Usage:    "Name of the feature",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "data",
						Aliases:  []string{"d"},
						Usage:    "Data to store",
						Required: true,
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}

func toggleFeature(c *cli.Context) error {
	featureName := c.String("feature")
	for i, toggle := range featureToggles {
		if toggle.Name == featureName {
			toggle.Enabled = !toggle.Enabled
			fmt.Printf("Feature '%s' toggled to: %t\n", featureName, toggle.Enabled)
			featureToggles[i] = toggle
			return nil
		}
	}
	return fmt.Errorf("feature '%s' not found", featureName)
}

func storeUserData(c *cli.Context) error {
	userID := c.String("user")
	featureName := c.String("feature")
	data := c.String("data")

	handleUserData(userID, featureName, data)
	return nil
}
