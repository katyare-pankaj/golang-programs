package main

import (
	"flag"
	"fmt"
)

var (
	// General flags available to all users
	host    string
	port    int
	verbose bool

	// Admin-specific flags
	adminMode bool
	debug     bool
)

func main() {
	// Register flags
	flag.StringVar(&host, "host", "localhost", "Host to connect to")
	flag.IntVar(&port, "port", 8080, "Port to listen on")
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose mode")
	flag.BoolVar(&adminMode, "admin", false, "Enable admin mode")
	flag.BoolVar(&debug, "debug", false, "Enable debug mode")

	// Parse command-line flags
	flag.Parse()

	// Determine user role
	userRole := "admin" // Replace this with actual role determination logic

	// Load and validate flags based on user role
	if err := loadFlagsByRole(userRole); err != nil {
		fmt.Println("Error loading flags:", err)
		return
	}

	// Display parsed flags
	fmt.Println("Host:", host)
	fmt.Println("Port:", port)
	fmt.Println("Verbose:", verbose)
	fmt.Println("Admin Mode:", adminMode)
	fmt.Println("Debug Mode:", debug)
}

func loadFlagsByRole(role string) error {
	// Check role and enable/disable flags accordingly
	switch role {
	case "admin":
		if !adminMode {
			flag.Set("admin", "true") // Enable admin mode if not set
		}
	case "user":
		// Disable admin-specific flags
		if adminMode {
			flag.Set("admin", "false")
		}
		if debug {
			flag.Set("debug", "false")
		}
	default:
		return fmt.Errorf("unknown role: %s", role)
	}

	// Validate flags
	if err := validateFlags(); err != nil {
		return err
	}

	return nil
}

func validateFlags() error {
	// Validate that port is within a valid range
	if port < 0 || port > 65535 {
		return fmt.Errorf("port must be between 0 and 65535, got %d", port)
	}

	// Additional validation logic can be added here
	return nil
}