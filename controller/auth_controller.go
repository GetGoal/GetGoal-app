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
	api.POST("/auth/external-sign-in", controller.ProviderSignIn)
	api.POST("/auth/reset-password", controller.ResetPassword)
	api.POST("/auth/verify-password-reset", controller.VerifyReset)
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

	token := c.MustGet("access_token").(string)
	err := controller.AuthService.SignOut(token)
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

// Provider Sign In godoc
// @summary Provider Sign In
// @description Sign in outsider provider
// @tags Authentication
// @id ProviderSignIn
// @accept json
// @produce json
// @response 200 {object} model.GeneralResponse "OK"
// @response 400 {object} model.GeneralResponse "Bad Request"
// @Router /api/v1/auth/external-sign-in [post]
func (controller *AuthController) ProviderSignIn(c *gin.Context) {
	var request model.ProviderSignInRequest
	if err := common.Bind(c, &request); err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Request to sign in",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	access, refresh, err := controller.AuthService.ExternalSignIn(request)
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

// Reset Password godoc
// @summary Reset Password
// @description Reset Password
// @tags Authentication
// @id ProviderSignIn
// @accept json
// @produce json
// @response 200 {object} model.GeneralResponse "OK"
// @response 400 {object} model.GeneralResponse "Bad Request"
// @Router /api/v1/auth/external-sign-in [post]
func (controller *AuthController) ResetPassword(c *gin.Context) {
	var request model.ResetPasswordRequest
	if err := common.Bind(c, &request); err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Request",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	err := controller.AuthService.ResetPassword(request)
	if err != nil {
		if err.Error() == "email not found" {
			c.JSON(http.StatusNotFound, model.GeneralResponse{
				Code:    http.StatusNotFound,
				Message: "Email not found",
				Data:    nil,
				Error:   nil,
			})
			return
		}
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
		Message: "If email exist, reset password link has been sent to your email.",
		Data:    nil,
		Error:   nil,
	})

}

func (controller *AuthController) VerifyToken(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad request",
			Data:    nil,
			Error:   "Token required",
		})
		return
	}

	_, err := common.ValidateAccessToken(token)
	if err != nil {
		if err.Error() == "invalid token" {
			c.JSON(http.StatusUnauthorized, model.GeneralResponse{
				Code:    http.StatusUnauthorized,
				Message: "Invalid token",
				Data:    nil,
				Error:   nil,
			})
			return
		}
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
		Message: "Token is valid",
		Data:    nil,
		Error:   nil,
	})

}

func (controller *AuthController) VerifyReset(c *gin.Context) {
	var request model.VerifyResetRequest
	if err := common.Bind(c, &request); err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Request",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	err := controller.AuthService.VerifyPasswordReset(request)
	if err != nil {
		if err.Error() == "email not found" {
			c.JSON(http.StatusUnauthorized, model.GeneralResponse{
				Code:    http.StatusUnauthorized,
				Message: "Email not found",
				Data:    nil,
				Error:   nil,
			})
			return
		}
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
		Message: "code is valid",
		Data:    nil,
		Error:   nil,
	})

}
