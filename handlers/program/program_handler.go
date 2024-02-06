package handlers

import "github.com/gin-gonic/gin"

type ProgramHandler interface {
	GetSearchProgram(c *gin.Context)
	FindAllPrograms(c *gin.Context)
	FindProgramByID(c *gin.Context)
	AddNewProgram(c *gin.Context)
	FilterProgram(c *gin.Context)
}
