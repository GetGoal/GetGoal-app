package repositories

import "github.com/xbklyn/getgoal-app/entities"

type LabelRepository interface {
	FindAllLabels() ([]entities.Label, error)
	FindLabelByID(id int) (entities.Label, error)
	FindLabelByName(name string) (entities.Label, error)
	GetSearchLabel() ([]entities.Label, error)
	Save(label *entities.Label) error
}
