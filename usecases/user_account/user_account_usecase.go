package usecases

import "github.com/xbklyn/getgoal-app/entities"

type UserAccountUsecase interface {
	FindAllUser() ([]entities.UserAccount, error)
	FindOneUser(condition interface{}) (entities.UserAccount, error)
}
