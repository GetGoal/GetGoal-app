package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/modules/label"
)

func GetRoutes() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	label.LabelAnonymousRegister(v1.Group("/labels"))
	label.LabelRegister(v1.Group("/labels"))

	landing := r.Group("/pong")
	landing.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}
