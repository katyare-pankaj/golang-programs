package main

import (
	"fmt"
	"sync"
)

// Node represents a single node in the AVL tree
type Node struct {
	key    int
	height int
	left   *Node
	right  *Node
}

// AVLTree represents the AVL tree
type AVLTree struct {
	root *Node
	mu   sync.RWMutex // Mutex for concurrent access
}

// NewNode creates a new node
func NewNode(key int) *Node {
	return &Node{key: key, height: 1}
}

// Height gets the height of a node
func Height(node *Node) int {
	if node == nil {
		return 0
	}
	return node.height
}

// RightRotate performs a right rotation on the subtree rooted with node
func RightRotate(y *Node) *Node {
	x := y.left
	T2 := x.right

	// Perform rotation
	x.right = y
	y.left = T2

	// Update heights
	y.height = max(Height(y.left), Height(y.right)) + 1
	x.height = max(Height(x.left), Height(x.right)) + 1

	return x
}

// LeftRotate performs a left rotation on the subtree rooted with node
func LeftRotate(x *Node) *Node {
	y := x.right
	T2 := y.left

	// Perform rotation
	y.left = x
	x.right = T2

	// Update heights
	x.height = max(Height(x.left), Height(x.right)) + 1
	y.height = max(Height(y.left), Height(y.right)) + 1

	return y
}

// GetBalance gets the balance factor of the node
func GetBalance(node *Node) int {
	if node == nil {
		return 0
	}
	return Height(node.left) - Height(node.right)
}

// Insert inserts a node and maintains AVL Tree properties
func (t *AVLTree) Insert(key int) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.root = insert(t.root, key)
}

func insert(node *Node, key int) *Node {
	if node == nil {
		return NewNode(key)
	}

	if key < node.key {
		node.left = insert(node.left, key)
	} else if key > node.key {
		node.right = insert(node.right, key)
	} else {
		return node // Duplicate keys are not allowed
	}

	node.height = 1 + max(Height(node.left), Height(node.right))

	balance := GetBalance(node)

	// Left Left Case
	if balance > 1 && key < node.left.key {
		return RightRotate(node)
	}

	// Right Right Case
	if balance < -1 && key > node.right.key {
		return LeftRotate(node)
	}

	// Left Right Case
	if balance > 1 && key > node.left.key {
		node.left = LeftRotate(node.left)
		return RightRotate(node)
	}

	// Right Left Case
	if balance < -1 && key < node.right.key {
		node.right = RightRotate(node.right)
		return LeftRotate(node)
	}

	return node
}

// Max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// InOrder traverses the tree in-order
func (t *AVLTree) InOrder() {
	t.mu.RLock()
	defer t.mu.RUnlock()

	inOrder(t.root)
}

func inOrder(node *Node) {
	if node != nil {
		inOrder(node.left)
		fmt.Print(node.key, " ")
		inOrder(node.right)
	}
}

func main() {
	avl := AVLTree{}
	avl.Insert(10)
	avl.Insert(20)
	avl.Insert(30)
	avl.Insert(40)
	avl.Insert(50)

	avl.InOrder() // Output: 10 20 30 40 50
}
