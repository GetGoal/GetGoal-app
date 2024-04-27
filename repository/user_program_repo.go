package repository

import (
	"github.com/xbklyn/getgoal-app/entity"
	"github.com/xbklyn/getgoal-app/model"
)

type UserProgramRepo interface {
	Save(actionId uint64, programId uint64, userId uint64) error
	FindActionByUserId(userId uint64, actionId uint64) ([]entity.UserProgram, error)
	FindUserProgramByProgramId(programId uint64) (entity.UserProgram, error)
	GetStatistic(programId uint64) (model.ProgramStat, error)
}
