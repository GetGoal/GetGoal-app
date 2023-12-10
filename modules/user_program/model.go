package user_program

import (
	"time"

	"github.com/xbklyn/getgoal-app/common"
)

type UserProgram struct {
	UserAccountID uint64     `gorm:"column:user_account_id;not null" json:"user_account_id"`
	ProgramID     uint64     `gorm:"column:program_id;not null" json:"program_id"`
	ActionID      uint64     `gorm:"column:action_id;not null" json:"action_id"`
	CreatedAt     time.Time  `gorm:"column:created_at;not null;default:current_timestamp" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"column:updated_at;not null;default:current_timestamp" json:"updated_at"`
	DeletedAt     *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

func Migrate() {
	db := common.GetDB()
	db.AutoMigrate(&UserProgram{})
}

// TableName sets the table name for the UserProgram model.
func (UserProgram) TableName() string {
	return "user_program"
}

func SaveOne(actionId uint64, programId uint64, userAccountId uint64) error {
	db := common.GetDB()

	// Create the new program
	if err := db.Create(&UserProgram{ActionID: actionId, ProgramID: programId, UserAccountID: userAccountId}).Error; err != nil {
		return err
	}

	return nil
}
