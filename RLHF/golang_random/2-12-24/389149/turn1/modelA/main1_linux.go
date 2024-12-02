//go:build linux && amd64
// +build linux,amd64

package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("uname", "-a")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error running command:", err)
		os.Exit(1)
	}
	fmt.Println("Linux-specific code executed successfully.")
}
