package main

import (
	"fmt"
	"os"
	"strings"
)

func checkVulnerableLibraries(dependencies []string) bool {
	vulnerableLibraries := []string{"old_library_version", "another_vulnerable_lib"}

	for _, dep := range dependencies {
		for _, vuln := range vulnerableLibraries {
			if strings.Contains(dep, vuln) {
				fmt.Printf("Vulnerable library found: %s\n", dep)
				return true
			}
		}
	}
	return false
}

func main() {
	// Sample data representing dependencies of a legacy application
	dependencies := []string{
		"package1>=1.2.3",
		"vulnerable_library=0.1.0",
		"safe_library=2.3.4",
		"old_library_version=1.0.0",
	}

	fmt.Println("Checking for vulnerable libraries...")
	if checkVulnerableLibraries(dependencies) {
		os.Exit(1) // Exit with status 1 to indicate failure
	}
	fmt.Println("No vulnerable libraries found.")
	fmt.Println("Security check completed successfully.")
}
