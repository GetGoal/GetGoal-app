package migrations

import (
	"github.com/xbklyn/getgoal-app/entities"
	"gorm.io/gorm"
)

func LabelMigrate(db *gorm.DB) {
	db.AutoMigrate(&entities.Label{})
}