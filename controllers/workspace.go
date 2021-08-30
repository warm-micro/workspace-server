package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListWorkspace(c *gin.Context) {
	userId, _ := c.GetQuery("userId")

	c.JSON(http.StatusOK, gin.H{
		"message": "test",
		"body":    userId,
	})
}

func CreateWorkspace(c *gin.Context) {

}
