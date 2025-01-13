package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Comment struct {
	ID       int
	Content  string
	Children [][]*Comment // Using an array of arrays to handle nested comments
}

func generateComments(parent *Comment, depth int, maxDepth int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Stop generating comments if the desired depth is reached
	if depth >= maxDepth {
		return
	}

	// Generate random number of child comments for the current parent
	numChildren := rand.Intn(5) + 1
	parent.Children = make([][]*Comment, numChildren)
	for i := 0; i < numChildren; i++ {
		child := &Comment{
			ID:       i + 1,
			Content:  fmt.Sprintf("Child comment %d of Parent %d", i+1, parent.ID),
			Children: nil, // Children will be populated later
		}
		parent.Children[i] = []*Comment{child}

		// Recursively generate nested comments for the current child
		wg.Add(1)
		go generateComments(child, depth+1, maxDepth, wg)
	}
}

func printCommentTree(root *Comment, indent string) {
	fmt.Printf("%sComment ID: %d, Content: %s\n", indent, root.ID, root.Content)
	for _, childGroup := range root.Children {
		for _, child := range childGroup {
			printCommentTree(child, indent+"  ")
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Root comment
	root := &Comment{
		ID:       1,
		Content:  "Root Comment",
		Children: nil,
	}

	var wg sync.WaitGroup
	// Generate nested comments up to a maximum depth of 8 (to demonstrate deep nesting)
	wg.Add(1)
	go generateComments(root, 1, 8, &wg)
	wg.Wait()

	// Print the entire comment tree
	printCommentTree(root, "")
}
