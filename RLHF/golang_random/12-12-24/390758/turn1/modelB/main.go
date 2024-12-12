package main

import (
	"fmt"
	"sync"
)

type Node struct {
	Key     int
	Left    *Node
	Right   *Node
	Height  int
	Balance int // -1, 0, 1
	rwmutex *sync.RWMutex
}

func NewNode(key int) *Node {
	return &Node{
		Key:     key,
		Height:  1,
		Balance: 0,
		rwmutex: &sync.RWMutex{},
	}
}

func (root *Node) RightRotate() *Node {
	y := root.Left
	T2 := y.Right

	y.Right = root
	root.Left = T2

	root.Height = max(height(root.Left), height(root.Right)) + 1
	y.Height = max(height(y.Left), height(y.Right)) + 1

	// Update balance factor
	root.Balance = balance(root)
	y.Balance = balance(y)

	return y
}

func (root *Node) LeftRotate() *Node {
	x := root.Right
	T3 := x.Left

	x.Left = root
	root.Right = T3

	root.Height = max(height(root.Left), height(root.Right)) + 1
	x.Height = max(height(x.Left), height(x.Right)) + 1

	// Update balance factor
	root.Balance = balance(root)
	x.Balance = balance(x)

	return x
}

func height(node *Node) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func balance(node *Node) int {
	if node == nil {
		return 0
	}
	return height(node.Left) - height(node.Right)
}

func rightRotateOnInsert(y *Node) *Node {
	x := y.Left
	T2 := x.Right

	// Perform rotation
	x.Right = y
	y.Left = T2

	// Update heights
	y.Height = max(height(y.Left), height(y.Right)) + 1
	x.Height = max(height(x.Left), height(x.Right)) + 1

	// Update balance factors
	y.Balance = balance(y)

	if x.Balance == -1 && y.Balance == 1 {
		x.Balance = 0
		y.Balance = 0
	} else if x.Balance == -1 {
		y.Balance = 0
	} else if x.Balance == 1 && y.Balance == -1 {
		x.Balance = 0
	} else if x.Balance == 1 {
		x.Balance = 0
	}

	return x
}

func leftRotateOnInsert(x *Node) *Node {
	y := x.Right
	T3 := y.Left

	// Perform rotation
	y.Left = x
	x.Right = T3

	// Update heights
	x.Height = max(height(x.Left), height(x.Right)) + 1
	y.Height = max(height(y.Left), height(y.Right)) + 1

	// Update balance factors
	x.Balance = balance(x)

	if x.Balance == 1 && y.Balance == -1 {
		x.Balance = 0
		y.Balance = 0
	} else if x.Balance == 1 {
		y.Balance = 0
	} else if x.Balance == -1 && y.Balance == 1 {
		x.Balance = 0
	} else if x.Balance == -1 {
		x.Balance = 0
	}

	return y
}

func (root *Node) Insert(key int) {
	root.rwmutex.Lock()
	defer root.rwmutex.Unlock()

	if root == nil {
		root = NewNode(key)
		return
	}

	if key < root.Key {
		root.Left = root.InsertRec(root.Left, key)
	} else if key > root.Key {
		root.Right = root.InsertRec(root.Right, key)
	}

	// Update height of this ancestor node
	root.Height = max(height(root.Left), height(root.Right)) + 1

	// Get the balance factor of this ancestor node to check whether this node became unbalanced
	bf := balance(root)

	// If this node becomes unbalanced, then there are 4 cases

	// Left Left Case
	if bf > 1 && key < root.Left.Key {
		return
	}

	// Right Right Case
	if bf < -1 && key > root.Right.Key {
		root.Right = root.RightRotate()
		return
	}

	// Left Right Case
	if bf > 1 && key > root.Left.Key {
		root.Left = root.LeftRotate()
		root = root.RightRotate()
		return
	}

	// Right Left Case
	if bf < -1 && key < root.Right.Key {
		root.Right = root.RightRotate()
		root = root.LeftRotate()
		return
	}
}

func (root *Node) InsertRec(node *Node, key int) *Node {
	if node == nil {
		return NewNode(key)
	}

	if key < node.Key {
		node.Left = node.InsertRec(node.Left, key)
	} else if key > node.Key {
		node.Right = node.InsertRec(node.Right, key)
	}

	// Update height of this ancestor node
	node.Height = max(height(node.Left), height(node.Right)) + 1

	// Get the balance factor of this ancestor node to check whether this node became unbalanced
	bf := balance(node)

	// If this node becomes unbalanced, then there are 4 cases

	// Left Left Case
	if bf > 1 && key < node.Left.Key {
		return rightRotateOnInsert(node)
	}

	// Right Right Case
	if bf < -1 && key > node.Right.Key {
		return leftRotateOnInsert(node)
	}

	// Left Right Case
	if bf > 1 && key > node.Left.Key {
		node.Left = leftRotateOnInsert(node.Left)
		return rightRotateOnInsert(node)
	}

	// Right Left Case
	if bf < -1 && key < node.Right.Key {
		node.Right = rightRotateOnInsert(node.Right)
		return leftRotateOnInsert(node)
	}

	return node
}

func (root *Node) Search(key int) bool {
	root.rwmutex.RLock()
	defer root.rwmutex.RUnlock()

	if root == nil || root.Key == key {
		return true
	}

	if key < root.Key {
		return root.SearchRec(root.Left, key)
	}
	return root.SearchRec(root.Right, key)
}

func (root *Node) SearchRec(node *Node, key int) bool {
	if node == nil || node.Key == key {
		return true
	}

	if key < node.Key {
		return root.SearchRec(node.Left, key)
	}
	return root.SearchRec(node.Right, key)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	avlTree := NewNode(10)

	avlTree.Insert(20)
	avlTree.Insert(30)
	avlTree.Insert(40)
	avlTree.Insert(50)
	avlTree.Insert(25)

	fmt.Println("Search 20:", avlTree.Search(20)) // Output: true
	fmt.Println("Search 60:", avlTree.Search(60)) // Output: false
}
