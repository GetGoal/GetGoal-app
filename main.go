package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.InitDB()
}
func main() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
