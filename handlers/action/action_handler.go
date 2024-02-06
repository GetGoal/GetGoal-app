package handlers

import "github.com/gin-gonic/gin"

type ActionHandler interface {
	FindAllActions(c *gin.Context)
	FindActionByID(c *gin.Context)
}
