package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/modules/label"
)

func GetRoutes() *gin.Engine {
	r := gin.Default()

	r.RedirectTrailingSlash = false
	r.RemoveExtraSlash = true

	v1 := r.Group("/api/v1")
	label.LabelAnonymousRegister(v1.Group("/labels"))
	label.LabelRegister(v1.Group("/labels"))

	landing := r.Group("/api/ping")
	landing.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}
