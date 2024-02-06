package repositories

import (
	"errors"
	"log"
	"strings"

	"github.com/xbklyn/getgoal-app/config"
	"github.com/xbklyn/getgoal-app/entities"
	"gorm.io/gorm"
)

type labelRepositoryImpl struct {
	db *gorm.DB
}

// FindLabelByName implements LabelRepository.
func (l *labelRepositoryImpl) FindLabelByName(name string) (entities.Label, error) {
	log.Default().Printf("Query label by name: %s \n", name)
	var label entities.Label

	err := l.db.Debug().Model(&entities.Label{}).Preload("Programs").Where("LOWER(REPLACE(label_name, ' ', '')) = LOWER(?)", strings.ToLower(strings.ReplaceAll(name, " ", ""))).First(&label).Error

	return label, err
}

// Save implements LabelRepository.
func (l *labelRepositoryImpl) Save(label *entities.Label) error {
	if _, err := l.FindLabelByName(label.LabelName); err != nil {

		err := l.db.Debug().Create(label).Error
		return err
	}

	return errors.New("label with this name already existed")
}

// GetSearchLabel implements LabelRepository.
func (l *labelRepositoryImpl) GetSearchLabel() ([]entities.Label, error) {
	search_limit := config.GetConfig().Search.LabelLimit

	log.Default().Println("Query search label")

	var labels []entities.Label
	err := l.db.Debug().Model(&entities.Label{}).Preload("Programs").Order("RANDOM()").Limit(search_limit).Find(&labels).Error

	return labels, err
}

// FindLabelByID implements LabelRepository.
func (l *labelRepositoryImpl) FindLabelByID(id uint64) (entities.Label, error) {
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
