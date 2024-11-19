package main

import (
	"container/list"
	"fmt"
)

// Graph represents the graph structure
type Graph struct {
	NumNodes int
	AdjList  map[int][]int
}

// NewGraph initializes a new graph with a specified number of nodes
func NewGraph(numNodes int) *Graph {
	return &Graph{
		NumNodes: numNodes,
		AdjList:  make(map[int][]int),
	}
}

// AddEdge adds a directed edge between source and destination nodes
func (g *Graph) AddEdge(source int, destination int) {
	g.AdjList[source] = append(g.AdjList[source], destination)
}

// BFS performs BFS starting from a given source node and prints the visited nodes
func (g *Graph) BFS(source int) {
	queue := list.New()
	visited := make([]bool, g.NumNodes)

	// Mark the source node as visited and enqueue it
	visited[source] = true
	queue.PushBack(source)

	for queue.Len() != 0 {
		current := queue.Remove(queue.Front()).(int)
		fmt.Printf("%d ", current)

		// Explore all neighbors of the current node
		for _, neighbor := range g.AdjList[current] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue.PushBack(neighbor)
			}
		}
	}
}

func main() {
	g := NewGraph(4)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 3)

	fmt.Println("BFS Traversal:")
	g.BFS(2)
}
