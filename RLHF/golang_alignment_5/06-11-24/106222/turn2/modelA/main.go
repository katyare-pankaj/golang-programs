package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

// FunctionSignature represents the signature of a function
type FunctionSignature struct {
	name       string
	params     []string
	returnType string
}

// Graph represents the codebase as a graph where nodes are functions and edges represent similarity
type Graph struct {
	nodes []*FunctionSignature
	edges [][]int
}

// NewGraph creates a new graph
func NewGraph() *Graph {
	return &Graph{nodes: make([]*FunctionSignature, 0), edges: make([][]int, 0)}
}

// AddNode adds a node to the graph
func (g *Graph) AddNode(fs *FunctionSignature) {
	g.nodes = append(g.nodes, fs)
	g.edges = append(g.edges, make([]int, len(g.nodes)))
}

// AddEdge adds an edge between two nodes
func (g *Graph) AddEdge(i, j int) {
	g.edges[i][j] = 1
	g.edges[j][i] = 1 // Undirected graph
}

// JaccardSimilarity calculates the Jaccard similarity between two sets of strings
func JaccardSimilarity(set1, set2 []string) float64 {
	intersection := 0
	union := len(set1) + len(set2)

	for _, s1 := range set1 {
		for _, s2 := range set2 {
			if s1 == s2 {
				intersection++
				break
			}
		}
	}

	return float64(intersection) / float64(union)
}

// detectSimilarFunctions uses graph algorithms to find similar functions
func detectSimilarFunctions(filePath string, threshold float64) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	// Create the graph
	graph := NewGraph()

	// Extract function signatures from the AST
	ast.Inspect(file, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.FuncDecl:
			fs := &FunctionSignature{
				name:       n.Name.Name,
				params:     make([]string, 0),
				returnType: strings.TrimSpace(fmt.Sprint(n.Type.Results.List[0].Type)),
			}

			for _, param := range n.Type.Params.List {
				fs.params = append(fs.params, strings.TrimSpace(fmt.Sprint(param.Type)))
			}

			graph.AddNode(fs)
			return true
		}
		return true
	})

	// Add edges between nodes if they have similar signatures
	for i := 0; i < len(graph.nodes); i++ {
		for j := i + 1; j < len(graph.nodes); j++ {
			sig1 := graph.nodes[i]
			sig2 := graph.nodes[j]

			paramSet1 := make([]string, len(sig1.params))
			copy(paramSet1, sig1.params)
			paramSet2 := make([]string, len(sig2.params))
			copy(paramSet2, sig2.params)

			similarity := JaccardSimilarity(paramSet1, paramSet2)
			if similarity >= threshold {
				graph.AddEdge(i, j)
			}
		}
	}

	// Find connected components (similar function groups)
	components := findConnectedComponents(graph)

	fmt.Println("Similar Functions:")
	for i, component := range components {
		fmt.Printf("Group %d:\n", i+1)
		for _, nodeIndex := range component {
			fs := graph.nodes[nodeIndex]
			fmt.Printf("%s(%s) %s\n", fs.name, strings.Join(fs.params, ", "), fs.returnType)
		}
		fmt.Println("--------------------")
	}
}

// findConnectedComponents uses Depth-First Search to find connected components in the graph
func findConnectedComponents(graph *Graph) [][]int {
	var components [][]int
	visited := make([]bool, len(graph.nodes))

	for i := 0; i < len(graph.nodes); i++ {
		if !visited[i] {
			component := make([]int, 0)
			dfs(graph, i, &component, &visited)
			components = append(components, component)
		}
	}

	return components
}

func dfs(graph *Graph, node int, component *[]int, visited *[]bool) {
	*visited = append(*visited, true)
	*component = append(*component, node)

	for i := 0; i < len(graph.nodes); i++ {
		if graph.edges[node][i] == 1 && !(*visited)[i] {
			dfs(graph, i, component, visited)
		}
	}
}

func main() {
	const filePath = "your_codebase.go" // Replace this with the actual Go file path
	const similarityThreshold = 0.7     // Adjust this threshold as needed
	detectSimilarFunctions(filePath, similarityThreshold)
}
