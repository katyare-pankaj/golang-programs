//go:build linux && windows && darwin
// +build linux,windows,darwin

package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Run platform-specific commands
	switch os.Getenv("GOOS") {
	case "linux":
		runLinuxCommand()
	case "windows":
		runWindowsCommand()
	case "darwin":
		runMacOSCommand()
	default:
		fmt.Println("Unsupported platform")
	}
}

func runLinuxCommand() {
	fmt.Println("Running Linux-specific command: uname -a")
	cmd := exec.Command("uname", "-a")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error running command:", err)
	}
}

func runWindowsCommand() {
	fmt.Println("Running Windows-specific command: ver")
	cmd := exec.Command("ver")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error running command:", err)
	}
}

func runMacOSCommand() {
	fmt.Println("Running macOS-specific command: sw_vers")
	cmd := exec.Command("sw_vers")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error running command:", err)
	}
}
