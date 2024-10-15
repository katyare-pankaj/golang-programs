package main

import (
	"fmt"
)

// Immutable User data structure
type User struct {
	ID      int
	Name    string
	Email   string
	version int
}

func (u User) WithName(name string) User {
	return User{
		ID:      u.ID,
		Name:    name,
		Email:   u.Email,
		version: u.version + 1,
	}
}

func (u User) WithEmail(email string) User {
	return User{
		ID:      u.ID,
		Name:    u.Name,
		Email:   email,
		version: u.version + 1,
	}
}

// Synchronization logic using immutable data
func synchronizeUser(existingUser User, updatedFields map[string]string) User {
	syncedUser := existingUser
	for key, value := range updatedFields {
		switch key {
		case "name":
			syncedUser = syncedUser.WithName(value)
		case "email":
			syncedUser = syncedUser.WithEmail(value)
		}
	}
	return syncedUser
}

func main() {
	existingUser := User{ID: 1, Name: "Alice", Email: "alice@example.com", version: 0}
	updatedFields := map[string]string{"name": "Bob", "email": "bob@example.com"}
	syncedUser := synchronizeUser(existingUser, updatedFields)
	fmt.Printf("Synchronized user: %+v\n", syncedUser)
}
