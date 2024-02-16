package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/common"
	"github.com/xbklyn/getgoal-app/model"
	"github.com/xbklyn/getgoal-app/service"
)

type ProgramController struct {
	service.ProgramService
}

func NewProgramController(programService service.ProgramService) *ProgramController {
	return &ProgramController{ProgramService: programService}
}

func (controller ProgramController) Route(api *gin.RouterGroup) {
	api.GET("/programs", controller.FindAllPrograms)
	api.GET("/programs/:id", controller.FindProgramByID)
	api.POST("/programs/search", controller.FindProgramByText)
	api.POST("/programs/filter", controller.FindProgramByLabel)
	api.POST("/programs", controller.SaveProgram)
}

// Find all programs  godoc
// @summary Find All program
// @description Find All program
// @tags Program
// @id FindAllProgram
// @produce json
// @response 200 {object} model.GeneralResponse "OK"
// @response 400 {object} model.GeneralResponse "Bad Request"
// @Router /api/v1/programs [get]
func (controller ProgramController) FindAllPrograms(c *gin.Context) {
	programs, err := controller.ProgramService.FindAllPrograms()
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Went Wrong",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	programsDTO := make([]model.ProgramDTO, 0)
	if len(programs) > 0 {
		programsDTO = model.ConvertToProgramDTOs(programs)
	}
	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Count:   len(programs),
		Data:    programsDTO,
		Error:   nil,
	})
}

// Find program by ID  godoc
// @summary Find Program by ID
// @description Find program by id
// @tags Program
// @id FindProgramById
// @param id path int true "Program ID"
// @produce json
// @response 200 {object} model.GeneralResponse "OK"
// @response 400 {object} model.GeneralResponse "Bad Request"
// @response 404 {object} model.GeneralResponse "Not Found"
// @Router /api/v1/programs/:id [get]
func (controller ProgramController) FindProgramByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	program, err := controller.ProgramService.FindProgramByID(id)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, model.GeneralResponse{
				Code:    http.StatusNotFound,
				Message: "Not Found",
				Data:    nil,
				Error:   err.Error(),
			})
			return
		}
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Went Wrong",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	programDTO := model.ConvertToProgramDTO(*program)
	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Count:   1,
		Data:    programDTO,
		Error:   nil,
	})
}

// Find program by Label  godoc
// @summary Filter Program
// @description Filter program by labels
// @tags Program
// @id FindProgramByLabel
// @param labels body []string true "label name"
// @produce json
// @response 200 {object} model.GeneralResponse "OK"
// @response 400 {object} model.GeneralResponse "Bad Request"
// @Router /api/v1/programs/filter [post]
func (controller ProgramController) FindProgramByLabel(c *gin.Context) {
	labels := new(model.Filter)
	if err := common.Bind(c, labels); err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Label 1",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	VaErr := common.Validate(labels)
	if VaErr != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Label",
			Data:    nil,
			Error:   VaErr.Error(),
		})
		return

	}

	programs, err := controller.ProgramService.FindProgramByLabel(labels.Labels)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Went Wrong",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	programsDTO := make([]model.ProgramDTO, 0)
	if len(programs) > 0 {
		programsDTO = model.ConvertToProgramDTOs(programs)
	}
	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Count:   len(programs),
		Data:    programsDTO,
		Error:   nil,
	})
}

// Find program by Text  godoc
// @summary Search Program
// @description Search program
// @tags Program
// @id FindProgramByText
// @param text body model.Search true "Search Text"
// @produce json
// @response 200 {object} model.GeneralResponse "OK"
// @response 400 {object} model.GeneralResponse "Bad Request"
// @response 404 {object} model.GeneralResponse "Not Found"
// @Router /api/v1/search [post]
func (controller ProgramController) FindProgramByText(c *gin.Context) {
	text := new(model.Search)
	if err := common.Bind(c, text); err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid texuigiuiiigut",
			Data:    nil,
			Error:   err.Error(),
		})
		return

	}
	VaErr := common.Validate(text)
	if VaErr != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid text",
			Data:    nil,
			Error:   VaErr.Error(),
		})
		return

	}
	programs, err := controller.ProgramService.FindProgramByText(text.SearchText)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Went Wrong",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	programsDTO := make([]model.ProgramDTO, 0)
	if len(programs) > 0 {
		programsDTO = model.ConvertToProgramDTOs(programs)
	}
	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Count:   len(programs),
		Data:    programsDTO,
		Error:   nil,
	})
}

// Save Program  godoc
// @summary Save new program
// @description Save new program
// @tags Program
// @id SaveProgram
// @param program body model.ProgramCreateOrUpdate true "Program Create or Update"
// @produce json
// @response 201 {object} model.GeneralResponse "Created"
// @response 400 {object} model.GeneralResponse "Bad Request"
// @Router /api/v1/programs [post]
func (controller ProgramController) SaveProgram(c *gin.Context) {
	program := new(model.ProgramCreateOrUpdate)
	if err := common.Bind(c, program); err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	programCreate, err := controller.ProgramService.Save(*program)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	programDTO := model.ConvertToProgramDTO(programCreate)
	c.JSON(http.StatusCreated, model.GeneralResponse{
		Code:    http.StatusCreated,
		Message: "Success",
		Count:   1,
		Data:    programDTO,
		Error:   nil,
	})
}
