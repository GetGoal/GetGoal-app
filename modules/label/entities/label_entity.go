package entities

import (
	"time"

	"gorm.io/gorm"
)

type Label struct {
	LabelID   uint64     `gorm:"column:label_id;primary_key;auto_increment" json:"label_id"`
	LabelName string     `gorm:"column:label_name;type:varchar(50);not null" json:"label_name"`
	CreatedAt time.Time  `gorm:"column:created_at;not null;default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;not null;default:current_timestamp" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;index" json:"deleted_at"`
}

func (label *Label) TableName() string {
	return "label"
}

func (label *Label) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	label.UpdatedAt = now
	return nil
}
