package main

import (
	"fmt"
	"net/http"
)

// Step 1: Encapsulate the data and behavior related to a "User"
type User struct {
	name  string
	email string
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) SetName(name string) {
	u.name = name
}

// Step 2: Create an abstract interface for "UserService"
type UserService interface {
	GetUser() *User
	SaveUser(user *User)
}

// Step 3: Implement the abstraction with a concrete struct
type InMemoryUserService struct {
	currentUser *User
}

func (ims *InMemoryUserService) GetUser() *User {
	return ims.currentUser
}

func (ims *InMemoryUserService) SaveUser(user *User) {
	ims.currentUser = user
}

func main() {
	// Step 4: Use the abstraction through an interface variable
	var userService UserService = &InMemoryUserService{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Step 5: Interact with the encapsulated data through the abstract interface
		user := userService.GetUser()
		if user == nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		fmt.Fprintf(w, "Hello, %s!\n", user.GetName())
	})

	userService.SaveUser(&User{name: "Alice"})

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
