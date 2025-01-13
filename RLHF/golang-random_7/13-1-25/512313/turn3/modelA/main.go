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
		close(transformedValues)
	}()

	// Limit concurrent processing by using a fixed number of workers
	workerCount := 4 // Adjust this based on your needs
	workerCh := make(chan struct{}, workerCount)

	// Perform the concurrent in-order traversal
	wg.Add(1)
	go func() {
		inorderTraversal(root, transformedValues, workerCh, &wg)
		wg.Done()
	}()

	// Wait for the traversal to complete
	wg.Wait()
}

func inorderTraversal(node *TreeNode, transformedValues chan<- int, workerCh chan struct{}, wg *sync.WaitGroup) {
	if node == nil {
		return
	}

	// Send a work token to wait for a available worker
	workerCh <- struct{}{}
	defer func() {
		<-workerCh
	}()

	wg.Add(1)
	go func() {
		inorderTraversal(node.Left, transformedValues, workerCh, wg)
		wg.Done()
	}()

	// Process current node and double its value
	transformedValues <- node.Val * 2

	wg.Add(1)
	go func() {
		inorderTraversal(node.Right, transformedValues, workerCh, wg)
		wg.Done()
	}()
}
