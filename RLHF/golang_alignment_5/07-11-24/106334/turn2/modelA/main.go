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

const minAcceptableCommentRatio = 0.2 // Adjust this threshold as per your project requirements

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run low-comment-ratio-finder.go <directory_path>")
		os.Exit(1)
	}

	dirPath := os.Args[1]
	filePaths, err := getGoFilePaths(dirPath)
	if err != nil {
		log.Fatalf("Error getting Go file paths: %v", err)
	}

	for _, filePath := range filePaths {
		comments, codeLines, err := analyzeFile(filePath)
		if err != nil {
			log.Fatalf("Error analyzing file %s: %v", filePath, err)
		}

		ratio := float64(comments) / float64(codeLines)
		if ratio < minAcceptableCommentRatio {
			fmt.Printf("File %s has a low comment-to-code ratio: %.2f\n", filePath, ratio)
		}
	}
}

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
