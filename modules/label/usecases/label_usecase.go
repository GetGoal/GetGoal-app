package usecases

import "github.com/xbklyn/getgoal-app/modules/label/entities"

type LabelUsecase interface {
	FindAllLabels() ([]entities.Label, error)
}
