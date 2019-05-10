package myMLServiceAsync

import (
	"github.com/mercadolibre/myML/src/api/domain/externalApis/userDomain"
	"github.com/mercadolibre/myML/src/api/domain/myMLDomain"
	"github.com/mercadolibre/myML/src/api/utils/errors"
	"net/http"
	"sync"
)

func FindMLInformationAsync(id int64) (*myMLDomain.MyML, *errors.ApiError){
	if id<=0 {
		return nil, &errors.ApiError{
			Message:"ID invalido",
			Status: http.StatusBadRequest,
		}
	}
	var result myMLDomain.Result
	result.Data = &myMLDomain.MyML{}
	result.Data.User = &userDomain.User{
		ID:id,
	}
	apiError := result.Data.User.Get()
	if apiError != nil{
		return nil, &errors.ApiError{
			Message:apiError.Message,
			Status: apiError.Status,
		}
	}
	c := make(chan *myMLDomain.Result,2)
	defer close(c)
	var wg sync.WaitGroup
	wg.Add(RoutinesCount)
	go func() {
		defer wg.Done()
		for i := 0; i < RoutinesCount-1 ; i++ {
			currentResult := <-c
			wg.Done()
			if currentResult.Error!=nil{
				return
			}
			if currentResult.Data.Site != nil {
				result.Data.Site = currentResult.Data.Site
				continue
			}

			if currentResult.Data.Country != nil {
				result.Data.Country = currentResult.Data.Country
				continue
			}
		}
	}()
	go func() {c <- loadCountry(result.Data.User.CountryID)} ()
	go func() {c <- loadSite(result.Data.User.SiteID)} ()
	wg.Wait()
	return result.Data,result.Error
}

