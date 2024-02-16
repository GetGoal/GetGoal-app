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
	api.POST("/auth/sign-in", controller.SignIn)
}

// Register user godoc
// @summary Register new user
// @description Register new user
// @tags Authentication
// @id Save user
// @accept json
// @produce json
// @param request body model.SignUpRequest true "Sign Up Request"
// @response 200 {object} model.GeneralResponse "OK"
// @response 400 {object} model.GeneralResponse "Bad Request"
// @Router /api/v1/auth/register [post]
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

// Verifyuser godoc
// @summary Verify new user
// @description check verficatoin code from email
// @tags Authentication
// @id Verifyuser
// @accept json
// @produce json
// @param request body model.VerifyRequest true "Sign Up Request"
// @response 200 {object} model.GeneralResponse "OK"
// @response 400 {object} model.GeneralResponse "Bad Request"
// @Router /api/v1/auth/verify [post]
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

// Sign In godoc
// @summary Sign In
// @description Sign in with email and password
// @tags Authentication
// @id SignIn
// @accept json
// @produce json
// @param request body model.Credentials true "Sign Up Request"
// @response 200 {object} model.GeneralResponse "OK"
// @response 400 {object} model.GeneralResponse "Bad Request"
// @Router /api/v1/auth/sign-in [post]
func (controller *AuthController) SignIn(c *gin.Context) {
	var request model.Credentials
	if err := common.Bind(c, &request); err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Request",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	access, refresh, err := controller.AuthService.SignIn(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad request",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	tokens := model.TokenResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	}
	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Sign in Success",
		Data:    tokens,
		Error:   nil,
	})
}