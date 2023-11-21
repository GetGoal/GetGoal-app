package program

import (
	"time"

	"github.com/gin-gonic/gin"
)

type ProgramSerializer struct {
	C *gin.Context
	Program
}

type ProgramResponse struct {
	ProgramID          uint64    `json:"program_id"`
	ProgramName        string    `json:"program_name"`
	Rating             float64   `json:"rating"`
	ProgramDescription string    `json:"program_description"`
	ExpectedTime       string    `json:"expected_time"`
	UpdatedAt          time.Time `json:"updated_at"`
	Labels             []Label   `json:"labels"`
}

func (s *ProgramSerializer) Response() ProgramResponse {

	return ProgramResponse{
		ProgramID:          s.ProgramID,
		ProgramName:        s.ProgramName,
		ProgramDescription: s.ProgramDescription,
		Rating:             s.Rating,
		ExpectedTime:       s.ExpectedTime,
		UpdatedAt:          s.UpdatedAt,
		Labels:             s.Labels,
	}
}

type ProgramsSerializer struct {
	C       *gin.Context
	Program []Program
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
	response["program"] = programResponses

	return response
}
