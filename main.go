package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/common"
	"github.com/xbklyn/getgoal-app/routes"
)

func init() {
	common.LoadEnvVariables()
	common.InitDB()
}
func main() {

	r := routes.GetRoutes()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
