package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Comment struct {
	ID       int
	Content  string
	Children []*Comment
}

func generateComments(parent *Comment, depth int, maxDepth int) {
	// Stop generating comments if the desired depth is reached
	if depth >= maxDepth {
		return
	}

	// Generate random number of child comments for the current parent
	numChildren := rand.Intn(5) + 1
	for i := 0; i < numChildren; i++ {
		child := &Comment{
			ID:       len(parent.Children) + 1,
			Content:  fmt.Sprintf("Child comment %d of Parent %d", i+1, parent.ID),
			Children: []*Comment{},
		}
		parent.Children = append(parent.Children, child)

		// Recursively generate nested comments for the current child
		generateComments(child, depth+1, maxDepth)
	}
}

func printCommentTree(root *Comment) {
	fmt.Printf("Comment ID: %d, Content: %s\n", root.ID, root.Content)
	for _, child := range root.Children {
		printCommentTree(child)
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

	// Generate nested comments up to a maximum depth of 4
	generateComments(root, 1, 4)

	// Print the entire comment tree
	printCommentTree(root)
}
