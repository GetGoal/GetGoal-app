package repositories

import "github.com/xbklyn/getgoal-app/entities"

type UserAccountRepository interface {
	FindAllUser() ([]entities.UserAccount, error)
	FindOneUser(condition interface{}) (entities.UserAccount, error)
}
