package entity

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	TaskID            uint64     `gorm:"column:task_id;primary_key;auto_increment" json:"task_id"`
	TaskName          string     `gorm:"column:task_name;type:varchar(150);not null" json:"task_name"`
	TaskStatus        int        `gorm:"column:task_status;not null" json:"task_status"`
	IsSetNotification int        `gorm:"column:is_set_noti;not null" json:"is_set_noti"`
	StartTime         time.Time  `gorm:"column:start_time;not null" json:"start_time"`
	EndTime           *time.Time `gorm:"column:end_time" json:"end_time"`
	Category          string     `gorm:"column:category;type:varchar(50)" json:"category"`
	TimeBeforeNotify  int        `gorm:"column:time_before_notify" json:"time_before_notify"`
	TaskDescription   string     `gorm:"column:task_description;type:varchar(250)" json:"task_description"`
	Link              string     `gorm:"column:link;type:varchar(255)" json:"link"`
	MediaURL          string     `gorm:"column:media_url;type:varchar(255)" json:"media_url"`
	CreatedAt         time.Time  `gorm:"column:created_at;not null;default:current_timestamp" json:"created_at"`
	UpdatedAt         time.Time  `gorm:"column:updated_at;not null;default:current_timestamp" json:"updated_at"`
	DeletedAt         *time.Time `gorm:"column:deleted_at" json:"deleted_at"`

	// Relationship
	ProgramID *int     `gorm:"column:program_id" json:"program_id"`
	Program   *Program `gorm:"foreignKey:ProgramID;references:program_id" json:"program"`

	UserAccountID int         `gorm:"column:user_account_id;not null" json:"user_account_id"`
	UserAccount   UserAccount `gorm:"foreignKey:UserID;references:user_account_id" json:"owner"`
}

func (task *Task) TableName() string {
	return "task"
}

func (task *Task) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	task.UpdatedAt = now
	return nil
}
