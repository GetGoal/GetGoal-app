package repository

import (
	"time"

	"github.com/xbklyn/getgoal-app/entity"
	"github.com/xbklyn/getgoal-app/model"
)

type UserRepo interface {
	FindUserByEmail(email string) (entity.UserAccount, error)
	FindUserByID(id uint64) (entity.UserAccount, error)
	FindDateWithTasks(date time.Time, id uint64) ([]model.DateHasTask, error)
	Save(user *entity.UserAccount) error
	Update(id uint64, user entity.UserAccount) error
}
