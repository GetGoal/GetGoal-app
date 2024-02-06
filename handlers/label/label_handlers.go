package handlers

import (
	"github.com/gin-gonic/gin"
)

type LabelHandler interface {
	GetSeachLabel(c *gin.Context)
	FindAllLabels(c *gin.Context)
	FindLabelByID(c *gin.Context)
	AddNewLabel(c *gin.Context)
}
