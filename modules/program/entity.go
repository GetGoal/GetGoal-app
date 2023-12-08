package program

import "time"

type Label struct {
	LabelID   uint64 `gorm:"column:label_id;primary_key;auto_increment" json:"label_id"`
	LabelName string `gorm:"column:label_name;type:varchar(50);not null" json:"label_name"`
}
type Task struct {
	TaskID            uint64     `gorm:"column:task_id;primary_key;auto_increment" json:"task_id"`
	TaskName          string     `gorm:"column:task_name;type:varchar(150);not null" json:"task_name"`
	TaskStatus        int        `gorm:"column:task_status;not null" json:"task_status"`
	UserAccountID     int        `gorm:"column:user_account_id;not null" json:"user_account_id"`
	IsSetNotification int        `gorm:"column:is_set_noti;not null" json:"is_set_noti"`
	StartTime         time.Time  `gorm:"column:start_time;not null" json:"start_time"`
	EndTime           *time.Time `gorm:"column:end_time" json:"end_time"`
	ProgramID         int        `gorm:"column:program_id" json:"program_id"`
	Category          string     `gorm:"column:category;type:varchar(50)" json:"category"`
	TimeBeforeNotify  int        `gorm:"column:time_before_notify" json:"time_before_notify"`
	TaskDescription   string     `gorm:"column:task_description;type:varchar(250)" json:"task_description"`
	Link              string     `gorm:"column:link;type:varchar(255)" json:"link"`
	MediaURL          string     `gorm:"column:media_url;type:varchar(255)" json:"media_url"`
	UpdatedAt         time.Time  `gorm:"column:updated_at;not null;default:current_timestamp" json:"updated_at"`
}

type UserAccount struct {
	UserID    uint64    `gorm:"column:user_id;primary_key;auto_increment" json:"user_id"`
	Email     string    `gorm:"column:email;type:varchar(100);not null" json:"email"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:current_timestamp" json:"updated_at"`

	ActionType []ActionType `gorm:"many2many:user_program;foreignKey:user_id;joinForeignKey:user_account_id;References:ActionID;JoinReferences:ActionID" json:"action_types"`
}

type ActionType struct {
	ActionID   uint64 `gorm:"column:action_id;primary_key;auto_increment" json:"action_id"`
	ActionName string `gorm:"column:action_name;type:varchar(50);not null" json:"action_name"`
}

func (label *Label) TableName() string {
	return "label"
}

func (task *Task) TableName() string {
	return "task"
}

func (userAccount *UserAccount) TableName() string {
	return "user_account"
}

func (actionType ActionType) TableName() string {
	return "action_type"
}
