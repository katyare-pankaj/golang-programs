package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	for {
		data := rand.Intn(100)
		fmt.Println("Sensor Data:", data)
		fmt.Fprintf(conn, "%d\n", data)
		time.Sleep(time.Second * 1)
	}
}
