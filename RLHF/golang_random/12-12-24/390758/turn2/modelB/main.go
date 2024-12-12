package main

import (
	"fmt"
	"sync"
)

// Node represents a single node in the BST
type Node struct {
	key   int
	left  *Node
	right *Node
	lock  *sync.RWMutex // Lock for concurrent access
}

// ConcurrentBST represents the BST
type ConcurrentBST struct {
	root *Node
}

// NewNode creates a new node
func NewNode(key int) *Node {
	return &Node{key: key, lock: &sync.RWMutex{}}
}

// Insert inserts a new node into the BST
func (t *ConcurrentBST) Insert(key int) {
	t.root = insert(t.root, key)
}

func insert(node *Node, key int) *Node {
	if node == nil {
		return NewNode(key)
	}

	node.lock.RLock()
	defer node.lock.RUnlock()

	if key < node.key {
		node.left = insert(node.left, key)
	} else if key > node.key {
		node.right = insert(node.right, key)
	}

	node.lock.Lock()
	defer node.lock.Unlock()

	// Here you would typically have balancing operations

	return node
}

// Search searches for a key in the BST
func (t *ConcurrentBST) Search(key int) bool {
	return search(t.root, key)
}

func search(node *Node, key int) bool {
	if node == nil {
		return false
	}

	node.lock.RLock()
	defer node.lock.RUnlock()

	if key == node.key {
		return true
	} else if key < node.key {
		return search(node.left, key)
	} else {
		return search(node.right, key)
	}
}

// InOrder traverses the BST in-order
func (t *ConcurrentBST) InOrder() {
	inOrder(t.root)
}

func inOrder(node *Node) {
	if node != nil {
		node.lock.RLock()
		defer node.lock.RUnlock()
		inOrder(node.left)
		fmt.Print(node.key, " ")
		inOrder(node.right)
	}
}

func main() {
	cbst := &ConcurrentBST{}
	cbst.Insert(10)
	cbst.Insert(20)
	cbst.Insert(5)
	cbst.Insert(15)
	cbst.Insert(30)

	fmt.Println("In-order Traversal:")
	cbst.InOrder() // Output might not be balanced since locks do not enforce balance

	fmt.Println("\nSearch for 15:", cbst.Search(15))
	fmt.Println("Search for 100:", cbst.Search(100))
}
