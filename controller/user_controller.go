package controller

import (
	"net/http"

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

func (controller UserController) Route(api *gin.RouterGroup) {
	api.GET("/users/profile", controller.FindUserByEmail)
	api.GET("/users/programs/saved", controller.FindSavedProgramByUser)
	api.POST("/users/reset-password", controller.ResetPassword)
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
