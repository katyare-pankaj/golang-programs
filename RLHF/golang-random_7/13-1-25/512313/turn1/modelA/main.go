package main

import (
	"fmt"
)

// TreeNode represents a node in the binary tree.
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// InOrderTraversalWithTransform performs an in-order traversal of a binary tree,
// transforms the value of each node, and returns the transformed values.
func InOrderTraversalWithTransform(root *TreeNode) []int {
	var transformedValues []int

	func traverse(node *TreeNode) {
		if node == nil {
			return
		}

		traverse(node.Left)
		// Transform the value and add it to the list
		transformedValues = append(transformedValues, node.Value*2)
		traverse(node.Right)
	}

	traverse(root)
	return transformedValues
}

func main() {
	// Creating a simple binary tree
	//       1
	//      / \
	//     2   3
	//    / \
	//   4   5
	root := &TreeNode{
		Value: 1,
		Left: &TreeNode{
			Value: 2,
			Left:  &TreeNode{Value: 4},
			Right: &TreeNode{Value: 5},
		},
		Right: &TreeNode{
			Value: 3,
		},
	}

	// Perform the transformation and get the result
	transformedValues := InOrderTraversalWithTransform(root)

	// Display the transformed values
	for _, value := range transformedValues {
		fmt.Println(value)
	}
}