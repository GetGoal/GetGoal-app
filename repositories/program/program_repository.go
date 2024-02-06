package repositories

import "github.com/xbklyn/getgoal-app/entities"

type ProgramRepository interface {
	FindAllPrograms() ([]entities.Program, error)
	FindOneProgram(condition interface{}) (entities.Program, error)
	FindProgramById(id uint64) (entities.Program, error)
	FindSearchProgram(text string) ([]entities.Program, error)
	FilterProgram(filter string) ([]entities.Program, error)
	Save(program *entities.Program, label_names []string) error
}
