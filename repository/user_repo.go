package repository

import "github.com/xbklyn/getgoal-app/entity"

type UserRepo interface {
	FindUserByEmail(email string) (entity.UserAccount, error)
	FindUserByID(id uint64) (entity.UserAccount, error)
	Save(user *entity.UserAccount) error
	Update(id uint64, user entity.UserAccount) error
	// FetchProgramByUserId(id uint64) ([]entity.Program, error)
}
