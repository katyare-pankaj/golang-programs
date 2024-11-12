// Example code for a simple recursive merge in Go
package main

import (
	"fmt"
)

type GameCharacter struct {
	Name   string
	Health int
}

func main() {
	char := GameCharacter{Name: "Alice", Health: 100}
	fmt.Println("Initial character:", char)
	// Merge changes from 'branch1'
	char = merge(char, branch1())
	fmt.Println("After merge from branch1:", char)
}

func merge(current Characters, other Characters) Characters {
	// Recursive merge of GameCharacter struct
	if current.Name != other.Name {
		fmt.Println("Name conflict, choosing current:", current.Name)
		return current
	}
	if current.Health != other.Health {
		fmt.Println("Health conflict, merging values:")
		current.Health = (current.Health + other.Health) / 2
	}
	return current
}

func branch1() Characters {
	// Changes from branch1
	return Characters{Name: "Alice", Health: 150}
}
