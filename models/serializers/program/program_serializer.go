package serializers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/entities"
)

type ProgramSerializer struct {
	C *gin.Context
	entities.Program
}

type ProgramResponse struct {
	ProgramID          uint64    `json:"program_id"`
	ProgramName        string    `json:"program_name"`
	Rating             float64   `json:"rating"`
	ProgramDescription string    `json:"program_description"`
	MediaURL           string    `json:"media_url"`
	ExpectedTime       string    `json:"expected_time"`
	UpdatedAt          time.Time `json:"updated_at"`
	Labels             []LabelP  `json:"labels"`
	Tasks              []TaskP   `json:"tasks"`
}

type LabelP struct {
	LabelID   uint64 `json:"label_id"`
	LabelName string `json:"label_name"`
}

type TaskP struct {
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
}

func (s *ProgramSerializer) Response() ProgramResponse {
	var labels []LabelP
	for _, label := range s.Labels {
		labels = append(labels, LabelP{
			LabelName: label.LabelName,
			LabelID:   label.LabelID,
		})
	}

	var tasks []TaskP
	for _, task := range s.Tasks {
		tasks = append(tasks, TaskP{
			TaskID:            task.TaskID,
			TaskName:          task.TaskName,
			TaskStatus:        task.TaskStatus,
			IsSetNotification: task.IsSetNotification,
			StartTime:         task.StartTime,
			EndTime:           task.EndTime,
			Category:          task.Category,
			TimeBeforeNotify:  task.TimeBeforeNotify,
			TaskDescription:   task.TaskDescription,
			Link:              task.Link,
			MediaURL:          task.MediaURL,
			UpdatedAt:         task.UpdatedAt,
		})
	}

	return ProgramResponse{
		ProgramID:          s.ProgramID,
		ProgramName:        s.ProgramName,
		ProgramDescription: s.ProgramDescription,
		MediaURL:           s.MediaURL,
		Rating:             s.Rating,
		ExpectedTime:       s.ExpectedTime,
		UpdatedAt:          s.UpdatedAt,
		Labels:             labels,
		Tasks:              tasks,
	}
}

type ProgramsSerializer struct {
	C       *gin.Context
	Program []entities.Program
	Count   int `json:"count"`
}

func (s *ProgramsSerializer) Response() map[string]interface{} {
	response := make(map[string]interface{})
	programResponses := []ProgramResponse{}

	for _, program := range s.Program {
		serializer := ProgramSerializer{s.C, program}
		programResponses = append(programResponses, serializer.Response())
	}

	response["count"] = s.Count
	response["programs"] = programResponses

	return response
}
