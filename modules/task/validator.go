package task

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
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

	parsedDate, err := time.Parse(time.RFC3339, s.GetTaskByEmailAndDateTask.Date)
	if err != nil {
		fmt.Printf("Error parsing date: %s\n", err.Error())
		return err
	}

	s.getTaskByEmailAndDateModel.Email = s.GetTaskByEmailAndDateTask.Email
	s.getTaskByEmailAndDateModel.Date = parsedDate

	fmt.Printf("Binded JSON payload: %+v\n", s.getTaskByEmailAndDateModel)

	return nil
}