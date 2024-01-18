package repositories

import "github.com/xbklyn/getgoal-app/entities"

type LabelRepository interface {
	FindAllLabels() ([]entities.Label, error)
}
