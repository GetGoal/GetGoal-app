package usecases

import (
	"github.com/xbklyn/getgoal-app/entities"
	prepositories "github.com/xbklyn/getgoal-app/repositories/program"
)

type ProgramUsecaseImpl struct {
	programRepo prepositories.ProgramRepository
}

// FilterProgram implements ProgramUsecase.
func (p *ProgramUsecaseImpl) FilterProgram(filter string) ([]entities.Program, error) {
	programs, err := p.programRepo.FilterProgram(filter)

	if err != nil {
		return nil, err
	}
	return programs, nil
}

// FindAllPrograms implements ProgramUsecase.
func (p *ProgramUsecaseImpl) FindAllPrograms() ([]entities.Program, error) {
	programs, err := p.programRepo.FindAllPrograms()

	if err != nil {
		return nil, err
	}
	return programs, nil
}

// FindProgramByID implements ProgramUsecase.
func (p *ProgramUsecaseImpl) FindProgramByID(id uint64) (*entities.Program, error) {
	program, err := p.programRepo.FindOneProgram(id)

	if err != nil {
		return nil, err
	}
	return &program, nil
}

// FindSearchProgram implements ProgramUsecase.
func (p *ProgramUsecaseImpl) FindSearchProgram(text string) ([]entities.Program, error) {
	programs, err := p.programRepo.FindSearchProgram(text)

	if err != nil {
		return nil, err
	}
	return programs, nil
}

// Save implements ProgramUsecase.
func (p *ProgramUsecaseImpl) Save(program *entities.Program, label_names []string) error {
	err := p.programRepo.Save(program, label_names)

	if err != nil {
		return err
	}
	return nil
}

func NewProgramUsecaseImpl(programRepo prepositories.ProgramRepository) ProgramUsecase {
	return &ProgramUsecaseImpl{programRepo}
}
