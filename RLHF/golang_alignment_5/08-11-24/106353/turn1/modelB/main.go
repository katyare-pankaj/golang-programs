package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	router := gin.Default()

	// Initialize MySQL database connection
	db, err := sqlx.Open("mysql", "root:your-password@tcp(127.0.0.1:3306)/your-database?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Mount routes
	router.GET("/users/:id", func(c *gin.Context) {
		userID := c.Param("id")

		// Validate userID input as integer
		userIDInt, err := strconv.Atoi(userID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		// Data validation passed, proceed with database query
		var user struct {
			ID    int    `db:"id"`
			Name  string `db:"name"`
			Email string `db:"email"`
			Age   int    `db:"age"`
		}

		err = db.Get(&user, "SELECT * FROM users WHERE id = ?", userIDInt)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user data"})
			return
		}

		c.JSON(http.StatusOK, user)
	})

	// Run the server
	router.Run(":8080")
}
