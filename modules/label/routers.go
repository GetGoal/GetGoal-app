package label

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/common"
)

func LabelRegister(router *gin.RouterGroup) {
	router.POST("/", LabelCreate)
}
func LabelAnonymousRegister(router *gin.RouterGroup) {
	router.GET("/", LabelList)
	router.GET("/filter", LabelFilterList)
	router.GET("/:id", LabelDetail)
}

func LabelCreate(c *gin.Context) {
	var label Label
	c.BindJSON(&label)

	if err := SaveOne(&label); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	c.JSON(http.StatusOK, label)
}

func LabelList(c *gin.Context) {

	label, err := FindAllLabel()
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("Label", err))
		return
	}

	serializer := LabelsSerializer{c, label}
	c.JSON(http.StatusOK, gin.H{"Lable": serializer.Response()})
}

func LabelFilterList(c *gin.Context) {

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
