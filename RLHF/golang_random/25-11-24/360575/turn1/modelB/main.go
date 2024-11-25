package main

import (
	"fmt"
	"go-programs/RLHF/golang_random/25-11-24/360575/turn1/modelB/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println("Database URL:", cfg.DatabaseURL)
	// Use the database URL as needed for your application
}
