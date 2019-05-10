package myMLController

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/myML/src/api/service/myMLService/myMLServiceAsync"
	"github.com/mercadolibre/myML/src/api/utils/errors"
	"net/http"
	"strconv"
)
func GetMLChainReplica(ctx *gin.Context){

	idUser,err := strconv.ParseInt(ctx.Param(pUserId), 10, 64)
	if err!=nil {
		apiErr := errors.ApiError{
			Message : err.Error(),
			Status: http.StatusBadRequest,
		}
		ctx.JSON(apiErr.Status,apiErr)
		return
	}
	if idUser<=0{
		apiErr := errors.ApiError{
			Message : "Error, ID invalido",
			Status: http.StatusBadRequest,
		}
		ctx.JSON(apiErr.Status,apiErr)
		return
	}
	myML, myMLErr := myMLServiceAsync.FindMLInformationReplica(idUser)
	if myMLErr != nil {
		ctx.JSON(myMLErr.Status, myMLErr)
		return
	}
	ctx.JSON(200,myML)
}

