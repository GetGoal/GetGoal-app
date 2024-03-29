package service

import (
	"github.com/xbklyn/getgoal-app/entity"
	"github.com/xbklyn/getgoal-app/model"
)

type AuthService interface {
	SignUp(request model.SignUpRequest) (useEntityr entity.UserAccount, err error)
	SignIn(request model.Credentials) (accessToken string, refreshToken string, err error)
	ExternalSignIn(request model.ProviderSignInRequest) (accessToken string, refreshToken string, err error)
	SignOut(token string) error
	Verify(request model.VerifyRequest) error
	VerifyPasswordReset(request model.VerifyResetRequest) error
	IsTokenBlacklisted(tokenString string) bool
	ResetPassword(request model.ResetPasswordRequest) error
}
