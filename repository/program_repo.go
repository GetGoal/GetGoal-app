package repository

import "github.com/xbklyn/getgoal-app/entity"

type ProgramRepo interface {
	FindAllPrograms() ([]entity.Program, error)
	FindProgramByID(id uint64) (entity.Program, error)
	FindProgramByIDs(ids []uint64) ([]entity.Program, error)
	FindProgramByText(str string) ([]entity.Program, error)
	FindSavedProgramByUserId(id uint64) ([]entity.Program, error)
	FindJoinedProgramByUserId(id uint64) ([]entity.Program, error)
	FindProgramByLabel(labels []string) ([]entity.Program, error)
	FindProgramByLabeWithLimits(labels []string, limit int) ([]entity.Program, error)
	FetchProgramByUserId(id uint64) ([]entity.Program, error)
	Save(program *entity.Program) (entity.Program, error)
	Update(id uint64, program *entity.Program) error
	Delete(program *entity.Program) error
}
