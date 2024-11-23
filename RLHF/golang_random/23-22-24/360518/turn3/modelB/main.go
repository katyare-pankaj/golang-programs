package main

import (
	"fmt"
)

type VertexID int
type Vertex interface {
	GetID() VertexID
	AddEdge(Edge)
	GetEdges() []Edge
	RemoveEdge(VertexID)
}
type Edge interface {
	GetSourceID() VertexID
	GetDestinationID() VertexID
	GetWeight() int
}
type Graph interface {
	AddVertex(Vertex)
	RemoveVertex(VertexID)
	GetVertex(VertexID) Vertex
	HasVertex(VertexID) bool
	IterateVertices(func(Vertex))
}
type BasicVertex struct {
	id    VertexID
	edges map[VertexID]Edge
}

func NewBasicVertex(id VertexID) Vertex {
	return &BasicVertex{
		id:    id,
		edges: make(map[VertexID]Edge),
	}
}
func (v *BasicVertex) GetID() VertexID {
	return v.id
}
func (v *BasicVertex) AddEdge(edge Edge) {
	v.edges[edge.GetDestinationID()] = edge
}
func (v *BasicVertex) GetEdges() []Edge {
	edges := make([]Edge, 0, len(v.edges))
	for _, edge := range v.edges {
		edges = append(edges, edge)
	}
	return edges
}
func (v *BasicVertex) RemoveEdge(destID VertexID) {
	delete(v.edges, destID)
}

type BasicEdge struct {
	sourceID      VertexID
	destinationID VertexID
	weight        int
}

func NewBasicEdge(sourceID, destinationID VertexID, weight int) Edge {
	return &BasicEdge{
		sourceID:      sourceID,
		destinationID: destinationID,
		weight:        weight,
	}
}
func (e *BasicEdge) GetSourceID() VertexID {
	return e.sourceID
}
func (e *BasicEdge) GetDestinationID() VertexID {
	return e.destinationID
}
func (e *BasicEdge) GetWeight() int {
	return e.weight
}

type AdjacencyListGraph struct {
	vertices map[VertexID]Vertex
}

func NewAdjacencyListGraph() Graph {
	return &AdjacencyListGraph{
		vertices: make(map[VertexID]Vertex),
	}
}
func (g *AdjacencyListGraph) AddVertex(vertex Vertex) {
	g.vertices[vertex.GetID()] = vertex
}
func (g *AdjacencyListGraph) RemoveVertex(vertexID VertexID) {
	delete(g.vertices, vertexID)
	// Remove edges incident on this vertex
	for _, vertex := range g.vertices {
		vertex.RemoveEdge(vertexID)
	}
}
func (g *AdjacencyListGraph) GetVertex(vertexID VertexID) Vertex {
	return g.vertices[vertexID]
}
func (g *AdjacencyListGraph) HasVertex(vertexID VertexID) bool {
	_, ok := g.vertices[vertexID]
	return ok
}
func (g *AdjacencyListGraph) IterateVertices(iterate func(Vertex)) {
	for _, vertex := range g.vertices {
		iterate(vertex)
	}
}
func main() {
	graph := NewAdjacencyListGraph()
	// Add vertices dynamically
	graph.AddVertex(NewBasicVertex(1))
	graph.AddVertex(NewBasicVertex(2))
	graph.AddVertex(NewBasicVertex(3))
	// Add edges
	graph.GetVertex(1).AddEdge(NewBasicEdge(1, 2, 5))
	graph.GetVertex(1).AddEdge(NewBasicEdge(1, 3, 3))
	graph.GetVertex(2).AddEdge(NewBasicEdge(2, 3, 4))
	// Traverse the graph
	fmt.Println("Graph Traversal:")
	graph.IterateVertices(func(vertex Vertex) {
		fmt.Printf("Vertex %d:\n", vertex.GetID())
		for _, edge := range vertex.GetEdges() {
			fmt.Printf("  -> Vertex %d (Weight: %d)\n", edge.GetDestinationID(), edge.GetWeight())
		}
	})
	// Remove a vertex
	graph.RemoveVertex(3)
	fmt.Println("\nGraph after removing vertex 3:")
	graph.IterateVertices(func(vertex Vertex) {
		fmt.Printf("Vertex %d:\n", vertex.GetID())
		for _, edge := range vertex.GetEdges() {
			fmt.Printf("  -> Vertex %d (Weight: %d)\n", edge.GetDestinationID(), edge.GetWeight())
		}
	})
}
