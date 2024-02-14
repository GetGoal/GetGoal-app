package repository

import "github.com/xbklyn/getgoal-app/entity"

type UserRepo interface {
	FindUserByEmail(email string) (entity.UserAccount, error)
	FindUserByID(id uint64) (entity.UserAccount, error)
}
