package main

import (
	"fmt"
	"runtime"
	"sync"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}

	transformedValues := make(chan int)
	var wg sync.WaitGroup

	go receiver(transformedValues)
	inorderTraversal(root, transformedValues, &wg)
	wg.Wait()
	close(transformedValues)
}

func receiver(transformedValues <-chan int) {
	for val := range transformedValues {
		fmt.Printf("%d ", val)
	}
}

func inorderTraversal(node *TreeNode, transformedValues chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	if node == nil {
		return
	}

	wg.Add(2)
	go func() {
		inorderTraversal(node.Left, transformedValues, wg)
		wg.Done()
	}()

	go func() {
		transformedValues <- node.Val * 2
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		inorderTraversal(node.Right, transformedValues, wg)
		wg.Done()
	}()
}
