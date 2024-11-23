package main

import (
	"fmt"
)

type Node interface {
	GetValue() int
	SetValue(int)
	AddChild(Node)
	GetChildren() []Node
}
type Tree struct {
	root Node
}

func NewTree() *Tree {
	return &Tree{
		root: nil,
	}
}
func (t *Tree) GetRoot() Node {
	return t.root
}
func (t *Tree) SetRoot(node Node) {
	t.root = node
}

// Simple Node implementation with basic properties
type BasicNode struct {
	value    int
	children []Node
}

func NewBasicNode(value int) *BasicNode {
	return &BasicNode{
		value:    value,
		children: nil,
	}
}
func (n *BasicNode) GetValue() int {
	return n.value
}
func (n *BasicNode) SetValue(value int) {
	n.value = value
}
func (n *BasicNode) AddChild(child Node) {
	n.children = append(n.children, child)
}
func (n *BasicNode) GetChildren() []Node {
	return n.children
}
func main() {
	tree := NewTree()
	root := NewBasicNode(1)
	tree.SetRoot(root)
	// Add children
	child1 := NewBasicNode(2)
	child2 := NewBasicNode(3)
	root.AddChild(child1)
	root.AddChild(child2)
	grandChild1 := NewBasicNode(4)
	child1.AddChild(grandChild1)
	// Traverse the tree
	fmt.Println("Tree Traversal:")
	traverseTree(tree.GetRoot())
}
func traverseTree(node Node) {
	fmt.Println(node.GetValue())
	for _, child := range node.GetChildren() {
		traverseTree(child)
	}
}
