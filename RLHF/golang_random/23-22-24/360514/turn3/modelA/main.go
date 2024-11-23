package main

import (
	"fmt"
)

// Vertex interface defines the operations that a vertex must support.
type Vertex interface {
	GetID() int
	SetID(int)
	AddEdge(Edge)
	GetEdges() []Edge
	RemoveEdge(Edge)
}

// Edge interface defines the operations that an edge must support.
type Edge interface {
	GetSource() Vertex
	GetDestination() Vertex
	SetWeight(int)
	GetWeight() int
}

// Graph struct represents the graph with a map of vertices.
type Graph struct {
	vertices map[int]Vertex
}

func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[int]Vertex),
	}
}

func (g *Graph) AddVertex(vertex Vertex) {
	g.vertices[vertex.GetID()] = vertex
}

func (g *Graph) GetVertex(id int) Vertex {
	return g.vertices[id]
}

func (g *Graph) RemoveVertex(id int) {
	delete(g.vertices, id)
}

func (g *Graph) TraverseGraph() {
	for _, vertex := range g.vertices {
		fmt.Printf("Vertex %d:\n", vertex.GetID())
		for _, edge := range vertex.GetEdges() {
			fmt.Printf("  -> Vertex %d (Weight: %d)\n", edge.GetDestination().GetID(), edge.GetWeight())
		}
	}
}

// Simple Vertex implementation with basic properties.
type BasicVertex struct {
	id    int
	edges []Edge
}

func NewBasicVertex(id int) *BasicVertex {
	return &BasicVertex{
		id:    id,
		edges: nil,
	}
}

func (v *BasicVertex) GetID() int {
	return v.id
}

func (v *BasicVertex) SetID(id int) {
	v.id = id
}

func (v *BasicVertex) AddEdge(edge Edge) {
	v.edges = append(v.edges, edge)
}

func (v *BasicVertex) GetEdges() []Edge {
	return v.edges
}

func (v *BasicVertex) RemoveEdge(edge Edge) {
	for i, e := range v.edges {
		if e == edge {
			v.edges = append(v.edges[:i], v.edges[i+1:]...)
			return
		}
	}
}

// Simple Edge implementation with basic properties.
type BasicEdge struct {
	source      Vertex
	destination Vertex
	weight      int
}

func NewBasicEdge(source Vertex, destination Vertex, weight int) *BasicEdge {
	return &BasicEdge{
		source:      source,
		destination: destination,
		weight:      weight,
	}
}

func (e *BasicEdge) GetSource() Vertex {
	return e.source
}

func (e *BasicEdge) GetDestination() Vertex {
	return e.destination
}

func (e *BasicEdge) SetWeight(weight int) {
	e.weight = weight
}

func (e *BasicEdge) GetWeight() int {
	return e.weight
}

func main() {
	graph := NewGraph()

	vertex1 := NewBasicVertex(1)
	vertex2 := NewBasicVertex(2)
	vertex3 := NewBasicVertex(3)

	edge12 := NewBasicEdge(vertex1, vertex2, 5)
	edge13 := NewBasicEdge(vertex1, vertex3, 3)
	edge23 := NewBasicEdge(vertex2, vertex3, 4)

	graph.AddVertex(vertex1)
	graph.AddVertex(vertex2)
	graph.AddVertex(vertex3)

	vertex1.AddEdge(edge12)
	vertex1.AddEdge(edge13)
	vertex2.AddEdge(edge23)

	fmt.Println("Graph Traversal:")
	graph.TraverseGraph()

	fmt.Println("\nRemoving vertex 3:")
	graph.RemoveVertex(vertex3.GetID())
	graph.TraverseGraph()
}
