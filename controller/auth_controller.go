package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/common"
	"github.com/xbklyn/getgoal-app/model"
	"github.com/xbklyn/getgoal-app/service"
)

type AuthController struct {
	service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return AuthController{authService}
}

func (controller *AuthController) Routes(api *gin.RouterGroup) {
	api.POST("/auth/register", controller.Register)
	api.POST("/auth/verify", controller.Verify)
}

func (controller *AuthController) Register(c *gin.Context) {

	var request model.SignUpRequest
	if err := common.Bind(c, &request); err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Request",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	user, err := controller.AuthService.SignUp(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad request",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Sign up Success",
		Data:    user,
		Error:   nil,
	})
}

func (controller *AuthController) Verify(c *gin.Context) {
	var request model.VerifyRequest
	if err := common.Bind(c, &request); err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Request",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	err := controller.AuthService.Verify(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad request",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Verify Success",
		Data:    nil,
		Error:   nil,
	})
}
