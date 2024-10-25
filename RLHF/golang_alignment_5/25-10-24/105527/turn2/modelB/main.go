package main

// Example of loop with complex branching (that can be difficult for the auto-vectorizer to handle)
func sumArrayWithBranching(arr []int32) int32 {
	total := int32(0)
	for _, value := range arr {
		if value%2 == 0 {
			total += value * 2
		} else {
			total += value
		}
	}
	return total
}
