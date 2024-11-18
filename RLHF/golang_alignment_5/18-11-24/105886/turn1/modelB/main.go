package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	r := gin.Default()

	// Template file path
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", func(c *gin.Context) {
		names := []string{"John", "Doe", "Jane"}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name": names,
		})
	})

	r.Run(":8080") // Listen and serve on port 8080
}
