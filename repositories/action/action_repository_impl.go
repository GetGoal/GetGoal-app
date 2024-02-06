package repositories

import (
	"log"

	"github.com/xbklyn/getgoal-app/entities"
	"gorm.io/gorm"
)

type actionRepositoryImpl struct {
	db *gorm.DB
}

// FindActionByID implements ActionRepository.
func (a *actionRepositoryImpl) FindActionByID(id int) (entities.ActionType, error) {
	log.Default().Println("Query action by ID")
	var action entities.ActionType

	err := a.db.Debug().Model(&entities.ActionType{}).First(&action, id).Error
	return action, err
}

// FindAllActions implements ActionRepository.
func (a *actionRepositoryImpl) FindAllActions() ([]entities.ActionType, error) {
	log.Default().Println("Query all actions")

	var actions []entities.ActionType

	err := a.db.Debug().Model(&entities.ActionType{}).Find(&actions).Error
	return actions, err
}

func NewActionRepositoryImpl(db *gorm.DB) ActionRepository {
	return &actionRepositoryImpl{db}
}
