package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ImmutableUser data structure
type ImmutableUser struct {
	ID    int
	Name  string
	Email string
	// Other fields
}

func (u ImmutableUser) UpdateName(newName string) ImmutableUser {
	// Create a new instance with the updated name
	return ImmutableUser{
		ID:    u.ID,
		Name:  newName,
		Email: u.Email,
		// Other fields remain the same
	}
}

// Synchronization function signature
type Synchronizer func(user ImmutableUser) ImmutableUser

func TestSynchronizationLogic(t *testing.T) {
	// BDD Scenario: Synchronizing user data should update the name
	t.Run("Update name during synchronization", func(t *testing.T) {
		// Given an initial user data
		initialUser := ImmutableUser{
			ID:    1,
			Name:  "Alice",
			Email: "alice@example.com",
		}

		// When synchronizing the user data with a function that updates the name
		synchronizer := updateNameSynchronizer("Bob")
		synchronizedUser := synchronizer(initialUser)

		// Then the name in the synchronized user data should be updated
		assert.Equal(t, "Bob", synchronizedUser.Name, "Name should be updated during synchronization")
		// And the original user data should remain unchanged
		assert.Equal(t, initialUser.ID, synchronizedUser.ID, "ID should remain the same")
	})
}

func updateNameSynchronizer(newName string) Synchronizer {
	return func(user ImmutableUser) ImmutableUser {
		// Return a new user with the updated name
		return user.UpdateName(newName)
	}
}

func main() {
	fmt.Println("Data synchronization tests using immutable data structures")
}
