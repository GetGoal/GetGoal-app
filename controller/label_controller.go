package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/xbklyn/getgoal-app/common"
	"github.com/xbklyn/getgoal-app/model"
	"github.com/xbklyn/getgoal-app/service"
)

type LabelController struct {
	service.LabelService
}

func NewLabelController(labelService service.LabelService) *LabelController {
	return &LabelController{LabelService: labelService}
}

func (controller LabelController) Route(api *gin.RouterGroup) {
	api.GET("/labels", controller.FindAllLabels)
	api.GET("/labels/search", controller.GetSearchLabel)
	api.GET("/labels/:id", controller.FindLabelByID)
	api.POST("/labels", controller.Save)
	api.PUT("/labels/:id", controller.Update)
	api.DELETE("/labels/:id", controller.Delete)
}

// FindAllLable godoc
// @summary Find All Label
// @description Find All Label
// @tags Label
// @id FindAllLabel
// @accept json
// @produce json
// @response 200 {object} model.GeneralResponse "OK"
// @response 400 {object} model.GeneralResponse "Bad Request"
// @Router /api/v1/labels [get]
func (controller LabelController) FindAllLabels(c *gin.Context) {
	labels, err := controller.LabelService.FindAllLabels()
	if err != nil {
		log.Default().Printf("Error: %v", err)
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Went Wrong",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	labelsDTO := make([]model.LabelModel, 0)
	if len(labels) > 0 {
		labelsDTO = model.ConvertToLabelModels(labels)
	}
	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Count:   len(labels),
		Data:    labelsDTO,
		Error:   nil,
	})
}

// FindLabelById godoc
// @summary Find Label By ID
// @description Search label by provided ID
// @tags Label
// @id FindLabelByID
// @param id path int true "Label ID"
// @accept json
// @produce json
// @response 200 {object} model.GeneralResponse "OK"
// @response 400 {object} model.GeneralResponse "Bad Request"
// @response 404 {object} model.GeneralResponse "Not Found"
// @Router /api/v1/labels/:id [get]
func (controller LabelController) FindLabelByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Default().Printf("Error: %v", err)
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Id",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	label, err := controller.LabelService.FindLabelByID(uint64(id))
	if err != nil {
		log.Default().Printf("Error: %v", err)
		c.JSON(http.StatusNotFound, model.GeneralResponse{
			Code:    http.StatusNotFound,
			Message: "Not Found",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	labelDTO := model.ConvertToLabelModel(*label)
	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Count:   1,
		Data:    labelDTO,
		Error:   nil,
	})
}

// GetSearchLabel godoc
// @summary Search Label Filter
// @description Get label with random order and limit
// @tags Label
// @id GetSearchLabel
// @accept json
// @produce json
// @response 200 {object} model.GeneralResponse "OK"
// @response 400 {object} model.GeneralResponse "Bad Request"
// @Router /api/v1/labels/search [get]
func (controller LabelController) GetSearchLabel(c *gin.Context) {
	labels, err := controller.LabelService.GetSearchLabel()
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Went Wrong",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	labelsDTO := make([]model.LabelModel, 0)
	if len(labels) > 0 {
		labelsDTO = model.ConvertToLabelModels(labels)
	}
	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Count:   len(labels),
		Data:    labelsDTO,
		Error:   nil,
	})
}

// SaveLabel godoc
// @summary Save Label
// @description Create new label
// @tags Label
// @id SaveLabel
// @accept json
// @param label body model.LabelRequest true "Label Request"
// @produce json
// @response 201 {object} model.GeneralResponse "Created"
// @response 400 {object} model.GeneralResponse "Bad Request"
// @Router /api/v1/labels [post]
func (controller LabelController) Save(c *gin.Context) {
	label := new(model.LabelRequest)
	if err := common.Bind(c, label); err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	labelCreate, err := controller.LabelService.Save(*label)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	labelDTO := model.ConvertToLabelModel(*labelCreate)
	c.JSON(http.StatusCreated, model.GeneralResponse{
		Code:    http.StatusCreated,
		Message: "Success",
		Count:   1,
		Data:    labelDTO,
		Error:   nil,
	})
}

// Update godoc
// @summary Update Label
// @description Update existing label
// @tags Label
// @id UpdateLabel
// @accept json
// @produce json
// @param id path int true "Label ID"
// @param label body model.LabelRequest true "Label Request"
// @response 200 {object} model.GeneralResponse "OK"
// @response 400 {object} model.GeneralResponse "Bad Request"
// @response 404 {object} model.GeneralResponse "Not Found"
// @Router /api/v1/labels/:id [put]
func (controller LabelController) Update(c *gin.Context) {
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
	label := new(model.LabelRequest)
	if err := common.Bind(c, label); err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Requests!",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	labelUpdate, err := controller.LabelService.Update(uint64(id), *label)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, model.GeneralResponse{
				Code:    http.StatusNotFound,
				Message: "Not Found",
				Data:    nil,
				Error:   err.Error(),
			})
			return
		}
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	labelDTO := model.ConvertToLabelModel(*labelUpdate)
	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Count:   1,
		Data:    labelDTO,
		Error:   nil,
	})
}

// Delete godoc
// @summary Delete Label
// @description Delete existing label
// @tags Label
// @id DeleteLabel
// @accept json
// @produce json
// @param id path int true "Label ID"
// @response 200 {object} model.GeneralResponse "OK"
// @response 400 {object} model.GeneralResponse "Bad Request"
// @Router /api/v1/labels/:id [delete]
func (controller LabelController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Requestsss",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	serviceErr := controller.LabelService.Delete(uint64(id))
	if serviceErr != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
			Data:    nil,
			Error:   serviceErr.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Deleted id " + strconv.FormatUint(uint64(id), 10) + " successfully",
		Data:    nil,
		Error:   nil,
	})
}
