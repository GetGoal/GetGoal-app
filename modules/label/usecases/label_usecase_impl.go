package usecases

import (
	"github.com/xbklyn/getgoal-app/modules/label/entities"
	"github.com/xbklyn/getgoal-app/modules/label/repositories"
)

type labelUsecaseImpl struct {
	labelRepository repositories.LabelRepository
}

// FindAllLabels implements LabelUsecase.
func (u *labelUsecaseImpl) FindAllLabels() ([]entities.Label, error) {
	labels, err := u.labelRepository.FindAllLabels()

	if err != nil {
		return nil, err
	}

	return labels, nil
}

func NewLabelUsecaseImpl(labelRepository repositories.LabelRepository) LabelUsecase {
	return &labelUsecaseImpl{labelRepository: labelRepository}
}
