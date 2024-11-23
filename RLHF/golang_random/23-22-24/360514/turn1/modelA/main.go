package main

import "fmt"

// Node interface defines the methods that nodes in the graph must implement.
type Node interface {
	ID() int
	Label() string
	Connections() []Node
}

// SimpleNode is a concrete implementation of the Node interface for this example.
type SimpleNode struct {
	id    int
	label string
	conn  []Node
}

func (s *SimpleNode) ID() int             { return s.id }
func (s *SimpleNode) Label() string       { return s.label }
func (s *SimpleNode) Connections() []Node { return s.conn }

// Graph represents the graph structure using adjacency list.
type Graph struct {
	nodes map[int]Node
}

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[int]Node),
	}
}

// AddNode adds a new node to the graph.
func (g *Graph) AddNode(id int, label string) {
	n := &SimpleNode{id: id, label: label}
	g.nodes[id] = n
}

// AddEdge adds an edge between two nodes in the graph.
func (g *Graph) AddEdge(id1, id2 int) {
	node1, ok := g.nodes[id1]
	node2, ok := g.nodes[id2]
	if !ok || !ok {
		fmt.Printf("Error: Nodes with IDs %d and %d do not exist.\n", id1, id2)
		return
	}

	node1.(*SimpleNode).conn = append(node1.(*SimpleNode).conn, node2)
	node2.(*SimpleNode).conn = append(node2.(*SimpleNode).conn, node1) // Undirected graph
}

// PrintGraph prints the graph in a readable format.
func (g *Graph) PrintGraph() {
	for _, node := range g.nodes {
		fmt.Printf("Node %d: %s -> ", node.ID(), node.Label())
		for _, conn := range node.Connections() {
			fmt.Printf("%d ", conn.ID())
		}
		fmt.Println()
	}
}

func main() {
	graph := NewGraph()

	graph.AddNode(1, "A")
	graph.AddNode(2, "B")
	graph.AddNode(3, "C")
	graph.AddNode(4, "D")

	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)

	fmt.Println("Graph:")
	graph.PrintGraph()
}
