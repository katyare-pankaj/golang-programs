package main

import (
	"fmt"
	"os"

	"os/exec"
)

func checkoutFile(filePath string) error {
	cmd := exec.Command("p4", "checkout", filePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to checkout file: %w", err)
	}
	fmt.Println("File checked out successfully.")
	return nil
}

func main() {
	filePath := "file.go"
	err := checkoutFile(filePath)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
