package main

import (
	"fmt"
)

// User struct
type User struct {
	ID   int
	Name string
}

// UserRepository interface
type UserRepository interface {
	FindUserByID(id int) (*User, error)
}

// InMemoryUserRepository struct
type InMemoryUserRepository struct {
	users map[int]*User
}

func (r *InMemoryUserRepository) FindUserByID(id int) (*User, error) {
	return r.users[id], nil
}

// Main function
func main() {
	users := map[int]*User{
		1: {ID: 1, Name: "Alice"},
		2: {ID: 2, Name: "Bob"},
	}
	repo := &InMemoryUserRepository{users: users}

	userID := 2
	user, err := repo.FindUserByID(userID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("User %d: %s\n", user.ID, user.Name) // Output: User 2: Bob
}
