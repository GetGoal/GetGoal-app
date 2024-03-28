package impl

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/xbklyn/getgoal-app/config"
	"github.com/xbklyn/getgoal-app/entity"
	repository "github.com/xbklyn/getgoal-app/repository"
	"gorm.io/gorm"
)

func NewlabelRepoImpl(db *gorm.DB) repository.LabelRepo {
	return &labelRepoImpl{db}
}

type labelRepoImpl struct {
	db *gorm.DB
}

// GetPreferenceLabel implements repository.LabelRepo.
func (l *labelRepoImpl) GetPreferenceLabel() ([]entity.Label, error) {

	pref_limit := config.GetConfig().Search.PreferenceLimit
	var labels []entity.Label
	err := l.db.Debug().Model(&entity.Label{}).Preload("Programs").Order("RANDOM()").Limit(pref_limit).Find(&labels).Error

	return labels, err
}

// Delete implements repositoryentity.LabelRepo.
func (l *labelRepoImpl) Delete(id uint64) error {
	lErr := l.db.Where("label_id = ?", id).Delete(&entity.Label{}).Error
	if lErr != nil {
		return lErr
	}
	return nil
}

// Update implements repositoryentity.LabelRepo.
func (l *labelRepoImpl) Update(id uint64, label entity.Label) (entity.Label, error) {
	log.Default().Printf("Update label: %s \n", label.LabelName)

	err := l.db.Debug().Model(&entity.Label{}).Where("label_id = ?", id).Updates(&entity.Label{LabelName: label.LabelName}).Error
	return label, err
}

// FindLabelByName implements LabelRepository.
func (l *labelRepoImpl) FindLabelByName(name string) (entity.Label, error) {
	log.Default().Printf("Query label by name: %s \n", name)
	var label entity.Label

	err := l.db.Debug().Model(&entity.Label{}).Preload("Programs").Where("LOWER(REPLACE(label_name, ' ', '')) = LOWER(?)", strings.ToLower(strings.ReplaceAll(name, " ", ""))).First(&label).Error

	return label, err
}

// Save implements LabelRepository.
func (l *labelRepoImpl) Save(labelModel *entity.Label) (entity.Label, error) {
	fmt.Print("Save label")
	existed, err := l.FindLabelByName(labelModel.LabelName)

	if err == nil {
		return existed, errors.New("label already existed")
	}

	er := l.db.Debug().Create(labelModel).Error
	return *labelModel, er
}

// GetSearchLabel implements LabelRepository.
func (l *labelRepoImpl) GetSearchLabel() ([]entity.Label, error) {
	fmt.Print("Get search label")
	search_limit := config.GetConfig().Search.LabelLimit

	log.Default().Println("Query search label")

	var labels []entity.Label
	err := l.db.Debug().Model(&entity.Label{}).Preload("Programs").Order("RANDOM()").Limit(search_limit).Find(&labels).Error

	return labels, err
}

// FindLabelByID implements LabelRepository.
func (l *labelRepoImpl) FindLabelByID(id uint64) (entity.Label, error) {
	log.Default().Printf("Query label by id: %d \n", id)

	var label entity.Label

	err := l.db.Debug().Model(&entity.Label{}).Preload("Programs").First(&label, id).Error

	return label, err
}

// FindAllLabels implements LabelRepository.
func (l *labelRepoImpl) FindAllLabels() ([]entity.Label, error) {
	log.Default().Println("Query all labels")

	var labels []entity.Label

	err := l.db.Debug().Model(&entity.Label{}).Preload("Programs").Find(&labels).Error
	return labels, err
}
