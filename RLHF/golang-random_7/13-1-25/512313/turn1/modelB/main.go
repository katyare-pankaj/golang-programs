package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Construct the binary tree
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}

	// In-order traversal using range loops
	transformedValues := make([]int, 0)
	for node := range inorderTraversal(root) {
		transformedValues = append(transformedValues, node.Val*2)
	}

	// Display the transformed values
	fmt.Println("Transformed values in order:")
	for _, val := range transformedValues {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}

func inorderTraversal(root *TreeNode) <-chan *TreeNode {
	c := make(chan *TreeNode)

	go func() {
		var stack []*TreeNode
		var current *TreeNode = root

		for current != nil || len(stack) > 0 {
			for current != nil {
				stack = append(stack, current)
				current = current.Left
			}

			current = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			c <- current

			current = current.Right
		}

		close(c)
	}()

	return c
}
