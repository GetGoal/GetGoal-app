package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/modules/label"
)

func GetRoutes() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/" + func() string {
		env := os.Getenv("ENV")
		if env == "qa" {
			return "qa"
		} else if env == "dev" {
			return "dev"
		}
		return ""
	}())

	v1 := api.Group("/v1")
	label.LabelAnonymousRegister(v1.Group("/labels"))
	label.LabelRegister(v1.Group("/labels"))

	landing := api.Group("/ping")
	landing.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}
