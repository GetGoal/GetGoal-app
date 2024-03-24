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
	api.POST("")
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
