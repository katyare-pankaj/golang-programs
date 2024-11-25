package main

import (
	"fmt"
	"go-programs/RLHF/golang_random/25-11-24/360575/turn2/modelA/envconfig"
)

func main() {
	cfg, err := envconfig.NewConfig()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Database URL: %q\n", cfg.DatabaseURL)
	fmt.Printf("Port: %d\n", cfg.Port)
	fmt.Printf("Debug: %v\n", cfg.Debug)

	// Use the configuration values as needed for your application
}
