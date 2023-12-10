package task

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/common"
)

type GetTaskByEmailAndDateValidator struct {
	GetTaskByEmailAndDateTask struct {
		Email string `json:"email" binding:"required"`
		Date  string `json:"date" binding:"required"`
	} `json:"params" binding:"required"`

	getTaskByEmailAndDateModel GetTaskByEmailAndDateTask
}

func NewGetTaskByEmailandDateValidator() GetTaskByEmailAndDateValidator {
	return GetTaskByEmailAndDateValidator{}
}

func (s *GetTaskByEmailAndDateValidator) Bind(c *gin.Context) error {
	err := c.ShouldBindJSON(s)
	if err != nil {
		log.Default().Printf("Error binding JSON: %s\n", err.Error())
		return err
	}

	parsedDate, err := time.Parse("2006-01-02", s.GetTaskByEmailAndDateTask.Date)
	if err != nil {
		fmt.Printf("Error parsing date: %s\n", err.Error())
		return err
	}

	s.getTaskByEmailAndDateModel.Email = s.GetTaskByEmailAndDateTask.Email
	s.getTaskByEmailAndDateModel.Date = parsedDate

	fmt.Printf("Binded JSON payload: %+v\n", s.getTaskByEmailAndDateModel)

	return nil
}

type JoinProgramkValidator struct {
	BulkTask []struct {
		IsSetNotification int    `json:"is_set_noti"`
		StartTime         string `json:"start_time" binding:"required"`
		TimeBeforeNotify  int    `json:"time_before_notify" binding:"max=250"`
	} `json:"modifications" binding:"required"`
	UserEmail string `json:"email" binding:"required"`

	bulkTaskModel []Task
}

func NewBulkTaskValidator() JoinProgramkValidator {
	return JoinProgramkValidator{}
}

func (s *JoinProgramkValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, s)
	if err != nil {
		log.Default().Printf("Error binding JSON: %s\n", err.Error())
		return err
	}
	for _, task := range s.BulkTask {
		parseTime, err := time.Parse("2006-01-02 15:04:05", task.StartTime)
		if err != nil {

			fmt.Printf("Error parsing date: %s\n", err.Error())
			return err
		}
		s.bulkTaskModel = append(s.bulkTaskModel, Task{
			IsSetNotification: task.IsSetNotification,
			StartTime:         parseTime,
			TimeBeforeNotify:  task.TimeBeforeNotify,
		})
	}

	fmt.Printf("Binded JSON payload: %+v\n", s.bulkTaskModel)

	return nil
}
