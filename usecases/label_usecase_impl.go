package usecases

import (
	"github.com/xbklyn/getgoal-app/entities"
	"github.com/xbklyn/getgoal-app/repositories"
)

type labelUsecaseImpl struct {
	labelRepository repositories.LabelRepository
}

// FindLabelByID implements LabelUsecase.
func (u *labelUsecaseImpl) FindLabelByID(id int) (*entities.Label, error) {
	label, err := u.labelRepository.FindLabelByID(id)

	if err != nil {
		return nil, err
	}
	return &label, nil
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
