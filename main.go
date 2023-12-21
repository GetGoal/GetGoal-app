package main

import (
	"github.com/xbklyn/getgoal-app/config"
	"github.com/xbklyn/getgoal-app/database"
	"github.com/xbklyn/getgoal-app/server"
)

func main() {
	cfg := config.GetConfig()

	db := database.NewPostgresDB(&cfg)

	server.NewGinServer(&cfg, db.GetDb()).Start()
}
