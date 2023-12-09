package task

import "time"

type Program struct {
	ProgramID          uint64    `gorm:"column:program_id;primary_key;auto_increment" json:"program_id"`
	ProgramName        string    `gorm:"column:program_name;type:varchar(150);not null" json:"program_name"`
	Rating             float64   `gorm:"column:rating;not null;default:0" json:"rating"`
	MediaURL           string    `gorm:"column:media_url;type:varchar(255)" json:"media_url"`
	ProgramDescription string    `gorm:"column:program_description;type:varchar(250)" json:"program_description"`
	ExpectedTime       string    `gorm:"column:expected_time;type:varchar(30)" json:"expected_time"`
	UpdatedAt          time.Time `gorm:"column:updated_at;not null;default:current_timestamp" json:"updated_at"`
}

type UserAccount struct {
	UserID    uint64    `gorm:"column:user_id;primary_key;auto_increment" json:"user_id"`
	FirstName string    `gorm:"column:first_name;type:varchar(70);not null" json:"first_name"`
	LastName  string    `gorm:"column:last_name;type:varchar(70);not null" json:"last_name"`
	Email     string    `gorm:"column:email;type:varchar(100);not null" json:"email"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:current_timestamp" json:"updated_at"`
}

func (program *Program) TableName() string {
	return "program"
}

func (userAccount *UserAccount) TableName() string {
	return "user_account"
}

type GetTaskByEmailAndDateTask struct {
	Email string
	Date  time.Time
}
