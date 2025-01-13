package main

import (
	"fmt"
)

// Comment represents a single comment with its content and child comments
type Comment struct {
	Content  string
	Children []*Comment
}

// AddChild adds a child comment to the current comment
func (c *Comment) AddChild(child *Comment) {
	c.Children = append(c.Children, child)
}

// Print recursively prints the comment and its children
func (c *Comment) Print() {
	fmt.Printf("%s\n", c.Content)
	for _, child := range c.Children {
		child.Print()
	}
}

func main() {
	// Creating a root comment
	root := &Comment{
		Content: "This is the root comment.",
	}

	// Adding some first-level child comments
	child1 := &Comment{
		Content: "First child comment.",
	}
	child2 := &Comment{
		Content: "Second child comment.",
	}

	// Adding second-level child comments to child1
	child1.AddChild(&Comment{
		Content: "First grandchild comment.",
	})
	child1.AddChild(&Comment{
		Content: "Second grandchild comment.",
	})

	// Adding child comments to the root
	root.AddChild(child1)
	root.AddChild(child2)

	// Print the entire comment tree
	root.Print()
}
