package impl

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/common"
	"github.com/xbklyn/getgoal-app/entity"
	"github.com/xbklyn/getgoal-app/repository"
	"github.com/xbklyn/getgoal-app/service"
)

func NewUserServiceImpl(userRepo repository.UserRepo) service.UserService {
	return UserServiceImpl{userRepo}
}

type UserServiceImpl struct {
	UserRepo repository.UserRepo
}

// FindUserByEmail implements service.UserService.
func (service UserServiceImpl) FindUserByEmail(c *gin.Context) (*entity.UserAccount, error) {
	claims := c.MustGet("claims").(*common.Claims)
	log.Default().Println(claims.Email)
	user, err := service.UserRepo.FindUserByEmail(claims.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
