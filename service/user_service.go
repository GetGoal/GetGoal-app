package service

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/entity"
	"github.com/xbklyn/getgoal-app/model"
)

type UserService interface {
	FindUserByEmail(c *gin.Context) (*entity.UserAccount, error)
	FindSavedProgram(c *gin.Context) ([]entity.Program, error)
	FindJoinedProgram(c *gin.Context) ([]entity.Program, error)
	FindDateWithTasks(c *gin.Context, date time.Time) ([]int, error)
	UpdateLabel(c *gin.Context, userModel model.UserModel) (*entity.UserAccount, error)
	ResetPassword(c *gin.Context, credential model.ChangePasswordRequest) (*entity.UserAccount, error)
}
