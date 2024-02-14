package main

import (
	"github.com/xbklyn/getgoal-app/config"
	"github.com/xbklyn/getgoal-app/database"
	"github.com/xbklyn/getgoal-app/server"
)

// @title Customers API
// @version 1.0
// @description.markdown

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @schemes https http
func main() {
	cfg := config.ReadConfig("./")

	db := database.NewPostgresDB(&cfg)

	gin := server.NewGinServer(&cfg, db.GetDb())

	gin.Start()

	// server.FiberServer(&cfg, db.GetDb()).Start()
}
