package main

import (
	"fmt"
)

// Step 2: Define an abstract UserService interface
type UserService interface {
	GetUser() *User
	SaveUser(*User)
}

// Step 3: Implement the abstraction with different storage strategies
type InMemoryUserService struct {
	currentUser *User
}

func (ims *InMemoryUserService) GetUser() *User {
	return ims.currentUser
}

func (ims *InMemoryUserService) SaveUser(user *User) {
	ims.currentUser = user
}

type DatabaseUserService struct {
	// Database implementation details hidden
}

func main() {
	var userService UserService = &InMemoryUserService{} // Can be switched to DatabaseUserService later

	userService.SaveUser(&User{name: "Alice"})
	fmt.Println("User Name:", userService.GetUser().GetName())

	// Switching to a different storage strategy is easy and scalable
	//userService = &DatabaseUserService{}
}
