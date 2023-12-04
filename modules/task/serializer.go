package task

import (
	"time"

	"github.com/gin-gonic/gin"
)

type TaskSerializer struct {
	C *gin.Context
	Task
}

type TaskResponse struct {
	TaskID            uint64     `json:"task_id"`
	TaskName          string     `json:"task_name"`
	TaskStatus        int        `json:"task_status"`
	UserAccountID     int        `json:"user_account_id"`
	IsSetNotification int        `json:"is_set_noti"`
	StartTime         time.Time  `json:"start_time"`
	EndTime           *time.Time `json:"end_time"`
	ProgramID         int        `json:"program_id"`
	Category          string     `json:"category"`
	TimeBeforeNotify  int        `json:"time_before_notify"`
	TaskDescription   string     `json:"task_description"`
	Link              string     `json:"link"`
	MediaURL          string     `json:"media_url"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

func (s *TaskSerializer) Response() TaskResponse {

	return TaskResponse{
		TaskID:            s.TaskID,
		TaskName:          s.TaskName,
		TaskStatus:        s.TaskStatus,
		UserAccountID:     s.UserAccountID,
		IsSetNotification: s.IsSetNotification,
		StartTime:         s.StartTime,
		EndTime:           s.EndTime,
		ProgramID:         s.ProgramID,
		Category:          s.Category,
		TimeBeforeNotify:  s.TimeBeforeNotify,
		TaskDescription:   s.TaskDescription,
		Link:              s.Link,
		MediaURL:          s.MediaURL,
		UpdatedAt:         s.UpdatedAt,
	}
}

type TasksSerializer struct {
	C     *gin.Context
	Tasks []Task
	Count int `json:"count"`
}

func (s *TasksSerializer) Response() map[string]interface{} {
	response := make(map[string]interface{})
	programResponses := []TaskResponse{}

	for _, task := range s.Tasks {
		serializer := TaskSerializer{s.C, task}
		programResponses = append(programResponses, serializer.Response())
	}

	response["count"] = s.Count
	response["tasks"] = programResponses

	return response
}
