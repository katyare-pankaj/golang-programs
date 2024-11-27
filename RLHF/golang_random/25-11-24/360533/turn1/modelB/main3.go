package main
import (
    "fmt"
    "github.com/go-perforce/p4"
)
func main() {
    // Connect to a Perforce server
    server, err := p4.Connect("p4.example.com:1666")
    if err != nil {
        fmt.Println("Error connecting to server:", err)
        return
    }
    // Authenticate with the server
    err = server.Login("username", "password")
    if err != nil {
        fmt.Println("Error authenticating:", err)
        return
    }
    // Create a new changelist
    changelist, err := server.NewChangelist()
    if err != nil {
        fmt.Println("Error creating changelist:", err)
        return
    }
    // Add files to the changelist
    err = changelist.Add("main.go", "README.md")
    if err != nil {
        fmt.Println("Error adding files to changelist:", err)
        return