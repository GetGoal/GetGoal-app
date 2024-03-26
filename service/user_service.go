package service

import (
	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/entity"
	"github.com/xbklyn/getgoal-app/model"
)

type UserService interface {
	FindUserByEmail(c *gin.Context) (*entity.UserAccount, error)
	ResetPassword(c *gin.Context, credentail model.ChangePasswordRequest) (*entity.UserAccount, error)
}
