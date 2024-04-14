package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/model"
	"github.com/xbklyn/getgoal-app/service"
)

type UserController struct {
	service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return UserController{userService}
}
func (controller UserController) RouteAnonymous(api *gin.RouterGroup) {
	api.POST("/users/reset-password", controller.ResetPassword)
}

func (controller UserController) Route(api *gin.RouterGroup) {
	api.GET("/users/profile", controller.FindUserByEmail)
	api.GET("/users/programs/saved", controller.FindSavedProgramByUser)
	api.GET("/users/programs/joined", controller.FindSavedProgramByUser)
	api.PUT("/users/labels", controller.UpdateUserLabel)
	api.GET("/users/calendar", controller.FindDateWithTask)
}
func (controller UserController) FindUserByEmail(c *gin.Context) {
	user, err := controller.UserService.FindUserByEmail(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Went Wrong",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	userDto := model.ConvertToUserDTO(*user)
	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    userDto,
		Error:   nil,
	})
}

func (controller UserController) ResetPassword(c *gin.Context) {
	var credential model.ChangePasswordRequest
	if err := c.ShouldBindJSON(&credential); err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Request",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	user, err := controller.UserService.ResetPassword(c, credential)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Went Wrong",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	userDto := model.ConvertToUserDTO(*user)
	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    userDto,
		Error:   nil,
	})
}

func (controller UserController) FindSavedProgramByUser(c *gin.Context) {
	savedPrograms, err := controller.UserService.FindSavedProgram(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Went Wrong",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	programDto := model.ConvertToProgramDTOs(savedPrograms)
	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    programDto,
		Error:   nil,
	})
}

func (controller UserController) FindJoinedProgramByUser(c *gin.Context) {
	joinedPrograms, err := controller.UserService.FindJoinedProgram(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Went Wrong",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	programDto := model.ConvertToProgramDTOs(joinedPrograms)
	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    programDto,
		Error:   nil,
	})
}

func (controller UserController) UpdateUserLabel(c *gin.Context) {
	var user model.UserModel
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Request",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	updated, err := controller.UserService.UpdateLabel(c, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Went Wrong",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	dto := model.ConvertToUserDTO(*updated)
	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    dto,
		Error:   nil,
	})
}

func (controller UserController) FindDateWithTask(c *gin.Context) {

	//check if date is valid
	dateStr := c.Query("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Request",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	datesWithTask, err := controller.UserService.FindDateWithTasks(c, date)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Went Wrong",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    datesWithTask,
		Error:   nil,
	})
}
