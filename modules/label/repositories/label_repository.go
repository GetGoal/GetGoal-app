package repositories

import "github.com/xbklyn/getgoal-app/modules/label/entities"

type LabelRepository interface {
	FindAllLabels() ([]entities.Label, error)
}
