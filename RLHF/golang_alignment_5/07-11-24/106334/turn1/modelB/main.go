package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Check if the number of arguments is correct
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run comment-to-code-ratio.go <codebase_directory>")
		os.Exit(1)
	}

	// Directory containing the codebase
	codebaseDir := os.Args[1]

	// Get the file paths of all Go files in the codebase
	filePaths, err := getGoFilePaths(codebaseDir)
	if err != nil {
		log.Fatalf("Error getting Go file paths: %v", err)
	}

	// Calculate the comment-to-code ratio for each file
	totalComments := 0
	totalCodeLines := 0
	for _, filePath := range filePaths {
		comments, codeLines, err := analyzeFile(filePath)
		if err != nil {
			log.Fatalf("Error analyzing file %s: %v", filePath, err)
		}
		totalComments += comments
		totalCodeLines += codeLines
	}

	// Calculate the global comment-to-code ratio
	globalRatio := float64(totalComments) / float64(totalCodeLines)

	// Display the results
	fmt.Printf("Total comments: %d\n", totalComments)
	fmt.Printf("Total code lines: %d\n", totalCodeLines)
	fmt.Printf("Global comment-to-code ratio: %.2f\n", globalRatio)
}

// getGoFilePaths returns the file paths of all Go files within the given directory.
func getGoFilePaths(dir string) ([]string, error) {
	var filePaths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".go") {
			filePaths = append(filePaths, path)
		}
		return nil
	})
	return filePaths, err
}

// analyzeFile analyzes the given Go file and returns the number of comments and code lines.
func analyzeFile(filePath string) (int, int, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return 0, 0, err
	}

	comments := 0
	codeLines := 0

	// Traverse the AST and count comments and code lines
	ast.Inspect(f, func(n ast.Node) bool {
		if n == nil {
			return true
		}

		// Count comments
		if _, ok := n.(*ast.Comment); ok {
			comments++
		}

		// Count code lines
		if n.Pos() != 0 && n.End() != 0 {
			codeLines++
		}

		return true
	})

	return comments, codeLines, nil
}
