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
	"time"
)

const apiUrl = "countries/"

var countryBreaker = breaker.New(3, 1, 2*time.Minute)

func (c *Country) Get() *errors.ApiError {

	if c.ID == "" {
		return &errors.ApiError{
			Message: "ID invalido",
			Status:  http.StatusBadRequest,
		}
	}
	var apiError *errors.ApiError = nil
	url := fmt.Sprintf("%s%s%s", externalApis.BaseUrl, apiUrl, c.ID)
	result := countryBreaker.Run(func() error {
		response, err := http.Get(url)
		if response != nil {
			defer response.Body.Close()
		}
		if err != nil {
			return err
		}
		if response.StatusCode != http.StatusOK {
			if response.StatusCode >= http.StatusInternalServerError {
				return errors2.New(response.Status)
			}
			apiError = &errors.ApiError{
				Message: response.Status,
				Status:  response.StatusCode,
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
	switch result {
	case nil:
		return apiError
	case breaker.ErrBreakerOpen:
		return &errors.ApiError{
			Message: "Abierto circuitBreaker",
			Status:  http.StatusServiceUnavailable,
		}
	default:
		return &errors.ApiError{
			Message: result.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
}
