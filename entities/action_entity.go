package entities

import (
	"time"

	"github.com/xbklyn/getgoal-app/common"
	"gorm.io/gorm"
)

type ActionType struct {
	ActionID   uint64     `gorm:"column:action_id;primary_key;auto_increment" json:"action_id"`
	ActionName string     `gorm:"column:action_name;type:varchar(50);not null" json:"action_name"`
	CreatedAt  time.Time  `gorm:"column:created_at;not null;default:current_timestamp" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"column:updated_at;not null;default:current_timestamp" json:"updated_at"`
	DeletedAt  *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

func Migrate() {
	db := common.GetDB()
	db.AutoMigrate(&ActionType{})
}

func (actionType *ActionType) TableName() string {
	return "action_type"
}

func (actionType *ActionType) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	actionType.UpdatedAt = now
	return nil
}
