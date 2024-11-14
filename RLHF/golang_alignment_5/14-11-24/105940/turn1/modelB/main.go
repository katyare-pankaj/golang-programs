package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Employee struct from the domain model
type Employee struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func handleGetEmployee(ctx events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	employee := Employee{
		ID:    "1",
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}

	jsonEmp, err := json.Marshal(employee)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       fmt.Sprintf("Error: %v", err),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(jsonEmp),
	}, nil
}

func main() {
	lambda.Start(handleGetEmployee)
}
