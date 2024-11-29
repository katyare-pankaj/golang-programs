package yourmodule

import (
	"fmt"

	"github.com/fathis/featureflag"
)

const (
	// FeatureNewFeature is a flag for the new feature
	FeatureNewFeature featureflag.Feature = "newFeature"

	// FeatureImprovedPerformance is a flag for improved performance
	FeatureImprovedPerformance featureflag.Feature = "improvedPerformance"
)

func UseNewFeature() {
	if featureflag.IsEnabled(FeatureNewFeature) {
		// Implement the new feature code here
		fmt.Println("Using the new feature.")
	} else {
		// Fall back to the old feature code here
		fmt.Println("Using the old feature.")
	}
}

func OptimizePerformance() {
	if featureflag.IsEnabled(FeatureImprovedPerformance) {
		// Implement performance optimization code here
		fmt.Println("Optimizing performance.")
	} else {
		// Don't perform any performance optimizations
		fmt.Println("Performance optimization disabled.")
	}
}
