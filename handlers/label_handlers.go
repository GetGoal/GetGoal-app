package handlers

import (
	"github.com/gin-gonic/gin"
)

type LabelHandler interface {
	FindAllLabels(c *gin.Context)
	FindLabelByID(c *gin.Context)
}
