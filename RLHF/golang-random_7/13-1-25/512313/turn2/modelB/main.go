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
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}

	transformedValues := make(chan int)
	var wg sync.WaitGroup

	// Start a concurrent receiver to read from the channel
	go func() {
		for val := range transformedValues {
			fmt.Printf("%d ", val)
		}
	}()

	// Perform the concurrent in-order traversal
	wg.Add(1)
	go func() {
		inorderTraversal(root, transformedValues, &wg)
		wg.Done()
	}()

	// Wait for the traversal to complete
	wg.Wait()
	close(transformedValues)

	fmt.Println()
}

func inorderTraversal(node *TreeNode, transformedValues chan<- int, wg *sync.WaitGroup) {
	if node == nil {
		return
	}

	wg.Add(1)
	go func() {
		inorderTraversal(node.Left, transformedValues, wg)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		// Double the node's value and send it to the channel
		transformedValues <- node.Val * 2
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		inorderTraversal(node.Right, transformedValues, wg)
		wg.Done()
	}()
}
