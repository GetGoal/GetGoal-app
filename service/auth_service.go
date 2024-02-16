package service

import (
	"github.com/xbklyn/getgoal-app/entity"
	"github.com/xbklyn/getgoal-app/model"
)

type AuthService interface {
	SignUp(request model.SignUpRequest) (useEntityr entity.UserAccount, err error)
	Verify(request model.VerifyRequest) error
}
