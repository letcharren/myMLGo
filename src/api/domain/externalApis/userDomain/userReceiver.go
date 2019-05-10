package userDomain

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/myML/src/api/utils/errors"
	"github.com/mercadolibre/myML/src/api/domain/externalApis"
	"io/ioutil"
	"net/http"
)

const url  = "https://api.mercadolibre.com/users/"
//const url  = "http://localhost:8081/users/"
func (u *User) Get() *errors.ApiError {

	if u.ID<=0 {
		return &errors.ApiError{
			Message:"ID invalido",
			Status: http.StatusBadRequest,
		}
	}

	url:= fmt.Sprintf("%s%s%d",externalApis.BaseUrl ,"users/",  u.ID)
	response, err := http.Get(url)
	//falta manejar error
	defer response.Body.Close()
	if err!=nil{
		return &errors.ApiError{
			Message:err.Error(),
			Status: http.StatusBadRequest,
		}
	}
	data, err := ioutil.ReadAll(response.Body)
	if err!=nil{
		return &errors.ApiError{
			Message:err.Error(),
			Status: http.StatusBadRequest,
		}
	}
	err = json.Unmarshal(data,u)
	if err!=nil{
		return &errors.ApiError{
			Message:err.Error(),
			Status: http.StatusBadRequest,
		}
	}
	return nil
}
