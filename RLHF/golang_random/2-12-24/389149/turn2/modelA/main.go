package main

import (
	"log"
	"runtime"
)

func main() {
	// Retrieve the module's metadata
	modRoot, _ := runtime.LookupPath(".")
	modData, err := modroot.ModData(modRoot)
	if err != nil {
		log.Fatalf("Could not read module data: %v", err)
	}

	// Display module information
	log.Printf("Module Name: %s", modData.Module.Path)
	log.Printf("Module Version: %s", modData.Module.Version)

	// Retrieve and display dependency versions
	for _, require := range modData.Module.Require {
		log.Printf("Dependency: %s@%s", require.Path, require.Version)
	}
}
