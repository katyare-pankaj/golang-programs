package fileshare

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// StoreData stores the given data to a file for the specified user.
func StoreData(user, filename, data string) error {
	// Simulate a file storage location based on the user.
	filePath := fmt.Sprintf("./data/%s/%s", user, filename)
	err := os.MkdirAll(fmt.Sprintf("./data/%s", user), os.ModePerm)
	if err != nil {
		return err
	}

	// Create and write the file.
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return err
	}

	// Debugging: Output the file creation path.
	fmt.Println("File created at:", filePath)

	return nil
}

// DeleteData deletes the specified file for the given user.
func DeleteData(user, filename string) error {
	// Simulate a file storage location based on the user.
	filePath := fmt.Sprintf("./data/%s/%s", user, filename)

	// Debugging: Output the file deletion path.
	fmt.Println("Attempting to delete file at:", filePath)

	// Check if the file exists.
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return errors.New("file does not exist")
	}

	// Check if the user is trying to delete another user's file.
	if !strings.HasPrefix(filePath, fmt.Sprintf("./data/%s", user)) && user != "admin" {
		return errors.New("user does not have permission to delete this file")
	}

	// Simulate admin check by allowing any file deletion by admin.
	if user == "admin" {
		err := os.Remove(filePath)
		if err != nil {
			return err
		}
		return nil
	}

	// Proceed to delete the file for the specified user.
	err := os.Remove(filePath)
	if err != nil {
		return err
	}

	return nil
}
