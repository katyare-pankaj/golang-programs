package main

import (
	"fmt"
)

// User struct
type User struct {
	ID   int
	Name string
}

// UserDAO interface
type UserDAO interface {
	GetUser(id int) (*User, error)
}

// MockUserDAO struct
type MockUserDAO struct{}

func (m *MockUserDAO) GetUser(id int) (*User, error) {
	users := map[int]*User{
		1: {ID: 1, Name: "Alice"},
		2: {ID: 2, Name: "Bob"},
	}

	user, ok := users[id]
	if !ok {
		return nil, fmt.Errorf("user not found with id: %d", id)
	}
	return user, nil
}

// Main function
func main() {
	dao := &MockUserDAO{} // Replace this with a real DAO implementation

	userID := 2
	user, err := dao.GetUser(userID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("User %d: %s\n", user.ID, user.Name) // Output: User 2: Bob
}
