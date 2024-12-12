package main

import (
	"fmt"
	"sync"
)

type Node struct {
	key   int
	left  *Node
	right *Node
	mu    sync.Mutex // Node-level lock
}

// ConcurrentBST is the structure for the concurrent binary search tree
type ConcurrentBST struct {
	root *Node
}

func NewConcurrentBST() *ConcurrentBST {
	return &ConcurrentBST{}
}

// Insert inserts a new key into the BST
func (tree *ConcurrentBST) Insert(key int) {
	if tree.root == nil {
		tree.root = &Node{key: key}
		return
	}
	tree.insert(tree.root, key)
}

func (tree *ConcurrentBST) insert(node *Node, key int) {
	node.mu.Lock()
	defer node.mu.Unlock()

	if key < node.key {
		if node.left == nil {
			node.left = &Node{key: key}
		} else {
			tree.insert(node.left, key)
		}
	} else if key > node.key {
		if node.right == nil {
			node.right = &Node{key: key}
		} else {
			tree.insert(node.right, key)
		}
	}
}

// Search checks whether a key is present in the BST
func (tree *ConcurrentBST) Search(key int) bool {
	return tree.search(tree.root, key)
}

func (tree *ConcurrentBST) search(node *Node, key int) bool {
	if node == nil {
		return false
	}

	// First lock the node to prevent modification during read
	node.mu.Lock()
	defer node.mu.Unlock()

	if key < node.key {
		return tree.search(node.left, key)
	} else if key > node.key {
		return tree.search(node.right, key)
	}
	return true // Key found
}

// InOrder traverses the tree in-order
func (tree *ConcurrentBST) InOrder() {
	tree.inOrder(tree.root)
}

func (tree *ConcurrentBST) inOrder(node *Node) {
	if node == nil {
		return
	}

	// Note: In this traversal, you'll not use locks to allow read-only access for traversal.
	// However, this should be used carefully and may not be suitable for truly concurrent writes.
	tree.inOrder(node.left)
	fmt.Print(node.key, " ")
	tree.inOrder(node.right)
}

// Remove deletes a key from the BST
// Note: We need to implement the removal in a thread-safe manner.
func (tree *ConcurrentBST) Remove(key int) {
	tree.root = tree.remove(tree.root, key)
}

func (tree *ConcurrentBST) remove(node *Node, key int) *Node {
	if node == nil {
		return nil
	}

	node.mu.Lock()
	defer node.mu.Unlock()

	if key < node.key {
		node.left = tree.remove(node.left, key)
	} else if key > node.key {
		node.right = tree.remove(node.right, key)
	} else {
		// Node with only one child or no child
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}
		// Node with two children: get the inorder successor (smallest in the right subtree)
		minNode := node.right
		for minNode != nil && minNode.left != nil {
			minNode = minNode.left
		}
		node.key = minNode.key
		node.right = tree.remove(node.right, minNode.key)
	}
	return node
}

// Main for testing
func main() {
	tree := NewConcurrentBST()
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(4)

	fmt.Println("In-order traversal of BST:")
	tree.InOrder()

	fmt.Println("\nSearching for key 4:", tree.Search(4))
	fmt.Println("Searching for key 6:", tree.Search(6))

	tree.Remove(3)
	fmt.Println("In-order traversal after removing 3:")
	tree.InOrder()
}
