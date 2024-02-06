package serializers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/entities"
)

type TaskSerializer struct {
	C *gin.Context
	entities.Task
}

type TaskResponse struct {
	TaskID            uint64     `json:"task_id"`
	TaskName          string     `json:"task_name"`
	TaskStatus        int        `json:"task_status"`
	IsSetNotification int        `json:"is_set_noti"`
	StartTime         time.Time  `json:"start_time"`
	EndTime           *time.Time `json:"end_time"`
	Category          string     `json:"category"`
	TimeBeforeNotify  int        `json:"time_before_notify"`
	TaskDescription   string     `json:"task_description"`
	Link              string     `json:"link"`
	MediaURL          string     `json:"media_url"`
	UpdatedAt         time.Time  `json:"updated_at"`

	Program     *entities.Program    `json:"program"`
	UserAccount entities.UserAccount `json:"owner"`
}

func (s *TaskSerializer) Response() TaskResponse {
	response := TaskResponse{
		TaskID:            s.TaskID,
		TaskName:          s.TaskName,
		TaskStatus:        s.TaskStatus,
		IsSetNotification: s.IsSetNotification,
		StartTime:         s.StartTime,
		EndTime:           s.EndTime,
		Category:          s.Category,
		TimeBeforeNotify:  s.TimeBeforeNotify,
		TaskDescription:   s.TaskDescription,
		Link:              s.Link,
		MediaURL:          s.MediaURL,
		UpdatedAt:         s.UpdatedAt,
		UserAccount:       s.UserAccount,
	}

	// Check if the Program is not nil before dereferencing
	if s.Program != nil {
		response.Program = s.Program
	}

	return response
}

type TasksSerializer struct {
	C     *gin.Context
	Tasks []entities.Task
	Count int
}

func (s *TasksSerializer) Response() map[string]interface{} {
	response := make(map[string]interface{})
	tasksResponses := []TaskResponse{}

	for _, task := range s.Tasks {
		serializer := TaskSerializer{s.C, task}
		tasksResponses = append(tasksResponses, serializer.Response())
	}

	response["count"] = s.Count
	response["tasks"] = tasksResponses

	return response
}

type TasksPlanningSerializer struct {
	C     *gin.Context
	Tasks []entities.Task
	Count int `json:"count"`
}

type TasksPlanningResponse struct {
	TaskID            uint64    `json:"task_id"`
	TaskName          string    `json:"task_name"`
	IsSetNotification int       `json:"is_set_noti"`
	StartTime         time.Time `json:"start_time"`
	TimeBeforeNotify  int       `json:"time_before_notify"`
	TaskDescription   string    `json:"task_description"`
}

func (s *TasksPlanningSerializer) Response() map[string]interface{} {
	response := make(map[string]interface{})
	taskResponses := []TasksPlanningResponse{}

	for _, task := range s.Tasks {
		taskResponses = append(taskResponses, TasksPlanningResponse{
			TaskID:            task.TaskID,
			TaskName:          task.TaskName,
			IsSetNotification: task.IsSetNotification,
			StartTime:         task.StartTime,
			TimeBeforeNotify:  task.TimeBeforeNotify,
			TaskDescription:   task.TaskDescription,
		})
	}

	response["tasks"] = taskResponses
	response["count"] = s.Count

	return response
}
