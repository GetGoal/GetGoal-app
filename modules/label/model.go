package label

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/xbklyn/getgoal-app/common"
)

type Label struct {
	LabelID   uint64     `gorm:"column:label_id;primary_key;auto_increment" json:"label_id"`
	LabelName string     `gorm:"column:label_name;type:varchar(50);not null" json:"label_name"`
	CreatedAt time.Time  `gorm:"column:created_at;not null;default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;not null;default:current_timestamp" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;index" json:"deleted_at"`

	Programs []Program `gorm:"many2many:label_program;foreignKey:LabelID;joinForeignKey:LabelID;References:ProgramID;JoinReferences:ProgramID" json:"programs"`
}

type Program struct {
	ProgramID          uint64    `gorm:"column:program_id;primary_key;auto_increment" json:"program_id"`
	ProgramName        string    `gorm:"column:program_name;type:varchar(150);not null" json:"program_name"`
	Rating             float64   `gorm:"column:rating;not null;default:0" json:"rating"`
	MediaURL           string    `gorm:"column:media_url;type:varchar(255)" json:"media_url"`
	ProgramDescription string    `gorm:"column:program_description;type:varchar(250)" json:"program_description"`
	ExpectedTime       string    `gorm:"column:expected_time;type:varchar(30)" json:"expected_time"`
	UpdatedAt          time.Time `gorm:"column:updated_at;not null;default:current_timestamp" json:"updated_at"`
}

func Migrate() {
	db := common.GetDB()
	db.AutoMigrate(&Label{}, &Program{})
}

func (lable *Label) TableName() string {
	return "label"
}

func (program *Program) TableName() string {
	return "program"
}

func (label *Label) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	label.UpdatedAt = now
	return nil
}

func getSearchLimitFromEnv() int {
	// Get the limit from the environment variable
	limitStr := os.Getenv("SEARCH_LIMIT")

	// Default limit in case the environment variable is not set
	defaultLimit := 13

	// Convert the limit string to an integer
	limit, err := strconv.Atoi(limitStr)

	if err != nil {
		return defaultLimit
	}
	return limit
}

func FindAllLabel() ([]Label, error) {
	db := common.GetDB()

	var labels []Label

	err := db.Debug().Model(&Label{}).Preload("Programs").Find(&labels).Error

	for _, label := range labels {
		fmt.Printf("Label ID: %d\n", label.LabelID)
		fmt.Printf("Number of Programs: %d\n", len(label.Programs))
	}
	return labels, err
}

func FindSearchLabel() ([]Label, error) {
	db := common.GetDB()

	var labels []Label
	var limit int = getSearchLimitFromEnv()

	err := db.Order("RANDOM()").Limit(limit).Find(&labels).Error

	return labels, err
}

func FindOneLable(condition interface{}) (Label, error) {
	db := common.GetDB()

	var label Label

	err := db.Where(condition).First(&label).Error
	return label, err
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Create(data).Error
	return err
}
