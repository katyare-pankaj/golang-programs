package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Azure/azure-functions-go/functions"
)

// DataTransformationInput is the input structure for the Azure Function.
type DataTransformationInput struct {
	Data string `json:"data"`
}

// DataTransformationOutput is the output structure for the Azure Function.
type DataTransformationOutput struct {
	TransformedData string `json:"transformedData"`
}

// DataTransformation is the Azure Function handler.
func DataTransformation(ctx context.Context, req *http.Request) (resp *http.Response, err error) {
	// Parse the request body into DataTransformationInput
	var input DataTransformationInput
	if err = json.NewDecoder(req.Body).Decode(&input); err != nil {
		return nil, functions.NewHttpError(http.StatusBadRequest, fmt.Sprintf("Error parsing request body: %v", err))
	}

	// Perform the data transformation (uppercasing each string)
	transformedData := strings.ToUpper(input.Data)

	// Create the DataTransformationOutput response
	output := DataTransformationOutput{
		TransformedData: transformedData,
	}

	// Marshal the response body
	respBody, err := json.Marshal(output)
	if err != nil {
		return nil, functions.NewHttpError(http.StatusInternalServerError, fmt.Sprintf("Error marshaling response body: %v", err))
	}

	// Return the response
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(strings.NewReader(string(respBody))),
	}, nil
}

func main() {
	port := os.Getenv("FUNCTIONS_WORKER_RUNTIME_PORT")
	if port == "" {
		log.Fatal("FUNCTIONS_WORKER_RUNTIME_PORT environment variable not set.")
	}

	// Start the Azure Function runtime
	log.Printf("Starting Azure Function on port %s", port)
	functions.Start(DataTransformation)
}
