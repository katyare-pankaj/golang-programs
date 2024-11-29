package main

import (
	"fmt"
	"yourmodule"

	"github.com/fathis/featureflag"
)

func main() {
	// Register the feature flags with the featureflag package
	featureflag.Register()

	// Check if the new feature is enabled
	if featureflag.IsEnabled(yourmodule.FeatureNewFeature) {
		fmt.Println("New feature is enabled.")
	} else {
		fmt.Println("New feature is disabled.")
	}

	// Check if the improved performance feature is enabled
	if featureflag.IsEnabled(yourmodule.FeatureImprovedPerformance) {
		fmt.Println("Improved performance feature is enabled.")
	} else {
		fmt.Println("Improved performance feature is disabled.")
	}
}
