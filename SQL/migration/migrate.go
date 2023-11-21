package migration

import (
	"log"

	"github.com/xbklyn/getgoal-app/modules/label"
	"github.com/xbklyn/getgoal-app/modules/program"
)

func Migrate() {
	log.Default().Println("Migrating...")

	program.Migrate()
	label.Migrate()

	log.Default().Println("Migration complete")
}
