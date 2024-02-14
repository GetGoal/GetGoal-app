package repository

import "github.com/xbklyn/getgoal-app/entity"

type ProgramRepo interface {
	FindAllPrograms() ([]entity.Program, error)
	FindProgramByID(id uint64) (entity.Program, error)
	FindProgramByText(str string) ([]entity.Program, error)
	FindProgramByLabel(labels []string) ([]entity.Program, error)
	Save(program *entity.Program) (entity.Program, error)
	Update(id uint64, program entity.Program) (entity.Program, error)
	Delete(id uint64) error
}