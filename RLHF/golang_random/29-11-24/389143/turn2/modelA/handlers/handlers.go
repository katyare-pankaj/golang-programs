package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HelloHandler is the struct that will define the request and response structure for the endpoint
type HelloHandler struct {
}

// NewHelloHandler creates a new instance of HelloHandler
func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

// HandleHello is the handler function for the GET /v1/hello endpoint
func (h *HelloHandler) HandleHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}
