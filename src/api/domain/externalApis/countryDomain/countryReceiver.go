package countryDomain

import (
	"encoding/json"
	errors2 "errors"
	"fmt"
	"github.com/mercadolibre/myML/src/api/domain/externalApis"
	"github.com/mercadolibre/myML/src/api/utils/errors"
	"gopkg.in/eapache/go-resiliency.v1/breaker"
	"io/ioutil"
	"net/http"
)

const apiUrl = "countries/"

var countryBreaker = breaker.New(externalApis.ErrorThreshold, externalApis.SuccessThreshold, externalApis.TimeOpen)

func (c *Country) Get() *errors.ApiError {

	if c.ID == "" {
		return &errors.ApiError{
			Message: "ID invalido",
			Status:  http.StatusBadRequest,
		}
	}
	var apiError *errors.ApiError = nil
	url := fmt.Sprintf("%s%s%s", externalApis.BaseUrl, apiUrl, c.ID)
	//Lamada a Circuit Breaker
	//Se considera que la llamada fallo cuando es un error 5xx o
	//falla la conexion (http.Get(url) retorna error)
	err := countryBreaker.Run(func() error {
		response, err := http.Get(url)
		if response != nil {
			defer response.Body.Close()
		}
		if err != nil {
			return err
		}
		if response.StatusCode != http.StatusOK {
			apiError = &errors.ApiError{
				Message: response.Status,
				Status:  response.StatusCode,
			}
			if response.StatusCode >= http.StatusInternalServerError {
				return errors2.New("Server-Error")
			}
			return nil
		}
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}
		err = json.Unmarshal(data, c)
		if err != nil {
			return err
		}
		return nil
	})
	switch err {
	case nil:
		return apiError
	case breaker.ErrBreakerOpen:
		return &errors.ApiError{
			Message: "circuit Breaker Country Open",
			Status:  http.StatusServiceUnavailable,
		}
	default:
		if err.Error() == "Server-Error" {
			return apiError
		}
		return &errors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
}
