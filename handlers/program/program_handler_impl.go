package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/common"
	pserializers "github.com/xbklyn/getgoal-app/models/serializers/program"
	pvalidators "github.com/xbklyn/getgoal-app/models/validators/program"
	pusecases "github.com/xbklyn/getgoal-app/usecases/program"
)

type ProgramHandlerImpl struct {
	programUsecase pusecases.ProgramUsecase
}

// AddNewProgram implements ProgramHandler.
func (p *ProgramHandlerImpl) AddNewProgram(c *gin.Context) {
	programValidator := pvalidators.NewProgramValidator()
	if err := programValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	if err := p.programUsecase.Save(&programValidator.ProgramModel, programValidator.LabelName); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Database", err))
		return
	}

	serializer := pserializers.ProgramSerializer{C: c, Program: programValidator.ProgramModel}
	c.JSON(http.StatusOK, gin.H{"data": serializer.Response()})

}

// FilterProgram implements ProgramHandler.
func (p *ProgramHandlerImpl) FilterProgram(c *gin.Context) {
	filter := c.Query("filter")

	program, err := p.programUsecase.FilterProgram(filter)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("Program", err))
		return
	}

	serializer := pserializers.ProgramsSerializer{C: c, Program: program, Count: len(program)}
	c.JSON(http.StatusOK, gin.H{"data": serializer.Response()})
}

// FindAllPrograms implements ProgramHandler.
func (p *ProgramHandlerImpl) FindAllPrograms(c *gin.Context) {
	programs, err := p.programUsecase.FindAllPrograms()
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("err", err))
		return
	}

	serializer := pserializers.ProgramsSerializer{C: c, Program: programs, Count: len(programs)}
	c.JSON(http.StatusOK, gin.H{"data": serializer.Response()})
}

// FindProgramByID implements ProgramHandler.
func (p *ProgramHandlerImpl) FindProgramByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewError("Program", err))
		return
	}

	program, err := p.programUsecase.FindProgramByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("Program", err))
		return
	}

	serializer := pserializers.ProgramSerializer{C: c, Program: *program}
	c.JSON(http.StatusOK, gin.H{"Program": serializer.Response()})

}

// GetSearchProgram implements ProgramHandler.
func (p *ProgramHandlerImpl) GetSearchProgram(c *gin.Context) {
	text := c.Query("text")

	program, err := p.programUsecase.FindSearchProgram(text)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("Program", err))
		return
	}

	serializer := pserializers.ProgramsSerializer{C: c, Program: program, Count: len(program)}
	c.JSON(http.StatusOK, gin.H{"data": serializer.Response()})
}

func NewProgramHandlerImpl(programUsecase pusecases.ProgramUsecase) ProgramHandler {
	return &ProgramHandlerImpl{programUsecase}
}
