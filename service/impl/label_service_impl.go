package impl

import (
	"github.com/xbklyn/getgoal-app/common"
	"github.com/xbklyn/getgoal-app/entity"
	"github.com/xbklyn/getgoal-app/model"
	repository "github.com/xbklyn/getgoal-app/repository"
	"github.com/xbklyn/getgoal-app/service"
)

type labelServiceImpl struct {
	LabelRepo repository.LabelRepo
}

// FindAllLabels implements service.LabelService.
func (service *labelServiceImpl) FindAllLabels() ([]entity.Label, error) {
	labels, err := service.LabelRepo.FindAllLabels()
	if err != nil {
		return nil, err
	}

	return labels, nil
}

// FindLabelByID implements service.LabelService.
func (service *labelServiceImpl) FindLabelByID(id uint64) (*entity.Label, error) {
	label, err := service.LabelRepo.FindLabelByID(id)

	if err != nil {
		return nil, err
	}
	return &label, nil
}

// GetSearchLabel implements service.LabelService.
func (service *labelServiceImpl) GetSearchLabel() ([]entity.Label, error) {
	labels, err := service.LabelRepo.GetSearchLabel()

	if err != nil {
		return nil, err
	}
	return labels, nil
}

// Save implements service.LabelService.
func (service *labelServiceImpl) Save(labelModel model.LabelRequest) (*entity.Label, error) {
	err := common.Validate(labelModel)
	if err != nil {
		return nil, err
	}
	labelE := entity.Label{
		LabelName: labelModel.LabelName,
	}

	label, serviceErr := service.LabelRepo.Save(&labelE)
	return &label, serviceErr
}

func NewLabelServiceImpl(labelRepo *repository.LabelRepo) service.LabelService {
	return &labelServiceImpl{*labelRepo}
}