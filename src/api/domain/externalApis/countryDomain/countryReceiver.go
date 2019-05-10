package countryDomain

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/myML/src/api/domain/externalApis"
	"github.com/mercadolibre/myML/src/api/utils/errors"
	"io/ioutil"
	"net/http"
)

const url  = "https://api.mercadolibre.com/countries/"
//const url  = "http://localhost:8081/countries/"

func (c *Country) Get() *errors.ApiError {

	if c.ID=="" {
		return &errors.ApiError{
			Message:"ID invalido",
			Status: http.StatusBadRequest,
		}
	}
	url:= fmt.Sprintf("%s%s%s",externalApis.BaseUrl,"country/" , c.ID)
	response, err := http.Get(url)
	if err!=nil{
		return &errors.ApiError{
			Message:err.Error(),
			Status: http.StatusBadRequest,
		}
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err!=nil{
		return &errors.ApiError{
			Message:err.Error(),
			Status: http.StatusBadRequest,
		}
	}
	err = json.Unmarshal(data,c)
	if err!=nil{
		return &errors.ApiError{
			Message:err.Error(),
			Status: http.StatusBadRequest,
		}
	}

	return nil
}

