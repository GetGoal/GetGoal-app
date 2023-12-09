package main

import (
	"github.com/xbklyn/getgoal-app/SQL/migration"
	"github.com/xbklyn/getgoal-app/common"
	"github.com/xbklyn/getgoal-app/routes"
)

func init() {
	common.LoadEnvVariables()
	common.SetTimeZone()
	common.InitDB()
}
func main() {

	r := routes.GetRoutes()

	migration.Migrate()
	r.Run()
}
