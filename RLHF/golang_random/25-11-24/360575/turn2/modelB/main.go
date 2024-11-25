package main

import (
	"fmt"
	"go-programs/RLHF/golang_random/25-11-24/360575/turn2/modelB/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println("Database URL:", cfg.DatabaseURL)
	fmt.Println("Port:", cfg.Port)
	fmt.Println("Enable Logging:", cfg.EnableLogging)
	fmt.Println("Log Level:", cfg.LogLevel)
}
