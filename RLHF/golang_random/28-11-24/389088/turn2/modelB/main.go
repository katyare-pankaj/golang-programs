package main

import (
	"fmt"
	"go-programs/RLHF/golang_random/28-11-24/389088/turn2/modelB/fileshare"
)

func main() {
	dataToStore := "Some content to store"
	fileToDelete := "myFile.txt"
	otherUserFile := "otherUserFile.txt"

	if err := fileshare.StoreData("user", "myFile.txt", dataToStore); err != nil {
		fmt.Println("Error storing data:", err)
	} else {
		fmt.Println("Data stored successfully.")
	}

	if err := fileshare.DeleteData("user", fileToDelete); err != nil {
		fmt.Println("Error deleting file:", err)
	} else {
		fmt.Println("File deleted successfully.")
	}
	if err := fileshare.DeleteData("user", otherUserFile); err != nil {
		fmt.Println("Error deleting file:", err) // User shouldn't be able to delete other users' files
	}
	if err := fileshare.DeleteData("admin", otherUserFile); err != nil {
		fmt.Println("Error deleting file:", err) // Admin should be able to delete any file
	}
}
