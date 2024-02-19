package repository

type UserProgramRepo interface {
	Save(actionId uint64, programId uint64, userId uint64) error
}
