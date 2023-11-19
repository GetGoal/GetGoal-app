package migration

import (
	"log"

	"github.com/xbklyn/getgoal-app/common"
	"github.com/xbklyn/getgoal-app/modules/label"
	"github.com/xbklyn/getgoal-app/modules/program"
)

func Migrate() {
	log.Default().Println("Migrating...")

	common.DB.AutoMigrate(&label.Label{})
	common.DB.AutoMigrate(&program.Program{})

	log.Default().Println("Migration complete")
}
