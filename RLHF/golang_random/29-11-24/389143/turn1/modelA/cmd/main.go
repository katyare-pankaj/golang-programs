package main

import (
	"fmt"
	"go-programs/RLHF/golang_random/29-11-24/389143/turn1/modelA/pkg/server"
)

func main() {
	fmt.Println("Starting server...")
	server.StartServer(":8080")
}
