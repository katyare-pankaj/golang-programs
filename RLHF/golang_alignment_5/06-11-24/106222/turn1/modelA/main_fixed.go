package main

import (
	"fmt"
	"strings"
)

// CodeFragment represents a piece of code
type CodeFragment struct {
	content string
	lines   []int
}

// Graph represents the code base as a graph
type Graph struct {
	nodes []*CodeFragment
	edges [][]int
}

// NewGraph creates a new graph
func NewGraph() *Graph {
	return &Graph{nodes: make([]*CodeFragment, 0), edges: make([][]int, 0)}
}

// AddNode adds a node to the graph
func (g *Graph) AddNode(cf *CodeFragment) {
	g.nodes = append(g.nodes, cf)
	g.edges = append(g.edges, make([]int, len(g.nodes)))
}

// AddEdge adds an edge between two nodes
func (g *Graph) AddEdge(i, j int) {
	g.edges[i][j] = 1
	g.edges[j][i] = 1 // Undirected graph
}

// BronKerbosch finds cliques in an undirected graph
func BronKerbosch(g *Graph, r []int, p []int, x []int, cliques [][]int) [][]int {
	if len(p) == 0 && len(x) == 0 {
		cliques = append(cliques, append([]int(nil), r...))
	} else {
		for i := 0; i < len(p); i++ {
			v := p[i]
			neighbors := []int{}
			for j := 0; j < len(g.edges[v]); j++ {
				if g.edges[v][j] == 1 {
					neighbors = append(neighbors, j)
				}
			}

			rnew := append(r, v)
			pnew := intersect(p, neighbors)
			xnew := intersect(x, neighbors)

			cliques = BronKerbosch(g, rnew, pnew, xnew, cliques)

			p = remove(p, v)
			x = append(x, v)
		}
	}
	return cliques
}

func intersect(a, b []int) []int {
	var result []int
	for _, x := range a {
		for _, y := range b {
			if x == y {
				result = append(result, x)
				break
			}
		}
	}
	return result
}

func remove(a []int, v int) []int {
	result := make([]int, 0)
	for _, x := range a {
		if x != v {
			result = append(result, x)
		}
	}
	return result
}

// detectCodeDuplication uses Bron-Kerbosch to find code clones
func detectCodeDuplication(codeFragments []*CodeFragment) [][]int {
	// Create the graph
	graph := NewGraph()
	for _, cf := range codeFragments {
		graph.AddNode(cf)
	}

	// Add edges between nodes if they have similar content
	for i := 0; i < len(graph.nodes); i++ {
		for j := i + 1; j < len(graph.nodes); j++ {
			if strings.Contains(graph.nodes[i].content, graph.nodes[j].content) || strings.Contains(graph.nodes[j].content, graph.nodes[i].content) {
				graph.AddEdge(i, j)
			}
		}
	}

	p := createNodeRange(len(graph.nodes))

	// Find cliques using Bron-Kerbosch
	cliques := BronKerbosch(graph, []int{}, p, []int{}, [][]int{})

	return cliques
}

// createNodeRange generates a range of node indices [0, len-1]
func createNodeRange(n int) []int {
	nodes := make([]int, n)
	for i := 0; i < n; i++ {
		nodes[i] = i
	}
	return nodes
}

func main() {
	codeFragments := []*CodeFragment{
		{content: "func add(a, b int) int { return a + b }", lines: []int{1, 2, 3}},
		{content: "func subtract(a, b int) int { return a - b }", lines: []int{4, 5, 6}},
		{content: "func multiply(a, b int) int { return a * b }", lines: []int{7, 8, 9}},
		{content: "func addOne(num int) int { return add(num, 1) }", lines: []int{10, 11, 12}},
		{content: "func isEven(num int) bool { return num % 2 == 0 }", lines: []int{13, 14}},
		{content: "func double(num int) int { return multiply(num, 2) }", lines: []int{15, 16, 17}},
		{content: "func subtractFive(num int) int { return subtract(num, 5) }", lines: []int{18, 19, 20}},
	}

	cliques := detectCodeDuplication(codeFragments)

	fmt.Println("Code Duplication Cliques:")
	for _, clique := range cliques {
		fmt.Println("Clique:")
		for _, nodeIndex := range clique {
			fmt.Printf("%s (lines: %v)\n", codeFragments[nodeIndex].content, codeFragments[nodeIndex].lines)
		}
		fmt.Println("--------------------")
	}
}
