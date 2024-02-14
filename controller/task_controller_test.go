package controller_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/xbklyn/getgoal-app/controller"
	"github.com/xbklyn/getgoal-app/entity"
	"github.com/xbklyn/getgoal-app/model"
)

type mockTaskService struct {
	mock.Mock
}

// JoinProgram implements service.TaskService.
func (m *mockTaskService) JoinProgram(programID uint64, modifications model.JoinProgramModifications) (*[]entity.Task, error) {
	args := m.Called(programID, modifications)
	return args.Get(0).(*[]entity.Task), args.Error(1)
}

func (m *mockTaskService) FindAllTasks() ([]entity.Task, error) {
	// Mock data for testing
	tasks := []entity.Task{
		{
			TaskID:            1,
			TaskName:          "Task 1",
			TaskStatus:        1,
			IsSetNotification: 1,
			StartTime:         time.Now(),
			EndTime:           nil,
			Category:          "Category 1",
			TimeBeforeNotify:  5,
			TaskDescription:   "Task 1 Description",
			Link:              "http://example.com",
			MediaURL:          "http://media.example.com",
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
			DeletedAt:         nil,
			ProgramID:         nil,
			Program:           nil,
			UserAccountID:     1,
			UserAccount:       entity.UserAccount{},
		},
	}
	return tasks, nil
}

func (m *mockTaskService) FindTaskByID(id uint64) (*entity.Task, error) {
	// Mock data for testing
	if id == 1 {
		task := entity.Task{
			TaskID:            id,
			TaskName:          "Task 1",
			TaskStatus:        1,
			IsSetNotification: 1,
			StartTime:         time.Now(),
			EndTime:           nil,
			Category:          "Category 1",
			TimeBeforeNotify:  5,
			TaskDescription:   "Task 1 Description",
			Link:              "http://example.com",
			MediaURL:          "http://media.example.com",
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
			DeletedAt:         nil,
			ProgramID:         nil,
			Program:           nil,
			UserAccountID:     1,
			UserAccount:       entity.UserAccount{},
		}
		return &task, nil
	}
	return nil, errors.New("record not found")
}

func (m *mockTaskService) FindTaskByEmailAndDate(request model.ToDoRequest) ([]entity.Task, error) {
	// Mock data for testing
	if request.Email == "test@example.com" && request.Date == "2024-02-14" {
		tasks := []entity.Task{
			{
				TaskID:            1,
				TaskName:          "Task 1",
				TaskStatus:        1,
				IsSetNotification: 1,
				StartTime:         time.Now(),
				EndTime:           nil,
				Category:          "Category 1",
				TimeBeforeNotify:  5,
				TaskDescription:   "Task 1 Description",
				Link:              "http://example.com",
				MediaURL:          "http://media.example.com",
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
				DeletedAt:         nil,
				ProgramID:         nil,
				Program:           nil,
				UserAccountID:     1,
				UserAccount:       entity.UserAccount{},
			},
		}
		return tasks, nil
	}
	return nil, errors.New("no tasks found for the given email and date")
}

func (m *mockTaskService) GetTaskFromProgramId(programId uint64) ([]entity.Task, error) {
	// Mock data for testing
	if programId == 1 {
		tasks := []entity.Task{
			{
				TaskID:            1,
				TaskName:          "Task 1",
				TaskStatus:        1,
				IsSetNotification: 1,
				StartTime:         time.Now(),
				EndTime:           nil,
				Category:          "Category 1",
				TimeBeforeNotify:  5,
				TaskDescription:   "Task 1 Description",
				Link:              "http://example.com",
				MediaURL:          "http://media.example.com",
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
				DeletedAt:         nil,
				ProgramID:         nil,
				Program:           nil,
				UserAccountID:     1,
				UserAccount:       entity.UserAccount{},
			},
		}
		return tasks, nil
	}
	return nil, errors.New("no tasks found for the given program ID")
}

func (m *mockTaskService) Save(task model.TaskCreateOrUpdate) (*entity.Task, error) {
	// Mock data for testing
	if task.TaskName == "" {
		return nil, errors.New("task name is required")
	}
	newTask := entity.Task{
		TaskID:            1,
		TaskName:          task.TaskName,
		TaskStatus:        1,
		IsSetNotification: task.IsSetNotification,
		StartTime:         task.StartTime,
		EndTime:           nil,
		Category:          task.Category,
		TimeBeforeNotify:  task.TimeBeforeNotify,
		TaskDescription:   task.TaskDescription,
		Link:              task.Link,
		MediaURL:          task.MediaURL,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		DeletedAt:         nil,
		ProgramID:         nil,
		Program:           nil,
		UserAccountID:     1,
		UserAccount:       entity.UserAccount{},
	}
	return &newTask, nil
}

func (m *mockTaskService) Update(id uint64, task model.TaskCreateOrUpdate) (*entity.Task, error) {
	// Mock data for testing
	if id == 1 {
		updatedTask := entity.Task{
			TaskID:            id,
			TaskName:          task.TaskName,
			TaskStatus:        1,
			IsSetNotification: task.IsSetNotification,
			StartTime:         task.StartTime,
			EndTime:           nil,
			Category:          task.Category,
			TimeBeforeNotify:  task.TimeBeforeNotify,
			TaskDescription:   task.TaskDescription,
			Link:              task.Link,
			MediaURL:          task.MediaURL,
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
			DeletedAt:         nil,
			ProgramID:         nil,
			Program:           nil,
			UserAccountID:     1,
			UserAccount:       entity.UserAccount{},
		}
		return &updatedTask, nil
	}
	return nil, errors.New("task not found")
}

func (m *mockTaskService) Delete(id uint64) error {
	// Mock data for testing
	if id == 1 {
		// Return nil indicating success
		return nil
	}
	return errors.New("task not found")
}

func (m *mockTaskService) UpdateStatus(id uint64, status int) (*entity.Task, error) {
	// Mock data for testing
	if id == 1 {
		task := entity.Task{}
		return &task, nil
	}
	return nil, errors.New("task not found")

}

var (
	_              *gin.Engine
	taskService    *mockTaskService
	taskController *controller.TaskController
)

func setupTaskController() {
	gin.SetMode(gin.TestMode)
	r = gin.New()
	taskService = &mockTaskService{}
	taskController = controller.NewTaskController(taskService)
	api := r.Group("/api/v1")
	taskController.Route(api)
}

func TestFindAllTasks_Success(t *testing.T) {
	setupTaskController()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/tasks", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response model.GeneralResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "Success", response.Message)
	// Add more assertions based on your response structure
}

func TestFindTaskByID_Success(t *testing.T) {
	setupTaskController()

	id := uint64(1)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/tasks/"+strconv.FormatUint(id, 10), nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response model.GeneralResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "Success", response.Message)
	// Add more assertions based on your response structure
}

func TestFindTaskByID_InvalidID(t *testing.T) {
	setupTaskController()

	id := "invalidID"
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/tasks/"+id, nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response model.GeneralResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, response.Code)
	assert.Contains(t, response.Message, "Invalid ID")
}

func TestFindTaskByID_TaskNotFound(t *testing.T) {
	setupTaskController()

	id := uint64(1000)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/tasks/"+strconv.FormatUint(id, 10), nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var response model.GeneralResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, response.Code)
	assert.Contains(t, response.Error, "task not found")
}

func TestSaveTask_Success(t *testing.T) {
	setupTaskController()

	payload := `{"task_name":"New Task", "description":"Test Description" owner_id:1}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/tasks", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response model.GeneralResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Equal(t, "Success", response.Message)
}

func TestSaveTask_InvalidPayload(t *testing.T) {
	setupTaskController()

	payload := `{"invalid_field":"New Task"}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/tasks", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response model.GeneralResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, response.Code)
	assert.Contains(t, response.Message, "Something Went Wrong")
	assert.Contains(t, response.Error, "task name is required")
}

func TestUpdateTask_Success(t *testing.T) {
	setupTaskController()

	id := uint64(1)
	payload := `{"task_name":"Updated Task", "description":"Updated Description"}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/v1/tasks/"+strconv.FormatUint(id, 10), strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response model.GeneralResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "Success", response.Message)
}

func TestUpdateTask_InvalidID(t *testing.T) {
	setupTaskController()

	id := "invalidID"
	payload := `{"task_name":"Updated Task", "description":"Updated Description"}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/v1/tasks/"+id, strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response model.GeneralResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, response.Code)
	assert.Contains(t, response.Message, "Invalid ID")
}

func TestDeleteTask_Success(t *testing.T) {
	setupTaskController()

	id := uint64(1)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/v1/tasks/"+strconv.FormatUint(id, 10), nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response model.GeneralResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "Deleted", response.Message)
}

func TestDeleteTask_InvalidID(t *testing.T) {
	setupTaskController()

	id := "invalidID"
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/v1/tasks/"+id, nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response model.GeneralResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, response.Code)
	assert.Contains(t, response.Message, "Invalid ID")
}

func TestJoinProgram_Success(t *testing.T) {
	// Set up test environment
	mockTaskService := &mockTaskService{}
	controller := controller.TaskController{TaskService: mockTaskService}
	programID := uint64(123)
	modifications := model.JoinProgramModifications{
		Email: "test@example.com",
		Modifications: []model.Modification{
			{
				IsSetNotification: 1,
				StartTime:         "2024-02-14T12:00:00Z",
				TimeBeforeNotify:  30,
			},
		},
	}
	expectedTasks := &[]entity.Task{}

	// Set up expected behavior of the mock
	mockTaskService.On("JoinProgram", programID, modifications).Return(expectedTasks, nil)

	// Make the actual function call
	tasks, err := controller.TaskService.JoinProgram(programID, modifications)

	// Check if the function behaves as expected
	assert.Equal(t, expectedTasks, tasks)
	assert.NoError(t, err)
}

func TestJoinProgram_ValidInput(t *testing.T) {
	// Set up test environment
	mockTaskService := &mockTaskService{}
	controller := controller.TaskController{TaskService: mockTaskService}

	// Define test case
	programID := uint64(123)
	modifications := model.JoinProgramModifications{
		Email: "test@example.com",
		Modifications: []model.Modification{
			{
				IsSetNotification: 1,
				StartTime:         "2024-02-14T12:00:00Z",
				TimeBeforeNotify:  30,
			},
		},
	}
	expectedTasks := &[]entity.Task{}

	// Set up expected behavior of the mock
	mockTaskService.On("JoinProgram", programID, modifications).Return(expectedTasks, nil)

	// Make the actual function call
	tasks, err := controller.TaskService.JoinProgram(programID, modifications)

	// Check if the function behaves as expected
	assert.Equal(t, expectedTasks, tasks)
	assert.Equal(t, nil, err)
}

func TestJoinProgram_InvalidProgramID(t *testing.T) {
	// Set up test environment
	mockTaskService := &mockTaskService{}
	controller := controller.TaskController{TaskService: mockTaskService}

	// Define test case
	programID := uint64(0) // Invalid program ID
	modifications := model.JoinProgramModifications{
		Email: "test@example.com",
		Modifications: []model.Modification{
			{
				IsSetNotification: 1,
				StartTime:         "2024-02-14T12:00:00Z",
				TimeBeforeNotify:  30,
			},
		},
	}
	expectedError := errors.New("invalid program ID")

	// Set up expected behavior of the mock
	mockTaskService.On("JoinProgram", programID, modifications).Return(&[]entity.Task{}, expectedError)

	// Make the actual function call
	tasks, err := controller.TaskService.JoinProgram(programID, modifications)

	// Check if the function behaves as expected
	assert.Equal(t, &[]entity.Task{}, tasks)
	assert.Equal(t, expectedError, err)
}
