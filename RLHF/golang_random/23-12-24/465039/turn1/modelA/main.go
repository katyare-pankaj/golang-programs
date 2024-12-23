package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Tenant represents a tenant in the application
type Tenant struct {
	ID   string
	Name string
}

// simulateTenantWork simulates work for a tenant with a delay
func simulateTenantWork(wg *sync.WaitGroup, tenant *Tenant) {
	defer wg.Done()

	fmt.Printf("Starting work for tenant: %s\n", tenant.Name)
	time.Sleep(time.Duration(runtime.GOMAXPROCS(0)) * time.Millisecond)
	fmt.Printf("Completed work for tenant: %s\n", tenant.Name)
}

func main() {
	// List of tenants
	tenants := []*Tenant{
		&Tenant{ID: "t1", Name: "Tenant One"},
		&Tenant{ID: "t2", Name: "Tenant Two"},
		&Tenant{ID: "t3", Name: "Tenant Three"},
		&Tenant{ID: "t4", Name: "Tenant Four"},
		&Tenant{ID: "t5", Name: "Tenant Five"},
	}

	// Create a WaitGroup to wait for all tenant work to complete
	var wg sync.WaitGroup

	// Start the work for each tenant
	for _, tenant := range tenants {
		wg.Add(1)
		go simulateTenantWork(&wg, tenant)
	}

	// Wait for all tenants to complete their work
	wg.Wait()

	// All tenant work is complete
	fmt.Println("All tenant work completed.")
}
