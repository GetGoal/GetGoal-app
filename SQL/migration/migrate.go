package migration

import (
	"log"

	"github.com/xbklyn/getgoal-app/modules/label"
	"github.com/xbklyn/getgoal-app/modules/program"
	"github.com/xbklyn/getgoal-app/modules/task"
	"github.com/xbklyn/getgoal-app/modules/user_account"
)

func Migrate() {
	log.Default().Println("Migrating...")

	program.Migrate()
	label.Migrate()
	task.Migrate()
	user_account.Migrate()

	log.Default().Println("Migration complete")
}
