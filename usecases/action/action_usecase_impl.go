package usecases

import (
	"github.com/xbklyn/getgoal-app/entities"
	arepository "github.com/xbklyn/getgoal-app/repositories/action"
)

type ActionUsecaseImpl struct {
	actionRepository arepository.ActionRepository
}

// FindActionByID implements ActionUsecase.
func (a *ActionUsecaseImpl) FindActionByID(id int) (*entities.ActionType, error) {
	action, err := a.actionRepository.FindActionByID(id)

	if err != nil {
		return nil, err
	}
	return &action, nil
}

// FindAllActions implements ActionUsecase.
func (a *ActionUsecaseImpl) FindAllActions() ([]entities.ActionType, error) {

	actions, err := a.actionRepository.FindAllActions()

	if err != nil {
		return nil, err
	}
	return actions, nil
}

func NewActionUsecaseImpl(actionRepository arepository.ActionRepository) ActionUsecase {
	return &ActionUsecaseImpl{
		actionRepository: actionRepository,
	}
}
