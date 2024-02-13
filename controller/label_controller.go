package controller

import (
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
}

func (controller LabelController) FindAllLabels(c *gin.Context) {
	labels, err := controller.LabelService.FindAllLabels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.GeneralResponse{
			Code:    http.StatusInternalServerError,
			Message: "Something Went Wrong",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	labelsDTO := model.ConvertToLabelModels(labels)
	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Count:   len(labels),
		Data:    labelsDTO,
		Error:   nil,
	})
}

func (controller LabelController) FindLabelByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	label, err := controller.LabelService.FindLabelByID(uint64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.GeneralResponse{
			Code:    http.StatusInternalServerError,
			Message: "Something Went Wrong",
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

func (controller LabelController) GetSearchLabel(c *gin.Context) {
	labels, err := controller.LabelService.GetSearchLabel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.GeneralResponse{
			Code:    http.StatusInternalServerError,
			Message: "Something Went Wrong",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	labelsDTO := model.ConvertToLabelModels(labels)
	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Count:   len(labels),
		Data:    labelsDTO,
		Error:   nil,
	})
}

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
		c.JSON(http.StatusInternalServerError, model.GeneralResponse{
			Code:    http.StatusInternalServerError,
			Message: "Something Went Wrong",
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
