package main

import (
	"log"
	"time"

	"github.com/xbklyn/getgoal-app/SQL/migration"
	"github.com/xbklyn/getgoal-app/common"
	"github.com/xbklyn/getgoal-app/routes"
)

func init() {
	common.SetTimezone("Asia/Bangkok")
	t := common.GetTime(time.Now())
	log.Println(t)
	common.LoadEnvVariables()
	common.InitDB()
}
func main() {

	r := routes.GetRoutes()

	migration.Migrate()
	r.Run()
}
