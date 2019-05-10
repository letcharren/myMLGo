package myMLController

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/myML/src/api/service/myMLService/myMLServiceSync"
	"github.com/mercadolibre/myML/src/api/utils/errors"
	"net/http"
	"strconv"
)

func GetMLSync(ctx *gin.Context){

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
	myML, myErr := myMLServiceSync.FindMLInformationSync(idUser)
	if myErr != nil {
		ctx.JSON(myErr.Status, myErr)
		return
	}
	ctx.JSON(200,myML)
}
