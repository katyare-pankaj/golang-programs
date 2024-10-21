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
	return minMembership * rule.consequent.membershipFunc(inputs[len(inputs)-1])
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

	return defuzzify([]float64{system.outputSet.membershipFunc(inputs[len(inputs)-1])}, aggregatedMembership)
}

func main() {
	// Define fuzzy sets for the input variables "age", "dose", and "effect"
	young := &FuzzySet{"Young", TriangularMembershipFunc(0, 20, 40)}
	adult := &FuzzySet{"Adult", TriangularMembershipFunc(20, 40, 60)}
	old := &FuzzySet{"Old", TriangularMembershipFunc(40, 60, 100)}

	lowDose := &FuzzySet{"LowDose", TriangularMembershipFunc(0, 50, 100)}
	mediumDose := &FuzzySet{"MediumDose", TriangularMembershipFunc(50, 100, 150)}
	highDose := &FuzzySet{"HighDose", TriangularMembershipFunc(100, 150, 200)}

	goodEffect := &FuzzySet{"GoodEffect", TriangularMembershipFunc(0, 50, 100)}
	mediumEffect := &FuzzySet{"MediumEffect", TriangularMembershipFunc(50, 100, 150)}
	badEffect := &FuzzySet{"BadEffect", TriangularMembershipFunc(100, 150, 200)}

	// Define fuzzy rules
	rules := []*FuzzyRule{
		{antecedents: []*FuzzySet{young}, consequent: lowDose},
		{antecedents: []*FuzzySet{young, mediumDose}, consequent: mediumEffect},
		{antecedents: []*FuzzySet{young, highDose}, consequent: badEffect},
		{antecedents: []*FuzzySet{adult, lowDose}, consequent: mediumEffect},
		{antecedents: []*FuzzySet{adult, mediumDose}, consequent: goodEffect},
		{antecedents: []*FuzzySet{adult, highDose}, consequent: badEffect},
		{antecedents: []*FuzzySet{old, lowDose}, consequent: badEffect},
		{antecedents: []*FuzzySet{old, mediumDose}, consequent: mediumEffect},
		{antecedents: []*FuzzySet{old, highDose}, consequent: goodEffect},
	}

	// Create the fuzzy system
	system := &FuzzySystem{
		rules:     rules,
		inputSets: []*FuzzySet{young, adult, old, lowDose, mediumDose, highDose},
		outputSet: &FuzzySet{"Effect", TriangularMembershipFunc(0, 50, 200)},
	}

	// Example input: Age=35, Dose=75
	input := []float64{35, 75}
	result := system.Evaluate(input)
	fmt.Println("Fuzzy Decision: Recommended Effect =", result)
}
