package controller

import (
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/model"
	"github.com/xbklyn/getgoal-app/service"
)

type TaskController struct {
	service.TaskService
}

func NewTaskController(s service.TaskService) *TaskController {
	return &TaskController{s}
}

func (controller TaskController) Route(api *gin.RouterGroup) {
	api.GET("/tasks", controller.FindAllTasks)
	api.GET("/tasks/:id", controller.FindTaskByID)
	api.GET("/tasks/to-do", controller.FindTaskFromEmailAndDate)
	api.GET("/tasks/plan/:id", controller.FindTaskFromProgramId)
	api.POST("/tasks", controller.Save)
	api.POST("/tasks/join-program/:program_id", controller.CreateTaskFromProgram)
	api.PUT("/tasks/done/:id", controller.UpdateStatusDone)
	api.PUT("/tasks/un-done/:id", controller.UpdateStatusTodo)
	api.PUT("/tasks/:id", controller.Update)
	api.DELETE("/tasks/:id", controller.Delete)
}

// FindAllTask godoc
// @summary Find All Tasks
// @description Find All Tasks
// @tags Task
// @id FindAllTask
// @accept json
// @produce json
// @response 200 {object} model.GeneralResponse "OK"
// @response 400 {object} model.GeneralResponse "Bad Request"
// @Router /api/v1/tasks [get]
func (controller TaskController) FindAllTasks(c *gin.Context) {
	tasks, err := controller.TaskService.FindAllTasks()
	if err != nil {
		c.JSON(http.StatusBadGateway, model.GeneralResponse{
			Code:    http.StatusBadGateway,
			Message: "Something Went Wrong",
			Data:    nil,
			Error:   err.Error(),
		})
	}
	tasksDTO := make([]model.TaskModel, 0)
	if len(tasks) > 0 {
		tasksDTO = model.ConvertToTaskModels(tasks)
	}
	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Count:   len(tasksDTO),
		Data:    tasksDTO,
		Error:   nil,
	})
}

// FindTaskById godoc
// @summary Find task by ID
// @description Find a task by passing ID
// @tags Task
// @id FindTaskById
// @accept json
// @produce json
// @param id path int true "Task ID"
// @response 200 {object} model.GeneralResponse "OK"
// @response 400 {object} model.GeneralResponse "Bad Request"
// @response 404 {object} model.GeneralResponse "Not Found"
// @Router /api/v1/tasks/:id [get]
func (controller TaskController) FindTaskByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	task, err := controller.TaskService.FindTaskByID(id)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, model.GeneralResponse{
				Code:    http.StatusNotFound,
				Message: "Not Found",
				Data:    nil,
				Error:   "task not found",
			})
			return
		}
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Went Wrong",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	taskDTO := model.ConvertToTaskModel(*task)
	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Count:   1,
		Data:    taskDTO,
		Error:   nil,
	})
}

// SaveTask godoc
// @summary Save Task
// @description Create new task
// @tags Task
// @id SaevTask
// @accept json
// @produce json
// @param task body model.TaskCreateOrUpdate true "Task Data"
// @response 201 {object} model.GeneralResponse "Created"
// @response 400 {object} model.GeneralResponse "Bad Request"
// @Router /api/v1/tasks [post]
func (controller TaskController) Save(c *gin.Context) {
	var task model.TaskCreateOrUpdate
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	taskE, err := controller.TaskService.Save(task, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Went Wrong",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	taskDTO := model.ConvertToTaskModel(*taskE)
	c.JSON(http.StatusCreated, model.GeneralResponse{
		Code:    http.StatusCreated,
		Message: "Success",
		Count:   1,
		Data:    taskDTO,
		Error:   nil,
	})
}

// UpdateTask godoc
// @summary Update Task
// @description Update a task by passing ID
// @tags Task
// @id UpdateTask
// @accept json
// @produce json
// @param id path int true "Task ID"
// @param task body model.TaskCreateOrUpdate true "Task Data"
// @response 200 {object} model.GeneralResponse "OK"
// @response 400 {object} model.GeneralResponse "Bad Request"
// @response 404 {object} model.GeneralResponse "Not Found"
// @Router /api/v1/tasks/:id  [put]
func (controller TaskController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	var task model.TaskCreateOrUpdate
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	taskE, err := controller.TaskService.Update(id, task, c)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, model.GeneralResponse{
				Code:    http.StatusNotFound,
				Message: "Not Found",
				Data:    nil,
				Error:   "task not found",
			})
			return
		}
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Went Wrong",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	taskDTO := model.ConvertToTaskModel(*taskE)
	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Count:   1,
		Data:    taskDTO,
		Error:   nil,
	})
}

// DeleteTask godoc
// @summary Delete Task
// @description Delete Existing Task
// @tags Task
// @id DeleteTask
// @accept json
// @produce json
// @param id path int true "Task ID"
// @response 200 {object} model.GeneralResponse "OK"
// @response 400 {object} model.GeneralResponse "Bad Request"
// @response 404 {object} model.GeneralResponse "Not Found"
// @Router /api/v1/tasks/:id [delete]
func (controller TaskController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	err = controller.TaskService.Delete(id, c)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, model.GeneralResponse{
				Code:    http.StatusNotFound,
				Message: "Not Found",
				Data:    nil,
				Error:   "task not found",
			})
			return
		}
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Went Wrong",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Deleted",
		Data:    nil,
		Error:   nil,
	})
}

// FindTaskByEmailDate godoc
// @summary Find Task by Email and Date
// @description Find Task by Email and Date
// @tags Task
// @id FindTaskByEmailDate
// @accept json
// @produce json
// @param task body model.ToDoRequest true "Data"
// @response 200 {object} model.GeneralResponse "OK"
// @response 400 {object} model.GeneralResponse "Bad Request"
// @response 404 {object} model.GeneralResponse "Not Found"
// @Router /api/v1/tasks/to-do [get]
func (controller TaskController) FindTaskFromEmailAndDate(c *gin.Context) {
	var toDoRequest model.ToDoRequest
	if err := c.ShouldBindJSON(&toDoRequest); err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Body",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	// Regular expression pattern for YYYY-HH-MM format
	pattern := `^\d{4}-\d{2}-\d{2}$`
	match, _ := regexp.MatchString(pattern, toDoRequest.Date)
	if !match {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Date",
			Data:    nil,
			Error:   "date format should be YYYY-MM-DD",
		})
		return
	}
	tasks, err := controller.TaskService.FindTaskByEmailAndDate(toDoRequest, c)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, model.GeneralResponse{
				Code:    http.StatusNotFound,
				Message: "Not Found",
				Data:    nil,
				Error:   "user not found",
			})
			return
		}
		c.JSON(http.StatusBadGateway, model.GeneralResponse{
			Code:    http.StatusBadGateway,
			Message: "Something Went Wrong",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	tasksDTO := make([]model.TaskModel, 0)
	if len(tasks) > 0 {
		tasksDTO = model.ConvertToTaskModels(tasks)
	}
	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Count:   len(tasksDTO),
		Data:    tasksDTO,
		Error:   nil,
	})
}

// plan program  godoc
// @summary list task to plan program
// @description List task related to that program_id for planning
// @tags Task
// @id PlanTask
// @accept json
// @produce json
// @param id path int true "Program ID"
// @response 200 {object} model.GeneralResponse "OK"
// @response 400 {object} model.GeneralResponse "Bad Request"
// @response 404 {object} model.GeneralResponse "Not Found"
// @Router /api/v1/tasks/plan/:program_id [get]
func (controller TaskController) FindTaskFromProgramId(c *gin.Context) {
	programId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	tasks, err := controller.TaskService.GetTaskFromProgramId(programId)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, model.GeneralResponse{
				Code:    http.StatusNotFound,
				Message: "Not Found",
				Data:    nil,
				Error:   "program not found",
			})
			return
		}
		c.JSON(http.StatusBadGateway, model.GeneralResponse{
			Code:    http.StatusBadGateway,
			Message: "Something Went Wrong",
			Data:    nil,
			Error:   err.Error(),
		})
	}

	now := time.Now()
	var originalStartTime time.Time

	for i := range tasks {
		originalTime := tasks[i].StartTime.Format("15:04:05")

		if i > 0 {

			diff := tasks[i].StartTime.Sub(originalStartTime)
			newStartTime := tasks[i-1].StartTime.Add(diff)

			originalStartTime = tasks[i].StartTime
			tasks[i].StartTime = newStartTime
			continue
		} else {

			newStartTime, _ := time.Parse(time.RFC3339, now.Add(time.Hour*24).Format("2006-01-02")+"T"+originalTime+"Z")

			originalStartTime = tasks[i].StartTime
			tasks[i].StartTime = newStartTime
			continue
		}
	}
	tasksDTO := make([]model.TaskModel, 0)
	if len(tasks) > 0 {
		tasksDTO = model.ConvertToTaskModels(tasks)
	}
	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Count:   len(tasksDTO),
		Data:    tasksDTO,
		Error:   nil,
	})
}

// Join program  godoc
// @summary join program
// @description Create tasks from program
// @tags Task
// @id JoinProgram
// @accept json
// @produce json
// @param id path int true "Program ID"
// @param modifications body model.JoinProgramModifications true "Modifications"
// @response 200 {object} model.GeneralResponse "OK"
// @response 400 {object} model.GeneralResponse "Bad Request"
// @response 404 {object} model.GeneralResponse "Not Found"
// @Router /api/v1/tasks/join-program/:program_id [post]
func (controller TaskController) CreateTaskFromProgram(c *gin.Context) {
	programId, err := strconv.ParseUint(c.Param("program_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	var modifications model.JoinProgramModifications
	if err := c.ShouldBindJSON(&modifications); err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Body",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	tasks, err := controller.TaskService.JoinProgram(programId, modifications, c)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, model.GeneralResponse{
				Code:    http.StatusNotFound,
				Message: "Not Found",
				Data:    nil,
				Error:   "program not found",
			})
			return
		}
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Went Wrong",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	tasksDTO := make([]model.TaskModel, 0)
	if len(*tasks) > 0 {
		tasksDTO = model.ConvertToTaskModels(*tasks)
	}
	c.JSON(http.StatusCreated, model.GeneralResponse{
		Code:    http.StatusCreated,
		Message: "Success",
		Count:   len(tasksDTO),
		Data:    tasksDTO,
		Error:   nil,
	})
}

// UpdateDone  godoc
// @summary update status to done
// @description update status to 1 (done)
// @tags Task
// @id UpdateDone
// @accept json
// @produce json
// @param id path int true "Task  ID"
// @response 200 {object} model.GeneralResponse "OK"
// @response 400 {object} model.GeneralResponse "Bad Request"
// @response 404 {object} model.GeneralResponse "Not Found"
// @Router /api/v1/tasks/done/:id [put]
func (controlller TaskController) UpdateStatusDone(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	task, err := controlller.TaskService.UpdateStatus(id, 2)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, model.GeneralResponse{
				Code:    http.StatusNotFound,
				Message: "Not Found",
				Data:    nil,
				Error:   "task not found",
			})
			return
		}
		c.JSON(http.StatusBadGateway, model.GeneralResponse{
			Code:    http.StatusBadGateway,
			Message: "Something Went Wrong",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	taskDTO := model.ConvertToTaskModel(*task)
	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Count:   1,
		Data:    taskDTO,
		Error:   nil,
	})
}

// UpdateTodo  godoc
// @summary update status to todo
// @description update status to 1 (todo)
// @tags Task
// @id UpdateTodo
// @accept json
// @produce json
// @param id path int true "Task  ID"
// @response 200 {object} model.GeneralResponse "OK"
// @response 400 {object} model.GeneralResponse "Bad Request"
// @response 404 {object} model.GeneralResponse "Not Found"
// @Router /api/v1/tasks/un-done/:id [put]
func (controlller TaskController) UpdateStatusTodo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	task, err := controlller.TaskService.UpdateStatus(id, 1)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, model.GeneralResponse{
				Code:    http.StatusNotFound,
				Message: "Not Found",
				Data:    nil,
				Error:   "task not found",
			})
			return
		}
		c.JSON(http.StatusBadGateway, model.GeneralResponse{
			Code:    http.StatusBadGateway,
			Message: "Something Went Wrong",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	taskDTO := model.ConvertToTaskModel(*task)
	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Count:   1,
		Data:    taskDTO,
		Error:   nil,
	})
}
