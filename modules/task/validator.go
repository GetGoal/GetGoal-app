package task

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type GetTaskByEmailAndDateValidator struct {
	Email string `json:"email" binding:"required"`
	Date  string `json:"date" binding:"required"`

	getTaskByEmailAndDateModel GetTaskByEmailAndDateTask
}

func NewGetTaskByEmailandDateValidator() GetTaskByEmailAndDateValidator {
	return GetTaskByEmailAndDateValidator{}
}

func (s *GetTaskByEmailAndDateValidator) Bind(c *gin.Context) error {

	err := c.BindJSON(s)
	if err != nil {
		log.Default().Printf("Error binding JSON: %s\n", err.Error())
		return err
	}

	parsedDate, err := time.Parse(time.RFC3339, s.Date)
	if err != nil {
		fmt.Printf("Error parsing date: %s\n", err.Error())
		return err
	}

	s.getTaskByEmailAndDateModel.Email = s.Email
	s.getTaskByEmailAndDateModel.Date = parsedDate

	fmt.Printf("Binded JSON payload: %+v\n", s.getTaskByEmailAndDateModel)

	return nil
}
