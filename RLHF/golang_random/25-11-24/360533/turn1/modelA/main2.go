package main

import (
	"fmt"

	"github.com/go-bazaar/bazaar/bzr"
)

func main() {
	// Initialize a new Bazaar repository
	repo, err := bzr.InitRepository("/path/to/my/repo")
	if err != nil {
		fmt.Println("Error initializing repository:", err)
		return
	}
	// Add files to the repository
	err = repo.Add("main.go", "README.md")
	if err != nil {
		fmt.Println("Error adding files:", err)
		return
	}
	// Commit the changes
	err = repo.Commit("Initial commit")
	if err != nil {
		fmt.Println("Error committing changes:", err)
		return
	}
}
