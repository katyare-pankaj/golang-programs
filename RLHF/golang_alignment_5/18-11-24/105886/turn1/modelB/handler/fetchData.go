package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUsers(c *gin.Context) {
	users, err := dbOperation.GetUsers()
	if err != nil {
		c.String(http.StatusInternalServerError, "Error fetching users: %v", err)
		return
	}
	c.JSON(http.StatusOK, users)
}
