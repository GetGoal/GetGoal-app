package impl

import (
	"github.com/xbklyn/getgoal-app/entity"
	repository "github.com/xbklyn/getgoal-app/repository"
	"gorm.io/gorm"
)

func NewUserProgramRepoImpl(db *gorm.DB) repository.UserProgramRepo {
	return &UserProgramRepoImpl{db}
}

type UserProgramRepoImpl struct {
	db *gorm.DB
}

// FindActionByUserId implements repository.UserProgramRepo.
func (u *UserProgramRepoImpl) FindActionByUserId(userId uint64, actionId uint64) ([]entity.UserProgram, error) {
	var activities []entity.UserProgram
	err := u.db.Debug().Model(&entity.UserProgram{}).
		Where("user_account_id = ?", userId).
		Where("action_id = ?", actionId).
		Find(&activities).Error

	return activities, err
}

// Save implements repository.UserProgramRepo.
func (u *UserProgramRepoImpl) Save(actionId uint64, programId uint64, userAccountId uint64) error {
	err := u.db.Create(&entity.UserProgram{ActionID: actionId, ProgramID: programId, UserAccountID: userAccountId}).Error

	return err
}
