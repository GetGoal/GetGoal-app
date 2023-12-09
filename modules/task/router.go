package task

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/common"
	"github.com/xbklyn/getgoal-app/modules/user_account"
)

func TaskRegister(router *gin.RouterGroup) {
	// router.POST("", ProgramCreate)
	router.GET("/to-do", TaskFromEmailAndDate)
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

func TaskFromEmailAndDate(c *gin.Context) {

	log.Default().Println("TaskFromEmailAndDate")
	taskByEmailAndDateValidator := NewGetTaskByEmailandDateValidator()
	if err := taskByEmailAndDateValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	log.Default().Println("email => ", taskByEmailAndDateValidator.getTaskByEmailAndDateModel.Email)

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
