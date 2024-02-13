package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserProgram struct {
	UserAccountID uint64     `gorm:"column:user_account_id;not null" json:"user_account_id"`
	ProgramID     uint64     `gorm:"column:program_id;not null" json:"program_id"`
	ActionID      uint64     `gorm:"column:action_id;not null" json:"action_id"`
	CreatedAt     time.Time  `gorm:"column:created_at;not null;default:current_timestamp" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"column:updated_at;not null;default:current_timestamp" json:"updated_at"`
	DeletedAt     *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

func (userProgram *UserProgram) TableName() string {
	return "user_account"
}

func (up *UserProgram) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	up.UpdatedAt = now
	return nil
}
