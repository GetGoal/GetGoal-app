package program

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/common"
)

func ProgramAnonymousRegister(router *gin.RouterGroup) {
	router.GET("", ProgramList)
	router.GET("/:id", ProgramDetail)
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

func ProgramDetail(c *gin.Context) {
	programId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewError("Program", err))
		return
	}

	program, err := FindOneProgram(&Program{ProgramID: programId})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("Program", err))
		return
	}

	serializer := ProgramSerializer{C: c, Program: program}
	c.JSON(http.StatusOK, gin.H{"Program": serializer.Response()})

}
