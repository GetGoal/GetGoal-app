package repository

import "github.com/xbklyn/getgoal-app/entity"

type UserProgramRepo interface {
	Save(actionId uint64, programId uint64, userId uint64) error
	FindActionByUserId(userId uint64, actionId uint64) ([]entity.UserProgram, error)
}
