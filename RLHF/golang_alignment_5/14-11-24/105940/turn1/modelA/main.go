package main

import (
	"github.com/gin-gonic/gin"
)

// Employee model
type Employee struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

var employees []Employee

func main() {
	router := gin.Default()

	// Sample data for demonstration
	employees = []Employee{
		{ID: 1, FirstName: "Alice", LastName: "Anderson", Email: "alice@example.com"},
		{ID: 2, FirstName: "Bob", LastName: "Brown", Email: "bob@example.com"},
	}

	// API Endpoints
	router.GET("/employees", getEmployees)
	router.GET("/employees/:id", getEmployee)
	router.POST("/employees", postEmployee)

	router.Run(":8080")
}

func getEmployees(c *gin.Context) {
	c.JSON(200, employees)
}

func getEmployee(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.String(400, "ID parameter is required")
		return
	}

	for _, emp := range employees {
		if emp.ID == int(id) {
			c.JSON(200, emp)
			return
		}
	}
	c.String(404, "Employee not found")
}

func postEmployee(c *gin.Context) {
	var newEmployee Employee
	if err := c.BindJSON(&newEmployee); err != nil {
		c.String(400, "Invalid employee data")
		return
	}

	newEmployee.ID = len(employees) + 1
	employees = append(employees, newEmployee)
	c.JSON(201, newEmployee)
}
