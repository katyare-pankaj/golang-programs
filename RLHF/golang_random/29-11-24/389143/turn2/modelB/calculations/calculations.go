package calculations

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CalculateHandler handles calculate API requests
func CalculateHandler(c *gin.Context) {
	op := c.Param("operator")
	aStr := c.Param("a")
	bStr := c.Param("b")

	a, err := strconv.ParseFloat(aStr, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid argument a"})
		return
	}

	b, err := strconv.ParseFloat(bStr, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid argument b"})
		return
	}

	var result float64
	var message string
	switch op {
	case "+":
		result = a + b
		message = fmt.Sprintf("%.2f + %.2f = %.2f", a, b, result)
	case "-":
		result = a - b
		message = fmt.Sprintf("%.2f - %.2f = %.2f", a, b, result)
	case "*":
		result = a * b
		message = fmt.Sprintf("%.2f * %.2f = %.2f", a, b, result)
	case "/":
		if b == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Cannot divide by zero"})
			return
		}
		result = a / b
		message = fmt.Sprintf("%.2f / %.2f = %.2f", a, b, result)
	default:
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid operator"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": message})
}
