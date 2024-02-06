package entities

import (
	"time"

	"gorm.io/gorm"
)

type UserAccount struct {
	UserID    uint64     `gorm:"column:user_id;primary_key;auto_increment" json:"user_id"`
	FirstName string     `gorm:"column:first_name;type:varchar(70);not null" json:"first_name"`
	LastName  string     `gorm:"column:last_name;type:varchar(70);not null" json:"last_name"`
	Email     string     `gorm:"column:email;type:varchar(100);not null" json:"email"`
	CreatedAt time.Time  `gorm:"column:created_at;not null;default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;not null;default:current_timestamp" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`

	//Relationship
	Tasks      *[]Task      `gorm:"foreignKey:UserAccountID" json:"tasks"`
	ActionType []ActionType `gorm:"many2many:user_program;foreignKey:user_id;joinForeignKey:user_account_id;References:ActionID;JoinReferences:ActionID" json:"action_types"`
}

func (user *UserAccount) TableName() string {
	return "user_account"
}

func (user *UserAccount) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	user.UpdatedAt = now
	return nil
}
