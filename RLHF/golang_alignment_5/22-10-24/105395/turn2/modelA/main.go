package main

import (
	"fmt"
)

// User represents a user entity
type User struct {
	ID   int
	Name string
}

// UserRepository interface defines the operations for retrieving users
type UserRepository interface {
	GetUserByID(id int) (*User, error)
}

// InMemoryUserRepository implements UserRepository using in-memory data
type InMemoryUserRepository struct {
	users map[int]*User
}

func (r *InMemoryUserRepository) GetUserByID(id int) (*User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, fmt.Errorf("user with id %d not found", id)
	}
	return user, nil
}

// InitializeInMemoryRepo is a helper function to populate the in-memory repo
func InitializeInMemoryRepo() *InMemoryUserRepository {
	repo := &InMemoryUserRepository{
		users: make(map[int]*User),
	}
	repo.users[1] = &User{ID: 1, Name: "Alice"}
	repo.users[2] = &User{ID: 2, Name: "Bob"}
	return repo
}

func main() {
	// Client code using the Repository interface
	repo := InitializeInMemoryRepo()

	userID := 2
	user, err := repo.GetUserByID(userID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("User Found: ID=%d, Name=%s\n", user.ID, user.Name)
}
