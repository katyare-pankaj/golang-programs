package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/mod/modfile"
)

func main() {
	// Open the go.mod file
	file, err := os.ReadFile("go.mod")
	if err != nil {
		log.Fatalf("Failed to read go.mod: %v", err)
	}

	// Parse the go.mod file
	modFile, err := modfile.Parse("go.mod", file, nil)
	if err != nil {
		log.Fatalf("Failed to parse go.mod: %v", err)
	}

	// Display module information
	fmt.Printf("Module Name: %s\n", modFile.Module.Mod.Path)
	fmt.Printf("Go Version: %s\n", modFile.Go.Version)

	// Display dependencies
	fmt.Println("Dependencies:")
	for _, req := range modFile.Require {
		fmt.Printf("- %s %s\n", req.Mod.Path, req.Mod.Version)
	}
}
