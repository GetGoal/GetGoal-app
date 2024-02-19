package service

import (
	"github.com/xbklyn/getgoal-app/entity"
	"github.com/xbklyn/getgoal-app/model"
)

type LabelService interface {
	FindAllLabels() ([]entity.Label, error)
	FindLabelByID(id uint64) (*entity.Label, error)
	GetSearchLabel() ([]entity.Label, error)
	Save(label model.LabelRequest) (*entity.Label, error)
	Update(id uint64, label model.LabelRequest) (*entity.Label, error)
	Delete(id uint64) error
}
