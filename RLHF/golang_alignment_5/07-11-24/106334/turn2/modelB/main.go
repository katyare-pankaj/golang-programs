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

const (
	// Ratio threshold below which files will be identified as having low comment-to-code ratio
	lowRatioThreshold = 0.2
)

func main() {
	// Check if the number of arguments is correct
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run low-comment-to-code-ratio.go <codebase_directory>")
		os.Exit(1)
	}

	// Directory containing the codebase
	codebaseDir := os.Args[1]

	// Get the file paths of all Go files in the codebase
	filePaths, err := getGoFilePaths(codebaseDir)
	if err != nil {
		log.Fatalf("Error getting Go file paths: %v", err)
	}

	// Analyze each file and print files with low comment-to-code ratio
	for _, filePath := range filePaths {
		comments, codeLines, err := analyzeFile(filePath)
		if err != nil {
			log.Fatalf("Error analyzing file %s: %v", filePath, err)
		}
		ratio := float64(comments) / float64(codeLines)
		if ratio < lowRatioThreshold {
			fmt.Printf("File with low comment-to-code ratio: %s (%.2f)\n", filePath, ratio)
		}
	}
}

// Remaining functions (getGoFilePaths, analyzeFile) remain the same as in the previous response

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

func analyzeFile(filePath string) (int, int, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return 0, 0, err
	}

	var comments int
	var codeLines int

	ast.Inspect(f, func(n ast.Node) bool {
		if n == nil {
			return true
		}

		if _, ok := n.(*ast.Comment); ok {
			comments++
		}

		if n.Pos() != 0 && n.End() != 0 {
			codeLines++
		}

		return true
	})

	return comments, codeLines, nil
}
