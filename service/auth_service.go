package service

import (
	"github.com/xbklyn/getgoal-app/entity"
	"github.com/xbklyn/getgoal-app/model"
)

type AuthService interface {
	SignUp(request model.SignUpRequest) (useEntityr entity.UserAccount, err error)
	SignIn(request model.Credentials) (accessToken string, refreshToken string, err error)
	SignOut() error
	Verify(request model.VerifyRequest) error
	IsTokenBlacklisted(tokenString string) bool
}
