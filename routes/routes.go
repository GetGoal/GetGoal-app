package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/modules/label"
	"github.com/xbklyn/getgoal-app/modules/program"
	"github.com/xbklyn/getgoal-app/modules/task"
	"github.com/xbklyn/getgoal-app/modules/user_account"
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
	program.ProgramRegister(v1.Group("/programs"))

	//Task groups
	task.TaskAnonymousRegister(v1.Group("/tasks"))
	task.TaskRegister(v1.Group("/tasks"))

	//User groups
	user_account.UserAnonymousRegister(v1.Group("/users"))
	user_account.UserRegister(v1.Group("/users"))

	//Ping group
	landing := api.Group("/ping")
	landing.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}
