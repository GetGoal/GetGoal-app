package task

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/common"
)

func TaskRegister(router *gin.RouterGroup) {
	// router.POST("", ProgramCreate)
}
func TaskAnonymousRegister(router *gin.RouterGroup) {
	router.GET("", TaskList)
	router.GET("/:id", TaskDetail)
}

func TaskList(c *gin.Context) {
	tasks, err := FindAllTask()
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("Task", err))
		return
	}

	serializer := TasksSerializer{C: c, Tasks: tasks, Count: len(tasks)}
	c.JSON(http.StatusOK, gin.H{"Task": serializer.Response()})
}

func TaskDetail(c *gin.Context) {
	taskId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewError("Task", err))
		return
	}

	task, err := FindOneTask(&Task{TaskID: taskId})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("Task", err))
		return
	}

	serializer := TaskSerializer{C: c, Task: task}
	c.JSON(http.StatusOK, gin.H{"Task": serializer.Response()})

}
