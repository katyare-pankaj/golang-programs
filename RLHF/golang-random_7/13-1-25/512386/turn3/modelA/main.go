package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Comment struct to represent each comment
type Comment struct {
	ID       int
	Content  string
	Children []*Comment
}

// GenerateComments recursively generates comments with a depth limiter
func GenerateComments(parent *Comment, depth int, maxDepth int, limit int) {
	if depth >= maxDepth || limit <= 0 {
		return
	}

	numChildren := rand.Intn(limit) + 1
	for i := 0; i < numChildren; i++ {
		child := &Comment{
			ID:       len(parent.Children) + 1,
			Content:  fmt.Sprintf("Child comment %d of Parent %d", i+1, parent.ID),
			Children: []*Comment{},
		}
		parent.Children = append(parent.Children, child)
		GenerateComments(child, depth+1, maxDepth, limit)
	}
}

// PrintCommentTree recursively prints the comment tree
func PrintCommentTree(root *Comment, depth int) {
	fmt.Printf("%sComment ID: %d, Content: %s\n", strings.Repeat("  ", depth), root.ID, root.Content)
	for _, child := range root.Children {
		PrintCommentTree(child, depth+1)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Root comment
	root := &Comment{
		ID:       1,
		Content:  "Root Comment",
		Children: []*Comment{},
	}

	// Generate nested comments with a maximum depth of 4 and a child limit of 3
	GenerateComments(root, 1, 4, 3)

	// Print the entire comment tree
	fmt.Println("Comment Tree:")
	PrintCommentTree(root, 0)
}
