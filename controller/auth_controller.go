package controller

import (
	"log"
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

func (controller *AuthController) RouteAnonymous(api *gin.RouterGroup) {
	api.POST("/auth/register", controller.Register)
	api.POST("/auth/verify", controller.Verify)
	api.POST("/auth/sign-in", controller.SignIn)
}

func (controller *AuthController) Route(api *gin.RouterGroup) {
	api.POST("/auth/sign-out", controller.SignOut)
}

// Register user godoc
// @summary Register new user
// @description Register new user
// @tags Authentication
// @id Save user
// @accept json
// @produce json
// @param request body model.SignUpRequest true "Sign Up Request"
// @response 201 {object} model.GeneralResponse "Created"
// @response 400 {object} model.GeneralResponse "Bad Request"
// @Router /api/v1/auth/register [post]
func (controller *AuthController) Register(c *gin.Context) {

	var request model.SignUpRequest
	if err := common.Bind(c, &request); err != nil {
		log.Default().Printf("Error: %v", err)
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Request",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	_, err := controller.AuthService.SignUp(request)
	if err != nil {
		log.Default().Printf("Error: %v", err)
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad request",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, model.GeneralResponse{
		Code:    http.StatusCreated,
		Message: "Sign up Success",
		Data:    nil,
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
		log.Default().Printf("Error: %v", err)
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
		log.Default().Printf("Error: %v", err)
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
		log.Default().Printf("Error: %v", err)
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
		log.Default().Printf("Error: %v", err)
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

// Sign Out godoc
// @summary Sign Out
// @description Sign out and black list token
// @tags Authentication
// @id SignOut
// @accept json
// @produce json
// @response 200 {object} model.GeneralResponse "OK"
// @response 400 {object} model.GeneralResponse "Bad Request"
// @Router /api/v1/auth/sign-out [post]
func (controller *AuthController) SignOut(c *gin.Context) {
	err := controller.AuthService.SignOut()
	if err != nil {
		log.Default().Printf("Error: %v", err)
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
		Message: "Sign out Success",
		Data:    nil,
		Error:   nil,
	})
}
