package service

import (
	"github.com/xbklyn/getgoal-app/entity"
	"github.com/xbklyn/getgoal-app/model"
)

type ProgramService interface {
	FindAllPrograms() ([]entity.Program, error)
	FindProgramByID(id uint64) (*entity.Program, error)
	FindProgramByText(str string) ([]entity.Program, error)
	FindProgramByLabel(labels []string) ([]entity.Program, error)
	Save(program model.ProgramCreateOrUpdate) (entity.Program, error)
	Update(id uint64, program model.ProgramCreateOrUpdate) (entity.Program, error)
	Delete(id uint64) error
}
