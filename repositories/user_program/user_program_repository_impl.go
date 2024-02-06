package repositories

import (
	"github.com/xbklyn/getgoal-app/entities"
	"gorm.io/gorm"
)

type UserProgramRepositoryImpl struct {
	db *gorm.DB
}

// SaveOne implements UserProgramRepository.
func (u *UserProgramRepositoryImpl) SaveOne(actionId uint64, programId uint64, userAccountId uint64) error {
	if err := u.db.Create(&entities.UserProgram{ActionID: actionId, ProgramID: programId, UserAccountID: userAccountId}).Error; err != nil {
		return err
	}

	return nil
}

func NewUserProgramRepositoryImpl(db *gorm.DB) UserProgramRepository {
	return &UserProgramRepositoryImpl{db}
}
