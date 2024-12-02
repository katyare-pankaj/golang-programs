package main

import (
	"log"
	"runtime"
	"strings"
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

	// Retrieve licenses
	licenses := modData.Module.Licenses
	for _, license := range licenses {
		licenseType := license.Type
		if licenseType != "" {
			licenseType = strings.ReplaceAll(licenseType, "spdx/", "")
		}
		log.Printf("License Type: %s", licenseType)
		log.Printf("License Name: %s", license.Name)
		log.Printf("License URL: %s", license.URL)
		log.Println()
	}
}
