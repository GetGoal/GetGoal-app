package repositories

import (
	"log"

	"github.com/xbklyn/getgoal-app/modules/label/entities"
	"gorm.io/gorm"
)

type labelRepositoryImpl struct {
	db *gorm.DB
}

// FindAllLabels implements LabelRepository.
func (l *labelRepositoryImpl) FindAllLabels() ([]entities.Label, error) {
	log.Default().Println("Query all labels")

	var labels []entities.Label

	err := l.db.Debug().Model(&entities.Label{}).Find(&labels).Error
	if err != nil {
		log.Default().Printf("Error when query all labels: %s \n", err)
		return nil, err
	}

	return labels, nil
}

func NewLabelRepositoryImpl(db *gorm.DB) LabelRepository {
	return &labelRepositoryImpl{db}
}
