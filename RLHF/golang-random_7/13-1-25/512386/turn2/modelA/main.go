package main

import (
	"fmt"
)

// Comment represents a single comment with content and a slice of children
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

	// Adding first-level child comments
	child1 := &Comment{
		Content: "First child comment.",
	}
	child2 := &Comment{
		Content: "Second child comment.",
	}

	// Adding second-level child comments
	grandchild1_1 := &Comment{
		Content: "First grandchild comment of child1.",
	}
	grandchild1_2 := &Comment{
		Content: "Second grandchild comment of child1.",
	}
	grandchild2_1 := &Comment{
		Content: "First grandchild comment of child2.",
	}

	// Adding third-level child comments
	greatgrandchild1_1_1 := &Comment{
		Content: "First great-grandchild comment of grandchild1_1.",
	}

	// Adding children to each level
	child1.AddChild(grandchild1_1)
	child1.AddChild(grandchild1_2)
	child2.AddChild(grandchild2_1)
	grandchild1_1.AddChild(greatgrandchild1_1_1)

	// Adding child comments to the root
	root.AddChild(child1)
	root.AddChild(child2)

	// Print the entire comment tree
	fmt.Println("Comment Tree:")
	root.Print()
}
