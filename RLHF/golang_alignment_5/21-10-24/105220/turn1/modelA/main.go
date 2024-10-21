package main

import "fmt"

// FuzzySet represents a fuzzy set with a name and membership function.
type FuzzySet struct {
	name           string
	membershipFunc func(float64) float64
}

// TriangularMembershipFunc creates a triangular membership function for a fuzzy set.
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

// TrapezoidalMembershipFunc creates a trapezoidal membership function for a fuzzy set.
func TrapezoidalMembershipFunc(min1, max1, min2, max2 float64) func(float64) float64 {
	return func(x float64) float64 {
		if x < min1 || x > max2 {
			return 0
		}
		if x <= max1 {
			return (x - min1) / (max1 - min1)
		}
		if x <= min2 {
			return 1
		}
		return (max2 - x) / (max2 - min2)
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

// FuzzyRule represents a fuzzy rule with antecedents and a consequent.
type FuzzyRule struct {
	antecedents []*FuzzySet
	consequent  *FuzzySet
}

// Evaluate the fuzzy rule for a given input value.
func (rule *FuzzyRule) Evaluate(input float64) float64 {
	minMembership := 1.0
	for _, antecedent := range rule.antecedents {
		membership := antecedent.membershipFunc(input)
		if membership < minMembership {
			minMembership = membership
		}
	}
	return minMembership * rule.consequent.membershipFunc(input)
}

// FuzzySystem represents a fuzzy system with rules and input/output variables.
type FuzzySystem struct {
	rules     []*FuzzyRule
	inputSets []*FuzzySet
	outputSet *FuzzySet
}

// Evaluate the fuzzy system for a given input value.
func (system *FuzzySystem) Evaluate(input float64) float64 {
	aggregatedMembership := make([]float64, len(system.rules))
	for i, rule := range system.rules {
		aggregatedMembership[i] = rule.Evaluate(input)
	}
	return defuzzify([]float64{system.outputSet.membershipFunc(input)}, aggregatedMembership)
}

func main() {
	// Define fuzzy sets for the input variable "age"
	young := FuzzySet{"Young", TriangularMembershipFunc(0, 20, 40)}
	adult := FuzzySet{"Adult", TriangularMembershipFunc(20, 40, 60)}
	old := FuzzySet{"Old", TriangularMembershipFunc(40, 60, 100)}

	fmt.Println(young, adult, old)

}
