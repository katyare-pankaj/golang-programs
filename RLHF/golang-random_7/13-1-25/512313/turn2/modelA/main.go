package main

import (
	"fmt"
	"sync"
)

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

	// Create a channel to receive transformed values
	resultCh := make(chan int)

	// Create a wait group to synchronize the main goroutine
	var wg sync.WaitGroup

	// Start a goroutine for the in-order traversal
	wg.Add(1)
	go func() {
		inorderTraversal(root, resultCh)
		wg.Done()
	}()

	// Collect the transformed values from the channel
	var transformedValues []int
	for val := range resultCh {
		transformedValues = append(transformedValues, val)
	}

	// Wait for the traversal goroutine to finish
	wg.Wait()

	// Display the transformed values
	fmt.Println("Transformed values in order:")
	for _, val := range transformedValues {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}

func inorderTraversal(root *TreeNode, resultCh chan<- int) {
	if root == nil {
		return
	}

	inorderTraversal(root.Left, resultCh)

	// Double the value and send it to the channel
	resultCh <- root.Val * 2

	inorderTraversal(root.Right, resultCh)
}
