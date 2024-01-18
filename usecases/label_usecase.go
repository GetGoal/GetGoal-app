package usecases

import "github.com/xbklyn/getgoal-app/entities"

type LabelUsecase interface {
	FindAllLabels() ([]entities.Label, error)
}
