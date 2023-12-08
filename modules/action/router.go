package action

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/common"
)

func ActionTypeRegister(router *gin.RouterGroup) {}

func ActionTypeAnonymousRegister(router *gin.RouterGroup) {
	router.GET("", ActionTypeList)
	router.GET("/:id", ActionTypeDetail)
}

func ActionTypeList(c *gin.Context) {
	actionTypes, err := FindAllActions()
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("ActionType", err))
		return
	}

	serializer := ActionTypesSerializer{C: c, Actions: actionTypes, Count: len(actionTypes)}
	c.JSON(http.StatusOK, gin.H{"ActionType": serializer.Response()})
}

func ActionTypeDetail(c *gin.Context) {
	actionId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewError("ActionType", err))
		return
	}

	actionType, err := FindOneAction(&ActionType{ActionID: actionId})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("ActionType", err))
		return
	}

	serializer := ActionTypeSerializer{C: c, ActionType: actionType}
	c.JSON(http.StatusOK, gin.H{"ActionType": serializer.Response()})

}
