package userDomain

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

const apiUrl = "users/"

var userBreaker = breaker.New(3, 1, 2*time.Minute)

func (u *User) Get() *errors.ApiError {

	if u.ID <= 0 {
		return &errors.ApiError{
			Message: "ID invalido",
			Status:  http.StatusBadRequest,
		}
	}
	var apiError *errors.ApiError = nil
	url := fmt.Sprintf("%s%s%d", externalApis.BaseUrl, apiUrl, u.ID)
	result := userBreaker.Run(func() error {
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
		err = json.Unmarshal(data, u)
		if err != nil {
			return err
		}
		// communicate with some external service and
		// return an error if the communication failed
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
			Status:  http.StatusServiceUnavailable,
		}
	}
}
