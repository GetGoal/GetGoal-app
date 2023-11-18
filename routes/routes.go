package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/modules/label"
)

func GetRoutes() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/" + func() string {
		if os.Getenv("ENV") == "qa" {
			return "qa"
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
