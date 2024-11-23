package main

import (
	"fmt"
)

// Node interface defines the methods a graph node must implement.
type Node interface {
	GetValue() interface{}
	SetValue(interface{})
	AddEdge(Node)
	GetEdges() []Node
}

// Graph interface defines the methods a graph must implement.
type Graph interface {
	AddNode(Node)
	GetNodes() []Node
}

// SimpleNode implements the Node interface.
type SimpleNode struct {
	value interface{}
	edges []Node
}

func NewSimpleNode(value interface{}) *SimpleNode {
	return &SimpleNode{
		value: value,
		edges: nil,
	}
}

func (n *SimpleNode) GetValue() interface{} {
	return n.value
}

func (n *SimpleNode) SetValue(value interface{}) {
	n.value = value
}

func (n *SimpleNode) AddEdge(edge Node) {
	n.edges = append(n.edges, edge)
}

func (n *SimpleNode) GetEdges() []Node {
	return n.edges
}

// SimpleGraph implements the Graph interface.
type SimpleGraph struct {
	nodes []Node
}

func NewSimpleGraph() *SimpleGraph {
	return &SimpleGraph{
		nodes: nil,
	}
}

func (g *SimpleGraph) AddNode(node Node) {
	g.nodes = append(g.nodes, node)
}

func (g *SimpleGraph) GetNodes() []Node {
	return g.nodes
}

func main() {
	graph := NewSimpleGraph()

	// Create nodes
	node1 := NewSimpleNode("A")
	node2 := NewSimpleNode("B")
	node3 := NewSimpleNode("C")

	// Add nodes to the graph
	graph.AddNode(node1)
	graph.AddNode(node2)
	graph.AddNode(node3)

	// Add edges
	node1.AddEdge(node2)
	node1.AddEdge(node3)
	node2.AddEdge(node3)

	// Traverse the graph
	fmt.Println("Graph Traversal:")
	traverseGraph(graph)
}

func traverseGraph(graph Graph) {
	for _, node := range graph.GetNodes() {
		fmt.Println(node.GetValue())
		for _, edge := range node.GetEdges() {
			fmt.Println("  ->", edge.GetValue())
		}
	}
}
