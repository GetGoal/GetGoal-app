package model

import (
	"time"

	"github.com/xbklyn/getgoal-app/entity"
)

type TaskCreateOrUpdate struct {
	TaskName          string    `json:"task_name" validate:"required,max=250"`
	IsSetNotification int       `json:"is_set_noti" default:"0"`
	StartTime         time.Time `json:"start_time"`
	Category          string    `json:"category"`
	TimeBeforeNotify  int       `json:"time_before_notify"`
	TaskDescription   string    `json:"task_description"`
	Link              string    `json:"link"`
	MediaURL          string    ` json:"media_url"`
	Owner             uint      `json:"owner" binding:"required" validate:"required"`
}

type ToDoRequest struct {
	Email string `json:"email" validate:"required,email"`
	Date  string `json:"date" validate:"required"`
}
type Modification struct {
	IsSetNotification int    `json:"is_set_noti"`
	StartTime         string `json:"start_time" binding:"required"`
	TimeBeforeNotify  int    `json:"time_before_notify"`
}
type JoinProgramModifications struct {
	Email         string         `json:"email" validate:"required,email"`
	Modifications []Modification `json:"modifications"`
}

type TaskModel struct {
	TaskID            uint64    `json:"task_id"`
	TaskName          string    `json:"task_name"`
	IsSetNotification int       `json:"is_set_noti"`
	StartTime         time.Time `json:"start_time"`
	Category          string    `json:"category"`
	TimeBeforeNotify  int       `json:"time_before_notify"`
	TaskDescription   string    `json:"task_description"`
	Link              string    `json:"link"`
	MediaURL          string    `json:"media_url"`
	RelatedProgram    uint      `json:"related_program_id"`
	OwnerID           uint      `json:"owner_id"`
}

func ConvertToTaskModel(entityTask entity.Task) TaskModel {
	if entityTask.ProgramID != nil {
		return TaskModel{
			TaskID:            entityTask.TaskID,
			TaskName:          entityTask.TaskName,
			IsSetNotification: entityTask.IsSetNotification,
			StartTime:         entityTask.StartTime,
			Category:          entityTask.Category,
			TimeBeforeNotify:  entityTask.TimeBeforeNotify,
			TaskDescription:   entityTask.TaskDescription,
			Link:              entityTask.Link,
			MediaURL:          entityTask.MediaURL,
			RelatedProgram:    uint(*entityTask.ProgramID),
			OwnerID:           uint(entityTask.UserAccountID),
		}
	}
	return TaskModel{
		TaskID:            entityTask.TaskID,
		TaskName:          entityTask.TaskName,
		IsSetNotification: entityTask.IsSetNotification,
		StartTime:         entityTask.StartTime,
		Category:          entityTask.Category,
		TimeBeforeNotify:  entityTask.TimeBeforeNotify,
		TaskDescription:   entityTask.TaskDescription,
		Link:              entityTask.Link,
		MediaURL:          entityTask.MediaURL,
		OwnerID:           uint(entityTask.UserAccountID),
	}
}

func ConvertToTaskModels(entityTasks []entity.Task) []TaskModel {
	var tasks []TaskModel
	for _, task := range entityTasks {
		tasks = append(tasks, ConvertToTaskModel(task))
	}
	return tasks
}