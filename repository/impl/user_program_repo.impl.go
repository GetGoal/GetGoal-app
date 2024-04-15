package impl

import (
	"github.com/xbklyn/getgoal-app/entity"
	"github.com/xbklyn/getgoal-app/model"
	repository "github.com/xbklyn/getgoal-app/repository"
	"gorm.io/gorm"
)

func NewUserProgramRepoImpl(db *gorm.DB) repository.UserProgramRepo {
	return &UserProgramRepoImpl{db}
}

type UserProgramRepoImpl struct {
	db *gorm.DB
}

// GetStatistic implements repository.UserProgramRepo.
func (u *UserProgramRepoImpl) GetStatistic(programId uint64) (model.ProgramStat, error) {
	var stats model.ProgramStat

	err := u.db.Debug().Model(&entity.UserProgram{}).
		Select("program_id,COUNT(CASE WHEN action_id = 2 THEN 1 END) AS joined,COUNT(CASE WHEN action_id = 3 THEN 1 END) AS saved,COUNT(CASE WHEN action_id = 4 THEN 1 END) AS viewed,MAX(CASE WHEN action_id = 2 THEN created_at END) AS last_joined").
		Where("program_id = ?", programId).
		Group("program_id").
		Find(&stats).
		Error

	return stats, err
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
