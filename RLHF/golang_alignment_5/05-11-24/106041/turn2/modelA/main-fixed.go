package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Node struct {
	Value int   `json:"value"`
	Next  *Node `json:"next"`
}

// SimpleNodeID represents a simple node structure with ID
type SimpleNodeID struct {
	ID int `json:"id"`
}

// SimpleNode represents the JSON structure for the Node
type SimpleNode struct {
	Value int           `json:"value"`
	Next  *SimpleNodeID `json:"next"`
}

// Custom Marshaler to break the circular reference
func (n *Node) MarshalJSON() ([]byte, error) {
	if n == nil {
		return json.Marshal(nil)
	}

	// Map to store node references
	nodeMap := make(map[*Node]int)
	nodeMap[n] = 0

	// Create a SimpleNode for JSON output
	out := SimpleNode{
		Value: n.Value,
		Next:  marshalNext(n.Next, nodeMap),
	}

	return json.Marshal(out)
}

// Helper function to create references to nodes by ID to avoid circular references
func marshalNext(node *Node, nodeMap map[*Node]int) *SimpleNodeID {
	if node == nil {
		return nil
	}
	id, ok := nodeMap[node]
	if !ok {
		id = len(nodeMap)
		nodeMap[node] = id
	}
	return &SimpleNodeID{ID: id}
}

func main() {
	// Example data with a circular reference
	a := &Node{Value: 1}
	b := &Node{Value: 2}
	c := &Node{Value: 3}

	a.Next = b
	b.Next = c
	c.Next = a // Circular reference back to A

	data, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		log.Fatal("Error marshaling data:", err)
	}
	fmt.Println(string(data))
}
