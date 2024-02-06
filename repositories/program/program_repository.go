package repositories

import "github.com/xbklyn/getgoal-app/entities"

type ProgramRepository interface {
	FindAllPrograms() ([]entities.Program, error)
	FindOneProgram(condition interface{}) (entities.Program, error)
	FindSearchProgram(text string) ([]entities.Program, error)
	FilterProgram(filter string) ([]entities.Program, error)
	Save(program *entities.Program, label_names []string) error
}
