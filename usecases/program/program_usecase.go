package usecases

import "github.com/xbklyn/getgoal-app/entities"

type ProgramUsecase interface {
	FindAllPrograms() ([]entities.Program, error)
	FindProgramByID(id uint64) (*entities.Program, error)
	FindSearchProgram(text string) ([]entities.Program, error)
	FilterProgram(filter string) ([]entities.Program, error)
	Save(program *entities.Program, label_names []string) error
}
