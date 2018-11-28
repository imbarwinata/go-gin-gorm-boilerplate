package controllers

import (
	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func (h HealthController) Status(c *gin.Context) {
	c.JSON(400, gin.H{"message": "Check Ok"})
	c.Abort()
	return
}
