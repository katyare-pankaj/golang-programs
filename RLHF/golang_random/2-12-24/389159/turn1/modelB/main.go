package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/dep/gps"
	"golang.org/x/mod/modfile"
)

func main() {
	// Retrieve the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}
	// Read the go.mod file content
	goModFileContent, err := ioutil.ReadFile(cwd + "/go.mod")
	if err != nil {
		log.Fatalf("Error reading go.mod file: %v", err)
	}
	// Parse the go.mod file content
	modFile, err := modfile.Parse(cwd+"/go.mod", goModFileContent, nil)
	if err != nil {
		log.Fatalf("Error parsing go.mod file: %v", err)
	}
	// Iterate through the module dependencies
	for _, require := range modFile.Require {
		modulePath := require.Mod.Path
		version := require.Mod.Version
		fmt.Printf("Module: %s, Version: %s\n", modulePath, version)
		// Use gps package to retrieve module metadata (if available)
		m, err := gps.NewContext().LoadModule(modulePath, version)
		if err != nil {
			log.Printf("Error retrieving metadata for module %s: %v", modulePath, err)
		} else {
			fmt.Printf("\tDescription: %s\n", m.Info.Desc)
			fmt.Printf("\tAuthor: %s\n", m.Info.Author)
			fmt.Printf("\tHomepage: %s\n", m.Info.Homepage)
		}
		fmt.Println()
	}
}
