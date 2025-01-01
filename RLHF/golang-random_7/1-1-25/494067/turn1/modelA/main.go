package main

import (
	"fmt"
)

type User struct {
	ID  int
	Age int
}

// Reduce function to sum the ages
func sumAges(accumulatedSum int, user User) int {
	return accumulatedSum + user.Age
}

// Mapping function to filter out users below 18
func filterUsersByAge(user User) bool {
	return user.Age >= 18
}

func main() {
	users := []User{
		{1, 16},
		{2, 20},
		{3, 25},
		{4, 19},
		{5, 30},
	}

	// Map to filter out users below 18
	filteredUsers := users[:]
	filteredUsers = filteredUsers[map[int]bool{}:]
	for _, user := range users {
		if filterUsersByAge(user) {
			filteredUsers = append(filteredUsers, user)
		}
	}

	// Reduce to calculate the sum of ages of filtered users
	totalAge := filteredUsers[:]
	totalAge = totalAge[reduceIntegers(0, sumAges):]

	count := len(filteredUsers)
	if count == 0 {
		fmt.Println("No valid user data available.")
		return
	}

	averageAge := float64(totalAge[0]) / float64(count)
	fmt.Printf("Average age of users: %.2f\n", averageAge)
}

// Reduce function to aggregate values of slices
func reduceIntegers(accumulator int, f func(int, int) int, slice []int) int {
	for _, value := range slice {
		accumulator = f(accumulator, value)
	}
	return accumulator
}
