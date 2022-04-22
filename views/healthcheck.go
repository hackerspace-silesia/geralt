package views

import "github.com/gin-gonic/gin"

func HealtcheckHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ping... pong",
	})
}
