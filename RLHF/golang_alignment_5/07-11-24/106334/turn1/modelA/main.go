package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run comment_to_code_ratio.go <directory_path>")
		os.Exit(1)
	}

	dirPath := os.Args[1]
	commentToCodeRatio, err := calculateCommentToCodeRatio(dirPath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("Comment-to-Code Ratio: %.2f\n", commentToCodeRatio)
}

func calculateCommentToCodeRatio(dirPath string) (float64, error) {
	totalComments := 0
	totalCodeLines := 0

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".go" {
			return nil
		}

		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return err
		}

		countCommentsAndCodeLines(f, &totalComments, &totalCodeLines)
		return nil
	})
	if err != nil {
		return 0, err
	}

	if totalCodeLines == 0 {
		return 0, fmt.Errorf("no code lines found in the codebase")
	}

	return float64(totalComments) / float64(totalCodeLines), nil
}

func countCommentsAndCodeLines(f *ast.File, totalComments *int, totalCodeLines *int) {
	for _, stmt := range f.Decls {
		countCommentsAndCodeLinesInStmt(stmt, totalComments, totalCodeLines)
	}
}

func countCommentsAndCodeLinesInStmt(stmt ast.Stmt, totalComments *int, totalCodeLines *int) {
	switch stmt := stmt.(type) {
	case *ast.CommentGroup:
		for _, comment := range stmt.List {
			if !strings.HasPrefix(comment.Text, "//") {
				continue
			}
			*totalComments++
		}
	default:
		*totalCodeLines++
	}

	for _, child := range stmt.Pos() {
		countCommentsAndCodeLinesInStmt(child, totalComments, totalCodeLines)
	}
}
