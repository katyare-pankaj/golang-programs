//go:build windows && amd64
// +build windows,amd64

package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("ver")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error running command:", err)
		os.Exit(1)
	}
	fmt.Println("Windows-specific code executed successfully.")
}
