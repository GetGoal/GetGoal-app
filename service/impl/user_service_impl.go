package impl

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/common"
	"github.com/xbklyn/getgoal-app/entity"
	"github.com/xbklyn/getgoal-app/model"
	"github.com/xbklyn/getgoal-app/repository"
	"github.com/xbklyn/getgoal-app/service"
	"github.com/zhenghaoz/gorse/client"
)

func NewUserServiceImpl(userRepo repository.UserRepo, programRepo repository.ProgramRepo, gorse client.GorseClient) service.UserService {
	return UserServiceImpl{userRepo, programRepo, gorse}
}

type UserServiceImpl struct {
	UserRepo    repository.UserRepo
	ProgramRepo repository.ProgramRepo
	GorseClient client.GorseClient
}

// FindDateWithTasks implements service.UserService.
func (service UserServiceImpl) FindDateWithTasks(c *gin.Context, date time.Time) ([]int, error) {
	claims := c.MustGet("claims").(*common.Claims)
	user, _ := service.UserRepo.FindUserByID(claims.UserID)
	if user.UserID == 0 {
		return nil, errors.New("user not found")
	}

	dates, err := service.UserRepo.FindDateWithTasks(date, user.UserID)
	if err != nil {
		return nil, err
	}
	log.Default().Printf("dates: %v", dates)
	datesWithTask := make([]int, 0)
	for _, date := range dates {
		parsedDate, _ := time.Parse("2006-01-02", date.Date)
		if date.Count > 0 {
			datesWithTask = append(datesWithTask, parsedDate.Day())
		}
	}
	return datesWithTask, nil

}

// UpdateLabel implements service.UserService.
func (service UserServiceImpl) UpdateLabel(c *gin.Context, userModel model.UserModel) (*entity.UserAccount, error) {
	claims := c.MustGet("claims").(*common.Claims)
	user, _ := service.UserRepo.FindUserByID(claims.UserID)
	if user.UserID == 0 {
		return nil, errors.New("user not found")
	}
	labelText, _ := json.Marshal(userModel.Labels)
	user.Labels = string(labelText)

	err := service.UserRepo.Update(user.UserID, user)
	if err != nil {
		return nil, err
	}
	userId := strconv.Itoa(int(user.UserID))
	_, gEer := service.GorseClient.UpdateUser(context.TODO(), userId, client.UserPatch{
		Labels: userModel.Labels,
	})
	if gEer != nil {
		return nil, gEer
	}
	return &user, nil
}

// FindJoinedProgram implements service.UserService.
func (service UserServiceImpl) FindJoinedProgram(c *gin.Context) ([]entity.Program, error) {
	claims := c.MustGet("claims").(*common.Claims)
	programs, err := service.ProgramRepo.FindSavedProgramByUserId(uint64(claims.UserID))
	if err != nil {
		return nil, err
	}
	return programs, nil
}

// FindSavedProgram implements service.UserService.
func (service UserServiceImpl) FindSavedProgram(c *gin.Context) ([]entity.Program, error) {
	claims := c.MustGet("claims").(*common.Claims)
	programs, err := service.ProgramRepo.FindSavedProgramByUserId(uint64(claims.UserID))
	if err != nil {
		return nil, err
	}
	return programs, nil
}

// UpdateUser implements service.UserService.
func (service UserServiceImpl) ResetPassword(c *gin.Context, credential model.ChangePasswordRequest) (*entity.UserAccount, error) {
	if err := common.Validate(credential); err != nil {
		return nil, err
	}
	user, _ := service.UserRepo.FindUserByEmail(credential.Email)
	if user.UserID == 0 {
		return nil, errors.New("user not found")
	}
	if credential.Email != user.Email {
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
