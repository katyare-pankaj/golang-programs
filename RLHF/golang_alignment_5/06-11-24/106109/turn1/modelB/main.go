package main

import (
	"fmt"
	"log"
	"os"
)

type Drone struct {
	Steps []Step
}

type Step interface {
	Run() error
}

type GoBuildStep struct{}

func (s *GoBuildStep) Run() error {
	fmt.Println("Running Go build step...")
	// Run the Go build command
	return os.Run("go", "build", "-o", "app")
}

func main() {
	d := &Drone{}
	d.Steps = append(d.Steps, &GoBuildStep{})

	// Add custom build steps here
	// For example, a Docker build step
	// d.Steps = append(d.Steps, &DockerBuildStep{})

	if err := d.Run(); err != nil {
		log.Fatalf("Drone build failed: %v", err)
	}
}

func (d *Drone) Run() error {
	for _, step := range d.Steps {
		if err := step.Run(); err != nil {
			return err
		}
	}
	return nil
}
