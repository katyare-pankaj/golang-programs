package main

import (
	"errors"
	"fmt"
	"sync"
)

// Graph represents an undirected graph using an adjacency list.
type Graph struct {
	vertices map[string]*Vertex
	mu       sync.RWMutex
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
	edges   []*Edge
	visited bool
	mu      sync.RWMutex
}

// NewVertex creates a new vertex with a given ID.
func NewVertex(id string) *Vertex {
	return &Vertex{
		id:      id,
		edges:   make([]*Edge, 0),
		visited: false,
	}
}

// Edge represents an edge connecting two vertices.
type Edge struct {
	from   *Vertex
	to     *Vertex
	weight float64 // Optional field for weighted edges
}

// NewEdge creates a new edge connecting two vertices with an optional weight.
func NewEdge(from *Vertex, to *Vertex, weight float64) *Edge {
	return &Edge{
		from:   from,
		to:     to,
		weight: weight,
	}
}

// AddVertex adds a new vertex to the graph.
func (g *Graph) AddVertex(v *Vertex) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.vertices[v.id] = v
}

// AddEdge adds an edge between two vertices.
func (g *Graph) AddEdge(from *Vertex, to *Vertex) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if _, ok := g.vertices[from.id]; !ok {
		return errors.New("vertex '" + from.id + "' not found")
	}

	if _, ok := g.vertices[to.id]; !ok {
		return errors.New("vertex '" + to.id + "' not found")
	}

	edge := NewEdge(from, to, 0)
	from.edges = append(from.edges, edge)
	to.edges = append(to.edges, edge) // Since the graph is undirected
	return nil
}

// BidirectionalSearch performs a bidirectional search from start to end vertex.
// Returns a list of vertices in the shortest path or an error if no path exists.
func (g *Graph) BidirectionalSearch(start string, end string) ([]string, error) {
	startVertex, endVertex, err := g.getVertices(start, end)
	if err != nil {
		return nil, err
	}

	queueStart := []*Vertex{startVertex}
	queueEnd := []*Vertex{endVertex}

	parentStart := make(map[*Vertex]*Vertex)
	parentEnd := make(map[*Vertex]*Vertex)

	visitedStart := make(map[string]bool)
	visitedEnd := make(map[string]bool)

	var meetingPoint *Vertex
	var wg sync.WaitGroup
	wg.Add(2)

	// Start search from start vertex
	go func() {
		defer wg.Done()
		for len(queueStart) > 0 {
			current := queueStart[0]
			queueStart = queueStart[1:]

			visitedStart[current.id] = true

			for _, edge := range current.edges {
				neighbor := edge.to
				if !visitedStart[neighbor.id] {
					visitedStart[neighbor.id] = true
					parentStart[neighbor] = current
					queueStart = append(queueStart, neighbor)
					if visitedEnd[neighbor.id] {
						meetingPoint = neighbor
						return
					}
				}
			}
		}
	}()

	// Start search from end vertex
	go func() {
		defer wg.Done()
		for len(queueEnd) > 0 {
			current := queueEnd[0]
			queueEnd = queueEnd[1:]

			visitedEnd[current.id] = true

			for _, edge := range current.edges {
				neighbor := edge.from
				if !visitedEnd[neighbor.id] {
					visitedEnd[neighbor.id] = true
					parentEnd[neighbor] = current
					queueEnd = append(queueEnd, neighbor)
					if visitedStart[neighbor.id] {
						meetingPoint = neighbor
						return
					}
				}
			}
		}
	}()

	wg.Wait()

	if meetingPoint == nil {
		return nil, errors.New("no path found")
	}

	// Reconstruct the path from start to meeting point
	var pathStart []string
	for v := meetingPoint; v != nil; v = parentStart[v] {
		pathStart = append([]string{v.id}, pathStart...)
	}

	// Reconstruct the path from end to meeting point
	var pathEnd []string
	for v := meetingPoint; v != nil; v = parentEnd[v] {
		pathEnd = append(pathEnd, v.id)
	}

	// Combine the paths and remove the meeting point from the second part
	pathStart = append(pathStart, pathEnd[1:]...)

	return pathStart, nil
}

func (g *Graph) getVertices(start string, end string) (*Vertex, *Vertex, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	startVertex, ok := g.vertices[start]
	if !ok {
		return nil, nil, errors.New("vertex '" + start + "' not found")
	}

	endVertex, ok := g.vertices[end]
	if !ok {
		return nil, nil, errors.New("vertex '" + end + "' not found")
	}

	return startVertex, endVertex, nil
}

func main() {
	// Create a graph
	graph := NewGraph()

	// Add vertices
	A := NewVertex("A")
	B := NewVertex("B")
	C := NewVertex("C")
	D := NewVertex("D")
	graph.AddVertex(A)
	graph.AddVertex(B)
	graph.AddVertex(C)
	graph.AddVertex(D)

	// Add edges
	graph.AddEdge(A, B)
	graph.AddEdge(B, D)
	graph.AddEdge(B, C)

	// Perform bidirectional search
	start := "A"
	end := "D"
	path, err := graph.BidirectionalSearch(start, end)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Shortest path:", path)
	}
}
