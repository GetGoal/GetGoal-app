package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/modules/label"
	"github.com/xbklyn/getgoal-app/modules/program"
)

func GetRoutes() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/" + func() string {
		env := os.Getenv("ENV")
		if env == "prod" {
			return ""
		} else {
			return env
		}
	}())

	v1 := api.Group("/v1")

	//Label groups
	label.LabelAnonymousRegister(v1.Group("/labels"))
	label.LabelRegister(v1.Group("/labels"))

	//Program groups
	program.ProgramAnonymousRegister(v1.Group("/programs"))

	landing := api.Group("/ping")
	landing.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}
