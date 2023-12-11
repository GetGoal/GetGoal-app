package task

import (
	"fmt"
	"time"

	"github.com/xbklyn/getgoal-app/common"
	"github.com/xbklyn/getgoal-app/modules/program"
	"github.com/xbklyn/getgoal-app/modules/user_account"
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

func Migrate() {
	db := common.GetDB()
	db.AutoMigrate(&Task{}, &Program{}, &UserAccount{})
}

func (task *Task) TableName() string {
	return "task"
}

func (task *Task) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	task.UpdatedAt = now
	return nil
}

func FindAllTask() ([]Task, error) {
	db := common.GetDB()

	var tasks []Task

	err := db.Debug().Model(&Task{}).Preload("Program").Preload("UserAccount").Find(&tasks).Error
	return tasks, err
}

func FindOneTask(condition interface{}) (Task, error) {
	db := common.GetDB()

	var task Task

	err := db.Debug().Model(&Task{}).Preload("Program").Preload("UserAccount").Where(condition).First(&task).Error
	return task, err
}

func FindTaskByDateAndEmail(condition *Task) ([]Task, error) {
	db := common.GetDB()

	var tasks []Task

	err := db.Debug().Model(&Task{}).Preload("Program").Preload("UserAccount").
		Where("DATE(start_time) = DATE(?)", condition.StartTime).
		Where("user_account_id = ?", condition.UserAccountID).
		Find(&tasks).Error
	return tasks, err
}

func GetTaskByProgramId(program_id uint64) ([]Task, error) {
	db := common.GetDB()

	var tasks []Task

	err := db.Debug().Model(&Task{}).Preload("UserAccount").Where("program_id = ?", program_id).Order("start_time ASC").Find(&tasks).Error
	return tasks, err
}

func SaveOne(task *Task) error {
	db := common.GetDB()

	if err := db.Create(task).Error; err != nil {
		return err
	}
	fmt.Println("Generated Task ID:", task.TaskID)
	return nil
}

func BindUser(user user_account.UserAccount) UserAccount {
	return UserAccount{
		UserID:    user.UserID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

func BindProgram(program program.Program) Program {
	return Program{
		ProgramID:          program.ProgramID,
		ProgramName:        program.ProgramName,
		Rating:             program.Rating,
		MediaURL:           program.MediaURL,
		ProgramDescription: program.ProgramDescription,
		ExpectedTime:       program.ExpectedTime,
	}
}
