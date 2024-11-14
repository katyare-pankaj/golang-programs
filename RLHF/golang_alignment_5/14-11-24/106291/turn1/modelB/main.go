package main

import (
    "fmt"
    "go/ast"
    "go/parser"
    "go/token"
    "io/ioutil"
)

// measureCC calculates McCabe's cyclomatic complexity of a Go function.
func measureCC(functionNode *ast.FuncDecl) int {
    var cc int

    // Count nodes: and, or, if, for, select, etc.
    // We use an ast.Visitor to traverse the AST
    type complexityVisitor struct {
        count int
    }
    v := &complexityVisitor{}

    ast.Walk(v, functionNode)

    // Visiting a statement node increments the counter
    func (v *complexityVisitor) Visit(node ast.Node) ast.Visitor {
        switch stmt := node.(type) {
        case *ast.IfStmt, *ast.ForStmt, *ast.SwitchStmt, *ast.SelectStmt,
            *ast.RangeStmt, *ast.BlockStmt:
            v.count++
        default:
        }
        return v
    }

    edges := 0
    // For simple functions, we assume only edge is the return. For more
    // involved calculations consider multiple exits, etc.
    edges = 1
    // Add 1 for the entry point
    edges++
    cc = edges - v.count + 2
    return cc
}

func main() {
    // Parse the Go file
    fset := token.NewFileSet()
    file, err := parser.ParseFile(fset, "main.go", ioutil.ReadFile("main.go"), 0)
    if err != nil {
        panic(err)
    }

    // Get the first function in the file
    for _, decl := range file.Decls {
        if funcDecl, ok := decl.(*ast.FuncDecl); ok && funcDecl.Name.Name == "main" {
        	// Measure CC of the main function
            cc := measureCC(funcDecl)
            fmt.Printf("Cyclomatic Complexity of 'main' function: %d\n", cc)
            return
        }
    }
    fmt.Println("Function 'main' not found.")
}