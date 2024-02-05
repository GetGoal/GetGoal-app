package repositories

import (
	"log"

	"github.com/xbklyn/getgoal-app/config"
	"github.com/xbklyn/getgoal-app/entities"
	"gorm.io/gorm"
)

type labelRepositoryImpl struct {
	db *gorm.DB
}

// FindLabelByName implements LabelRepository.
func (*labelRepositoryImpl) FindLabelByName(name string) (entities.Label, error) {
	panic("unimplemented")
}

// Save implements LabelRepository.
func (*labelRepositoryImpl) Save(label *entities.Label) error {
	panic("unimplemented")
}

// GetSearchLabel implements LabelRepository.
func (l *labelRepositoryImpl) GetSearchLabel() ([]entities.Label, error) {
	search_limit := config.GetConfig().Search.LabelLimit
	log.Default().Printf("Search limit: %d \n", search_limit)

	log.Default().Println("Query search label")

	var labels []entities.Label
	err := l.db.Debug().Model(&entities.Label{}).Preload("Programs").Order("RANDOM()").Limit(search_limit).Find(&labels).Error

	return labels, err
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
