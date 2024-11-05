package main  
import (  
    "encoding/json"  
    "fmt"  
    "log"
)  

type Node struct {  
    Value int    `json:"value"`  
    Next  *Node `json:"next"`  
}  

// Custom Marshaler to break the circular reference
func (n *Node) MarshalJSON() ([]byte, error) {  
    if n == nil {
        return json.Marshal(nil)
    }
    type SimpleNode struct {
        Value int `json:"value"`
        Next  *SimpleNodeID `json:"next"`
    }
    type SimpleNodeID struct {
        ID int `json:"id"`
    }

    var out SimpleNode
    out.Value = n.Value

    // Create a mapping to ensure next nodes reference each other by ID
    nodeMap := make(map[*Node]int)
    nodeMap[n] = len(nodeMap)

    func marshalNext(node *Node) (interface{}, error) {
        if node == nil {
            return nil, nil
        }
        id, ok := nodeMap[node]
        if !ok {
            id = len(nodeMap)
            nodeMap[node] = id
        }
        return &SimpleNodeID{ID: id}, nil
    }

    out.Next, err := marshalNext(n.Next)
    if err != nil {
        return nil, err
    }

    return json.Marshal(out)
}

func main() {  
    // Example data with a circular reference  
    a := &Node{Value: 1}  
    b := &Node{Value: 2}  
    c := &Node{Value: 3}  

    a.Next = b  
    b.Next = c  
    c.Next = a // Circular reference back to A

    data, err := json.MarshalIndent(a, "", "  ")  
    if err != nil {  
        log.Fatal("Error marshaling data:", err)  
    }  
    fmt.Println(string(data))
}