package usecases

type UserProgramUsecase interface {
	SaveOne(actionId uint64, programId uint64, userAccountId uint64) error
}
