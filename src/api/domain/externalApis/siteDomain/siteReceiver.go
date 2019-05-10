package siteDomain

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/myML/src/api/domain/externalApis"
	"github.com/mercadolibre/myML/src/api/utils/errors"
	"io/ioutil"
	"net/http"
)

const url  = "https://api.mercadolibre.com/sites/"
//const url  = "http://localhost:8081/sites/"

func (s *Site) Get() *errors.ApiError {

	if s.ID=="" {
		return &errors.ApiError{
			Message:"ID invalido",
			Status: http.StatusBadRequest,
		}
	}
	url:= fmt.Sprintf("%s%s%s",externalApis.BaseUrl,"sites/" , s.ID)
	response, err := http.Get(url)
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
	err = json.Unmarshal(data,s)
	if err!=nil{
		return &errors.ApiError{
			Message:err.Error(),
			Status: http.StatusBadRequest,
		}
	}
	return nil
}

