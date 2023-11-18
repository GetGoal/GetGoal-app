package program

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/common"
)

func ProgramAnonymousRegister(router *gin.RouterGroup) {
	router.GET("", ProgramList)
}

func ProgramList(c *gin.Context) {
	programs, err := FindAllProgram()
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("Program", err))
		return
	}

	serializer := ProgramsSerializer{C: c, Program: programs, Count: len(programs)}
	c.JSON(http.StatusOK, gin.H{"Programs": serializer.Response()})
}