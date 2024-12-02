package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("myprogram")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error running command:", err)
		os.Exit(1)
	}
	fmt.Println("Windows-specific code executed successfully.")
}
