package main

import (
	"log"

	"github.com/xbklyn/getgoal-app/common"
	"github.com/xbklyn/getgoal-app/modules/label"
)

func init() {
	common.LoadEnvVariables()
	common.InitDB()
}

func main() {
	log.Default().Println("Migrating...")

	common.DB.AutoMigrate(&label.Label{})

	log.Default().Println("Migration complete")
}
