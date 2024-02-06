package repositories

import "github.com/xbklyn/getgoal-app/entities"

type ActionRepository interface {
	FindAllActions() ([]entities.ActionType, error)
	FindActionByID(id int) (entities.ActionType, error)
}
