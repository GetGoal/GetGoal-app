package service

import (
	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/entity"
)

type UserService interface {
	FindUserByEmail(c *gin.Context) (*entity.UserAccount, error)
}
