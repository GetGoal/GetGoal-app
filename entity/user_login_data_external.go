package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserLoginDataExternal struct {
	UserLoginDataExternalID uint64     `gorm:"primary_key;auto_increment" json:"user_login_data_external_id"`
	UserID                  uint64     `gorm:"not null" json:"user_id"`
	ExternalProviderID      uint64     `gorm:"not null" json:"external_provider_id"`
	CreatedAt               time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt               time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeleteAt                *time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"delete_at"`
}

func (ex *UserLoginDataExternal) TableName() string {
	return "user_login_data_external"
}

func (ex *UserLoginDataExternal) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	ex.UpdatedAt = now
	return nil
}

func (ex *UserLoginDataExternal) Migrate(db *gorm.DB) {
	db.AutoMigrate(&UserLoginDataExternal{})
}
