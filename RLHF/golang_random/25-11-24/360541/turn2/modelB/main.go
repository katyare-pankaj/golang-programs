package main

import (
	"fmt"
	"log"

	"github.com/perforce/p4go"
)

func main() {
	// Create a new Perforce client
	client, err := p4go.New()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the Perforce server
	if err := client.Connect("p4.example.com:1666"); err != nil {
		log.Fatal(err)
	}

	// Login to the Perforce server
	if err := client.Login("your_username", "your_password"); err != nil {
		log.Fatal(err)
	}

	// Create a new branch
	branchName := "new_branch"
	if err := client.Branch(branchName, "mainline", p4go.BranchOptions{}); err != nil {
		log.Fatal(err)
	}

	// Edit a file in the branch
	filePath := "//depot/mainline/some_file.go"
	if err := client.Edit(filePath); err != nil {
		log.Fatal(err)
	}

	// Add changes to the change list
	changelist, err := client.NewChangelist("Description of changes")
	if err != nil {
		log.Fatal(err)
	}
	if err := changelist.Add(filePath); err != nil {
		log.Fatal(err)
	}

	// Submit the change list
	if err := changelist.Submit(); err != nil {
		log.Fatal(err)
	}

	// Delete the branch (optional)
	if err := client.DeleteBranch(branchName); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Changes submitted successfully!")
}
