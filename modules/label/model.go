package label

import (
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
}

func (Label) TableName() string {
	return "labels"
}

func (label *Label) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	label.UpdatedAt = now
	return nil
}

func FindAllLabel() ([]Label, error) {
	db := common.GetDB()

	var labels []Label

	err := db.Find(&labels).Error
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
	err := db.Save(data).Error
	return err
}
