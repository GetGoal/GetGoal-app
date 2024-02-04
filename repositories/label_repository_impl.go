package repositories

import (
	"log"

	"github.com/xbklyn/getgoal-app/entities"
	"gorm.io/gorm"
)

type labelRepositoryImpl struct {
	db *gorm.DB
}

// FindLabelByID implements LabelRepository.
func (l *labelRepositoryImpl) FindLabelByID(id int) (entities.Label, error) {
	log.Default().Printf("Query label by id: %d \n", id)

	var label entities.Label

	err := l.db.Debug().Model(&entities.Label{}).Preload("Programs").First(&label, id).Error
	return label, err
}

// FindAllLabels implements LabelRepository.
func (l *labelRepositoryImpl) FindAllLabels() ([]entities.Label, error) {
	log.Default().Println("Query all labels")

	var labels []entities.Label

	err := l.db.Debug().Model(&entities.Label{}).Preload("Programs").Find(&labels).Error
	return labels, err
}

func NewLabelRepositoryImpl(db *gorm.DB) LabelRepository {
	return &labelRepositoryImpl{db}
}
