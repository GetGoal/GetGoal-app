package entity

import (
	"time"

	"gorm.io/gorm"
)

type ExternalProvider struct {
	ExternalProviderID uint64     `gorm:"primary_key;auto_increment" json:"external_provider_id"`
	ProviderName       string     `gorm:"size:50;not null;unique" json:"provider_name"`
	CreatedAt          time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt          time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeleteAt           *time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"delete_at"`
}

func (ex *ExternalProvider) TableName() string {
	return "external_provider"
}

func (ex *ExternalProvider) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	ex.UpdatedAt = now
	return nil
}

func (ex *ExternalProvider) Migrate(db *gorm.DB) {
	db.AutoMigrate(&ExternalProvider{})
}
