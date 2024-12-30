package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Define the structure that the JSON data will be unmarshalled into.
type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func readAndUnmarshalFile() (user *User, err error) {
	// Read the file content.
	filename := "users.json"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON content into a User struct.
	user = &User{}
	err = json.Unmarshal(data, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func main() {
	// Read and unmarshal the JSON file.
	user, err := readAndUnmarshalFile()
	if err != nil {
		// Handle error: could be a file not found error or JSON syntax error, etc.
		switch err := err.(type) {
		case *os.PathError:
			log.Printf("Error: The file %s was not found.", err.Path)
		case json.SyntaxError:
			log.Printf("Error: JSON syntax error: %s", err.Error())
		default:
			log.Printf("An unexpected error occurred: %s", err.Error())
		}
		return
	}

	// If no error, display the user's information.
	fmt.Printf("Name: %s, Age: %d, Email: %s, Password: %s\n", user.Name, user.Age, user.Email, user.Password)
}
