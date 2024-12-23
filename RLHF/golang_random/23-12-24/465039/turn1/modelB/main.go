package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numTenants           = 10
	maxRequestsPerTenant = 5
)

type TenantRequest struct {
	tenantID  int
	requestID int
}

func worker(wg *sync.WaitGroup, requests <-chan TenantRequest) {
	defer wg.Done()
	for request := range requests {
		fmt.Printf("Processing tenant %d, request %d\n", request.tenantID, request.requestID)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func main() {
	wg := &sync.WaitGroup{}
	requestChan := make(chan TenantRequest)

	// Start workers
	for i := 0; i < numTenants; i++ {
		wg.Add(1)
		go worker(wg, requestChan)
	}

	// Generate requests for each tenant
	for tenantID := 1; tenantID <= numTenants; tenantID++ {
		for requestID := 1; requestID <= maxRequestsPerTenant; requestID++ {
			request := TenantRequest{tenantID, requestID}
			requestChan <- request
		}
	}
	close(requestChan)

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("All requests completed.")
}
