package usecases

import "github.com/xbklyn/getgoal-app/entities"

type ActionUsecase interface {
	FindAllActions() ([]entities.ActionType, error)
	FindActionByID(id int) (*entities.ActionType, error)
}
