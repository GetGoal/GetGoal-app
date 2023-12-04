package action

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

func FindAllActions() ([]ActionType, error) {
	db := common.GetDB()

	var actions []ActionType

	err := db.Debug().Model(&ActionType{}).Find(&actions).Error
	return actions, err
}

func FindOneAction(condition interface{}) (ActionType, error) {
	db := common.GetDB()

	var action ActionType

	err := db.Debug().Model(&ActionType{}).Where(condition).First(&action).Error
	return action, err
}

// func SaveOne(program *Program, labelNames []string) error {
// 	db := common.GetDB()

// 	// Create the new program
// 	if err := db.Create(program).Error; err != nil {
// 		return err
// 	}

// 	var labels []Label
// 	if len(labelNames) > 0 {
// 		for _, labelName := range labelNames {

// 			labelModel, err := getOrCreateLabel(db, labelName)
// 			if err != nil {
// 				return err
// 			}

// 			labels = append(labels, *labelModel)
// 		}
// 	}

// 	if err := db.Debug().Model(&program).Association("Labels").Append(labels); err != nil {
// 		return fmt.Errorf("failed to associate labels with program: %v", err)
// 	}

// 	return nil
// }

// func getOrCreateLabel(db *gorm.DB, labelName string) (*Label, error) {

// 	existingLabel, err := label.FindOneLableByName(labelName)
// 	if err != nil {
// 		fmt.Println("No label found with name: " + labelName)
// 		newLabel := label.Label{LabelName: labelName}
// 		if err := label.SaveOne(&newLabel); err != nil {
// 			return nil, err
// 		}
// 		fmt.Printf("Returning new label: %v", newLabel)
// 		return &Label{LabelID: newLabel.LabelID, LabelName: labelName}, nil
// 	} else {

// 		return &Label{LabelID: existingLabel.LabelID, LabelName: existingLabel.LabelName}, nil
// 	}

// }

// func FindSearchProgram(text string) ([]Program, error) {
// 	db := common.GetDB()

// 	var programs []Program

// 	err := db.Debug().Where("program_name ILIKE ?", "%"+text+"%").Find(&programs).Error

// 	return programs, err
// }

// func FilterProgram(filter string) ([]Program, error) {
// 	db := common.GetDB()

// 	var programs []Program
// 	err := db.Debug().Model(&Program{}).Joins("JOIN label_program ON program.program_id = label_program.program_id").
// 		Joins("JOIN label ON label_program.label_id = label.label_id AND label.label_name = ?", filter).
// 		Preload("Labels", "label_name = ?", filter).
// 		Find(&programs).Error

// 	return programs, err
// }
