package main

import (
	"math/rand"
	"testing"
)

// SelectVariant selects a variant based on a given percentage.
func SelectVariant(percentage float64) string {
	if percentage < 0 || percentage > 100 {
		panic("percentage must be between 0 and 100")
	}
	randomNumber := rand.Float64() * 100
	if randomNumber < percentage {
		return "variant_a"
	}
	return "variant_b"
}

func TestSelectVariant(t *testing.T) {
	rand.Seed(42) // Seed for reproducibility

	// Test case: Variant A should be selected 30% of the time
	for i := 0; i < 100; i++ {
		variant := SelectVariant(30)
		if variant != "variant_a" && variant != "variant_b" {
			t.Errorf("Unexpected variant: %s", variant)
		}
		if i%3 == 0 {
			assertVariantDistribution(t, variant, "variant_a", 30)
		}
	}

	// Test case: Variant B should be selected 70% of the time
	for i := 0; i < 100; i++ {
		variant := SelectVariant(70)
		if variant != "variant_a" && variant != "variant_b" {
			t.Errorf("Unexpected variant: %s", variant)
		}
		if i%3 == 0 {
			assertVariantDistribution(t, variant, "variant_b", 70)
		}
	}
}

func assertVariantDistribution(t *testing.T, variant string, expected string, percentage float64) {
	if variant != expected {
		t.Errorf("Expected variant %s, got %s", expected, variant)
	}
}

func main() {
	testing.Main()
}
