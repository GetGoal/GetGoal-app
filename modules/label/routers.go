package label

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/common"
)

func LabelRegister(router *gin.RouterGroup) {
	router.POST("", LabelCreate)
}
func LabelAnonymousRegister(router *gin.RouterGroup) {
	router.GET("", LabelList)
	router.GET("/search", LabelFilterList)
	router.GET("/:id", LabelDetail)
}

func LabelCreate(c *gin.Context) {
	labelValidator := NewLabelValidator()
	if err := labelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	if err := SaveOne(&labelValidator.labelModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Database", err))
		return
	}

	serializer := LabelSerializer{c, labelValidator.labelModel}
	c.JSON(http.StatusOK, gin.H{"Label": serializer.Response()})
}

func LabelList(c *gin.Context) {

	label, err := FindAllLabel()
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("Label", err))
		return
	}

	serializer := LabelsSerializer{C: c, Labels: label, Count: len(label)}
	c.JSON(http.StatusOK, gin.H{"Lable": serializer.Response()})
}

func LabelFilterList(c *gin.Context) {
	label, err := FindSearchLabel()
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewError("Label", err))
		return
	}

	serializer := LabelsSerializer{C: c, Labels: label, Count: len(label)}
	c.JSON(http.StatusOK, gin.H{"Lable": serializer.Response()})
}

func LabelDetail(c *gin.Context) {
	labelId, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadGateway, common.NewError("Label", err))
		return
	}

	label, err := FindOneLable(&Label{LabelID: labelId})

	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("Label", err))
		return
	}

	serializer := LabelSerializer{c, label}
	c.JSON(http.StatusOK, gin.H{"Lable": serializer.Response()})
}
