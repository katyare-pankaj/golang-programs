package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

const (
	networkTimeout           = time.Second * 5
	maxConcurrentConnections = 100
)

// PortScanner scans a single port for openness
func portScanner(wg *sync.WaitGroup, target string, port int) {
	defer wg.Done()

	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", target, port), networkTimeout)
	if err != nil {
		// Port is closed
		return
	}
	conn.Close()

	fmt.Printf("Port %d is open\n", port)
}

func main() {
	target := "example.com"
	ports := []int{22, 80, 443, 21, 25, 3306}

	var wg sync.WaitGroup
	sem := make(chan struct{}, maxConcurrentConnections)

	for _, port := range ports {
		wg.Add(1)
		sem <- struct{}{}
		go func(p int) {
			defer func() { <-sem }()
			portScanner(&wg, target, p)
		}(port)
	}

	wg.Wait()
	fmt.Println("Scan completed")
}
