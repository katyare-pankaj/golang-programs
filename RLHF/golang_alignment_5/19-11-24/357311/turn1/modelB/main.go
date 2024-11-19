
// Package bidirectsearch provides a bidirectional search algorithm for graph algorithms.
package bidirectsearch

import "fmt"

// Graph represents a general graph data structure.
type Graph struct {
	Nodes []int
	Edges map[int][]int
}

// NewGraph creates a new empty graph.
func NewGraph() *Graph {
	return &Graph{
		Edges: make(map[int][]int),
	}
}

// AddNode adds a new node to the graph.
func (g *Graph) AddNode(node int) {
	if _, ok := g.Edges[node]; !ok {
		g.Nodes = append(g.Nodes, node)
		g.Edges[node] = []int{}
	}
}

// AddEdge adds an edge between two nodes in the graph.
func (g *Graph) AddEdge(from, to int) {
	g.AddNode(from)
	g.AddNode(to)
	g.Edges[from] = append(g.Edges[from], to)
}

// BidirectionalSearch finds a path between two nodes in the graph using bidirectional search.
func BidirectionalSearch(graph *Graph, start, goal int) []int {
	frontier1 := map[int]struct{}{start: {}}
	frontier2 := map[int]struct{}{goal: {}}
	parents1 := map[int]int{start: -1}
	parents2 := map[int]int{goal: -1}
	visited1 := map[int]struct{}{}
	visited2 := map[int]struct{}{}

	for len(frontier1) > 0 && len(frontier2) > 0 {
		// Expand the frontiers in both directions
		if len(frontier1) < len(frontier2) {
			node := expandFrontier(graph, frontier1, parents1, visited1, frontier2, parents2, visited2)
			if node != -1 {
				return reconstructPath(parents1, parents2, start, goal, node)
			}
		} else {
			node := expandFrontier(graph, frontier2, parents2, visited2, frontier1, parents1, visited1)
			if node != -1 {
				return reconstructPath(parents1, parents2, start, goal, node)
			}
		}
	}

	// If no path is found, return an empty slice
	return []int{}
}

func expandFrontier(graph *Graph, frontier1, parents1, visited1 map[int]struct{}, frontier2, parents2, visited2 map[int]struct{}) int {
	node := -1
	for key := range frontier1 {
		node = key
		break
	}
	delete(frontier1, node)
	visited1[node] = struct{}{}

	for _, neighbor := range graph.Edges[node] {
		if _, ok := visited2[neighbor]; ok {
			return neighbor // We have found a meeting point!
		}
		if _, ok := visited1[neighbor]; !ok {
			frontier1[neighbor] = struct{}{}
			parents1[neighbor] = node
		}
	}
	return -1
}

func reconstructPath(parents1, parents2 map[int]int, start, goal, meetingPoint int) []int {
	// Reconstruct the path from start to meetingPoint using parents1
	path1 := []int{}
	current := meetingPoint
	for current != start {
		path1 = append(path1, current)
		current = parents1[current]
	}
	path1 = append(path1, start)
	reverse(path1)

	// Reconstruct the path from meetingPoint to goal using parents2
	path2 := []int{}
	current = meetingPoint