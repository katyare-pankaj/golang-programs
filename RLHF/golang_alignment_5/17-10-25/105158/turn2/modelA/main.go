package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/dataflowx/dataflowx-go"
)

func main() {
	// Initialize DataFlowX client
	dfx, err := dataflowx.NewClient("YOUR_DATAFLOWX_API_KEY")
	if err != nil {
		log.Fatalf("Error initializing DataFlowX client: %v", err)
	}

	// Define the data transformation pipeline using DataFlowX visual workflow builder
	pipelineID, err := dfx.CreatePipeline("MyDataTransformationPipeline")
	if err != nil {
		log.Fatalf("Error creating pipeline: %v", err)
	}

	err = dfx.AddReadCSVStep(pipelineID, "input.csv")
	if err != nil {
		log.Fatalf("Error adding Read CSV step: %v", err)
	}

	err = dfx.AddTransformStep(pipelineID, "uppercase", func(data interface{}) interface{} {
		str := data.(string)
		return strings.ToUpper(str)
	})
	if err != nil {
		log.Fatalf("Error adding Transform step: %v", err)
	}

	err = dfx.AddWriteCSVStep(pipelineID, "output.csv")
	if err != nil {
		log.Fatalf("Error adding Write CSV step: %v", err)
	}

	// Execute the data transformation pipeline
	err = dfx.RunPipeline(pipelineID)
	if err != nil {
		log.Fatalf("Error running pipeline: %v", err)
	}

	fmt.Println("Data transformation pipeline executed successfully!")
}
