package repositoryentity

import "github.com/xbklyn/getgoal-app/entity"

type LabelRepo interface {
	FindAllLabels() ([]entity.Label, error)
	FindLabelByID(id uint64) (entity.Label, error)
	FindLabelByName(name string) (entity.Label, error)
	GetSearchLabel() ([]entity.Label, error)
	Save(label *entity.Label) (entity.Label, error)
}