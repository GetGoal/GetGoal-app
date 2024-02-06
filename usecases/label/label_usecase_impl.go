package usecases

import (
	"github.com/xbklyn/getgoal-app/entities"
	lrepositories "github.com/xbklyn/getgoal-app/repositories/label"
)

type labelUsecaseImpl struct {
	labelRepository lrepositories.LabelRepository
}

// Save implements LabelUsecase.
func (u *labelUsecaseImpl) Save(label *entities.Label) error {
	err := u.labelRepository.Save(label)
	return err
}

// GetSeachLabel implements LabelUsecase.
func (u *labelUsecaseImpl) GetSearchLabel() ([]entities.Label, error) {
	labels, err := u.labelRepository.GetSearchLabel()

	if err != nil {
		return nil, err
	}
	return labels, nil
}

// FindLabelByID implements LabelUsecase.
func (u *labelUsecaseImpl) FindLabelByID(id uint64) (*entities.Label, error) {
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

func NewLabelUsecaseImpl(labelRepository lrepositories.LabelRepository) LabelUsecase {
	return &labelUsecaseImpl{labelRepository: labelRepository}
}
