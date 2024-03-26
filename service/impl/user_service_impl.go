package impl

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/common"
	"github.com/xbklyn/getgoal-app/entity"
	"github.com/xbklyn/getgoal-app/model"
	"github.com/xbklyn/getgoal-app/repository"
	"github.com/xbklyn/getgoal-app/service"
)

func NewUserServiceImpl(userRepo repository.UserRepo) service.UserService {
	return UserServiceImpl{userRepo}
}

type UserServiceImpl struct {
	UserRepo repository.UserRepo
}

// UpdateUser implements service.UserService.
func (service UserServiceImpl) ResetPassword(c *gin.Context, credential model.ChangePasswordRequest) (*entity.UserAccount, error) {
	claims := c.MustGet("claims").(*common.Claims)
	user, _ := service.UserRepo.FindUserByEmail(claims.Email)
	if user.UserID == 0 {
		return nil, errors.New("user not found")
	}
	if claims.Email != user.Email {
		return nil, errors.New("unauthorized")
	}

	isMatchedOldPassword, err := common.VerifyPassword(credential.NewPassword, user.PasswordSalt)
	if err != nil {
		return nil, err
	}
	if isMatchedOldPassword {
		return nil, errors.New("new password cannot be the same as the old password")
	}

	hashed, encodedHash, err := common.GenerateHashFromPassword(credential.NewPassword)
	if err != nil {
		return nil, err
	}

	user.PasswordHash = hashed
	user.PasswordSalt = encodedHash

	err = service.UserRepo.Update(user.UserID, user)
	if err != nil {
		return nil, err
	}
	return &user, nil
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
