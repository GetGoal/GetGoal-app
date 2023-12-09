package task

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/common"
	"github.com/xbklyn/getgoal-app/modules/user_account"
)

func TaskRegister(router *gin.RouterGroup) {
	router.GET("/to-do", TaskFromEmailAndDate)
}
func TaskAnonymousRegister(router *gin.RouterGroup) {
	router.GET("", TaskList)
	router.GET("/:id", TaskDetail)
	router.GET("/plan/:program_id", TaskPlanning)
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

func TaskFromEmailAndDate(c *gin.Context) {

	taskByEmailAndDateValidator := NewGetTaskByEmailandDateValidator()
	if err := taskByEmailAndDateValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	user, err := user_account.FindOneUser(&user_account.UserAccount{Email: taskByEmailAndDateValidator.getTaskByEmailAndDateModel.Email})

	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("User", err))
		return
	}

	tasks, err := FindTaskByDateAndEmail(&Task{UserAccountID: int(user.UserID), StartTime: taskByEmailAndDateValidator.getTaskByEmailAndDateModel.Date})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("Task", err))
		return
	}

	serializer := TasksSerializer{C: c, Tasks: tasks, Count: len(tasks)}
	c.JSON(http.StatusOK, gin.H{"Task": serializer.Response()})
}

func TaskPlanning(c *gin.Context) {
	programId, err := strconv.ParseUint(c.Param("program_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewError("Program", err))
		return
	}

	tasks, err := GetTaskByProgramId(programId)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("Task", err))
		return
	}

	now := time.Now()
	fmt.Println(now)
	var originalStartTime time.Time

	for i := range tasks {
		originalTime := tasks[i].StartTime.Format("15:04:05")

		if i > 0 {

			diff := tasks[i].StartTime.Sub(originalStartTime)
			newStartTime := tasks[i-1].StartTime.Add(diff)

			originalStartTime = tasks[i].StartTime
			tasks[i].StartTime = newStartTime
			continue
		} else {

			newStartTime, _ := time.Parse(time.RFC3339, now.Format("2006-01-02")+"T"+originalTime+"Z")

			originalStartTime = tasks[i].StartTime
			tasks[i].StartTime = newStartTime
			continue
		}
	}

	serializer := TasksPlanningSerializer{C: c, Tasks: tasks, Count: len(tasks)}
	c.JSON(http.StatusOK, gin.H{"Task": serializer.Response()})
}
