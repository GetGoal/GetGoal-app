package program

import (
	"fmt"
	"log"
	"time"

	"github.com/xbklyn/getgoal-app/common"
	"gorm.io/gorm"
)

type Program struct {
	ProgramID          uint64     `gorm:"column:program_id;primary_key;auto_increment" json:"program_id"`
	ProgramName        string     `gorm:"column:program_name;type:varchar(150);not null" json:"program_name"`
	Rating             float64    `gorm:"column:rating;not null;default:0" json:"rating"`
	MediaURL           string     `gorm:"column:media_url;type:varchar(255)" json:"media_url"`
	ProgramDescription string     `gorm:"column:program_description;type:varchar(250)" json:"program_description"`
	ExpectedTime       string     `gorm:"column:expected_time;type:varchar(30)" json:"expected_time"`
	CreatedAt          time.Time  `gorm:"column:created_at;not null;default:current_timestamp" json:"created_at"`
	UpdatedAt          time.Time  `gorm:"column:updated_at;not null;default:current_timestamp" json:"updated_at"`
	DeletedAt          *time.Time `gorm:"column:deleted_at;index" json:"deleted_at"`

	// Relationships
	Labels []Label `gorm:"many2many:label_program;foreignKey:ProgramID;joinForeignKey:ProgramID;References:LabelID;JoinReferences:LabelID" json:"labels"`
}

type Label struct {
	LabelID   uint64 `gorm:"column:label_id;primary_key;auto_increment" json:"label_id"`
	LabelName string `gorm:"column:label_name;type:varchar(50);not null" json:"label_name"`
}

func Migrate() {
	db := common.GetDB()
	db.AutoMigrate(&Program{}, &Label{})
}

func (program *Program) TableName() string {
	return "program"
}

func (label *Label) TableName() string {
	return "label"
}

// func (labelProgram *LabelProgram) TableName() string {
// 	return "label_program"
// }

func (program *Program) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	program.UpdatedAt = now
	return nil
}

func FindAllProgram() ([]Program, error) {
	db := common.GetDB()

	var program Program
	if result := db.Preload("Labels").First(&program); result.Error != nil {
		log.Fatalf("Error during Preload: %v", result.Error)
	} else {
		fmt.Printf("Preload Result: %+v\n", program)
	}

	var programs []Program

	err := db.Debug().Model(&Program{}).Preload("Labels").Find(&programs).Error

	for _, program := range programs {
		fmt.Printf("Program ID: %d\n", program.ProgramID)
		fmt.Printf("Number of Labels: %d\n", len(program.Labels))
	}
	return programs, err
}
func FindOneProgram(condition interface{}) (Program, error) {
	db := common.GetDB()

	var program Program

	err := db.Where(condition).First(&program).Error
	return program, err
}
func SaveOne(program *Program) error {
	db := common.GetDB()

	err := db.Create(program).Error
	return err
}
