package main

import (
	"fmt"
)

// Graph represents an undirected graph using an adjacency list.
type Graph struct {
	vertices map[string]*Vertex
}

// NewGraph creates a new graph instance.
func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[string]*Vertex),
	}
}

// Vertex represents a vertex in the graph.
type Vertex struct {
	id      string
	edges   []*Vertex
	visited bool
}

// NewVertex creates a new vertex with a given ID.
func NewVertex(id string) *Vertex {
	return &Vertex{
		id:      id,
		edges:   make([]*Vertex, 0),
		visited: false,
	}
}

// AddVertex adds a new vertex to the graph.
func (g *Graph) AddVertex(v *Vertex) {
	g.vertices[v.id] = v
}

// AddEdge adds an edge between two vertices.
func (g *Graph) AddEdge(from *Vertex, to *Vertex) {
	from.edges = append(from.edges, to)
	to.edges = append(to.edges, from) // Since the graph is undirected
}

// BFS performs a breadth-first search starting from the given vertex.
// Returns a list of vertices in the order they are visited.
func (g *Graph) BFS(startVertex string) []string {
	start, exists := g.vertices[startVertex]
	if !exists {
		fmt.Printf("Vertex %s does not exist in the graph.\n", startVertex)
		return nil
	}

	queue := []*Vertex{start}
	start.visited = true
	result := []string{}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current.id)

		for _, neighbor := range current.edges {
			if !neighbor.visited {
				queue = append(queue, neighbor)
				neighbor.visited = true
			}
		}
	}

	return result
}

func main() {
	graph := NewGraph()

	// Adding vertices
	v1 := NewVertex("A")
	v2 := NewVertex("B")
	v3 := NewVertex("C")
	v4 := NewVertex("D")
	v5 := NewVertex("E")
	v6 := NewVertex("F")

	// Adding edges
	graph.AddEdge(v1, v2)
	graph.AddEdge(v1, v3)
	graph.AddEdge(v2, v4)
	graph.AddEdge(v2, v5)
	graph.AddEdge(v3, v6)
	graph.AddEdge(v4, v6)
	graph.AddEdge(v5, v6)

	// Performing BFS starting from vertex A
	order := graph.BFS("A")
	fmt.Println("BFS Traversal Order:", order)
}
