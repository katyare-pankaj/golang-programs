package main

import (
	"fmt"
)

// FuzzySet represents a fuzzy set.
type FuzzySet struct {
	name           string
	membershipFunc func(float64) float64
}

// TriangularMembershipFunc creates a triangular membership function.
func TriangularMembershipFunc(min, peak, max float64) func(float64) float64 {
	return func(x float64) float64 {
		if x < min || x > max {
			return 0
		}
		if x <= peak {
			return (x - min) / (peak - min)
		}
		return (max - x) / (max - peak)
	}
}

// Defuzzify using Center of Gravity
func defuzzify(values []float64, memberships []float64) float64 {
	total := 0.0
	sum := 0.0
	for i := range values {
		mul := values[i] * memberships[i]
		total += mul
		sum += memberships[i]
	}
	if sum == 0 {
		return 0
	}
	return total / sum
}

// FuzzyRule represents a fuzzy rule.
type FuzzyRule struct {
	antecedents []*FuzzySet
	consequent  *FuzzySet
}

// Evaluate the fuzzy rule for a given input.
func (rule *FuzzyRule) Evaluate(inputs []float64) float64 {
	minMembership := 1.0
	for i, antecedent := range rule.antecedents {
		membership := antecedent.membershipFunc(inputs[i])
		if membership < minMembership {
			minMembership = membership
		}
	}
	return minMembership
}

// FuzzySystem represents a fuzzy system.
type FuzzySystem struct {
	rules     []*FuzzyRule
	inputSets []*FuzzySet
	outputSet *FuzzySet
}

// Evaluate the fuzzy system for a given input.
func (system *FuzzySystem) Evaluate(inputs []float64) float64 {
	aggregatedMembership := make([]float64, len(system.rules))
	for i, rule := range system.rules {
		aggregatedMembership[i] = rule.Evaluate(inputs)
	}

	// Defuzzification for treatment plan
	return defuzzify([]float64{0, 50, 100}, aggregatedMembership) // Adjust values for defuzzification
}

func main() {
	// Define fuzzy sets for the input variables "age", "blood_pressure", and "cholesterol"
	young := &FuzzySet{"Young", TriangularMembershipFunc(0, 20, 40)}
	adult := &FuzzySet{"Adult", TriangularMembershipFunc(20, 40, 60)}
	old := &FuzzySet{"Old", TriangularMembershipFunc(40, 60, 100)}

	normal := &FuzzySet{"Normal", TriangularMembershipFunc(90, 120, 150)}
	borderlineHigh := &FuzzySet{"Borderline High", TriangularMembershipFunc(120, 150, 180)}
	high := &FuzzySet{"High", TriangularMembershipFunc(150, 180, 250)}

	// Define fuzzy sets for the output variable "treatment_plan"
	lowDose := &FuzzySet{"Low Dose", TriangularMembershipFunc(0, 50, 100)}
	mediumDose := &FuzzySet{"Medium Dose", TriangularMembershipFunc(50, 100, 150)}
	highDose := &FuzzySet{"High Dose", TriangularMembershipFunc(100, 150, 200)}

	// Define fuzzy rules
	rules := []*FuzzyRule{
		{[]*FuzzySet{young, normal, high}, lowDose},
		{[]*FuzzySet{young, normal, borderlineHigh}, mediumDose},
		{[]*FuzzySet{young, normal, high}, highDose},
		{[]*FuzzySet{young, high, borderlineHigh}, highDose},
		{[]*FuzzySet{adult, normal, borderlineHigh}, mediumDose},
		{[]*FuzzySet{adult, normal, high}, highDose},
	}

	// Create the fuzzy system
	system := &FuzzySystem{
		rules:     rules,
		inputSets: []*FuzzySet{young, adult, old, normal, borderlineHigh, high},
		outputSet: lowDose, // Adjust according to your output set preference
	}

	// Patient data
	input := []float64{35, 130, 180} // Age, Blood Pressure, Cholesterol

	// Calculate the treatment plan using fuzzy logic
	treatmentPlan := system.Evaluate(input)
	fmt.Println("Treatment Plan: ", treatmentPlan)
}
