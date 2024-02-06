package usecases

import "github.com/xbklyn/getgoal-app/entities"

type LabelUsecase interface {
	FindAllLabels() ([]entities.Label, error)
	FindLabelByID(id uint64) (*entities.Label, error)
	GetSearchLabel() ([]entities.Label, error)
	Save(label *entities.Label) error
}
