package user_account

import (
	"time"

	"github.com/xbklyn/getgoal-app/common"
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
	Tasks *[]Task `gorm:"foreignKey:UserAccountID" json:"tasks"`
}

func Migrate() {
	db := common.GetDB()
	db.AutoMigrate(&UserAccount{}, &Task{})
}

func (user *UserAccount) TableName() string {
	return "user_account"
}

func (user *UserAccount) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	user.UpdatedAt = now
	return nil
}

func FindAllUsers() ([]UserAccount, error) {
	db := common.GetDB()

	var users []UserAccount

	err := db.Debug().Model(&UserAccount{}).Preload("Tasks").Find(&users).Error
	return users, err
}

func FindOneUser(condition interface{}) (UserAccount, error) {
	db := common.GetDB()

	var user UserAccount

	err := db.Debug().Model(&UserAccount{}).Preload("Tasks").Where(condition).First(&user).Error
	return user, err
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
