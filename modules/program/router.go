package program

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/common"
)

func ProgramRegister(router *gin.RouterGroup) {
	router.POST("", ProgramCreate)
}
func ProgramAnonymousRegister(router *gin.RouterGroup) {
	router.GET("", ProgramList)
	router.GET("/:id", ProgramDetail)
}

func ProgramCreate(c *gin.Context) {
	programValidator := NewProgramValidator()
	if err := programValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	if err := SaveOne(&programValidator.programModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Database", err))
		return
	}

	serializer := ProgramSerializer{c, programValidator.programModel}
	c.JSON(http.StatusOK, gin.H{"Program": serializer.Response()})

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
