package migration

import (
	"log"

	"github.com/xbklyn/getgoal-app/common"
	"github.com/xbklyn/getgoal-app/modules/label"
)

func Migrate() {
	log.Default().Println("Migrating...")

	common.DB.AutoMigrate(&label.Label{})

	log.Default().Println("Migration complete")
}
