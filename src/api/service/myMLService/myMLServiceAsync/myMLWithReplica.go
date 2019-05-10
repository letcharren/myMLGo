package myMLServiceAsync

import (
	"github.com/mercadolibre/myML/src/api/domain/myMLDomain"
	"github.com/mercadolibre/myML/src/api/utils/errors"
	"sync"
	"net/http"
)

type search func(query string) *myMLDomain.Result

type searchUser func(query int64) *myMLDomain.Result

func FindMLInformationReplica(id int64) (*myMLDomain.MyML, *errors.ApiError){
	if id<=0 {
		return nil, &errors.ApiError{
			Message:"ID invalido",
			Status: http.StatusBadRequest,
		}
	}
	c := make(chan *myMLDomain.Result,3)
	defer close(c)
	c <- firstUser(id,loadUser,loadUser,loadUser,loadUser,loadUser,loadUser,loadUser,loadUser,loadUser,loadUser)
	var result *myMLDomain.Result
	result = <- c
	if result.Error != nil {
		return nil, &errors.ApiError{
			Message: result.Error.Message,
			Status: result.Error.Status,
		}
	}
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
	go func() {c <- first(result.Data.User.CountryID,loadCountry,loadCountry,loadCountry,loadCountry,loadCountry)} ()
	go func() {c <- first(result.Data.User.SiteID,loadSite,loadSite,loadSite,loadSite,loadSite,loadSite)} ()
	wg.Wait()
	return result.Data,result.Error
}


func firstUser(idUser int64, replicas ...searchUser) *myMLDomain.Result {

	c := make(chan *myMLDomain.Result)
	searchReplica := func(i int) {
		c <- replicas[i](idUser)
	}
	for i := range replicas {
		go searchReplica(i)
	}
	return <- c
}

func first(id string, replicas ...search) *myMLDomain.Result {

	c := make(chan *myMLDomain.Result)
	searchReplica := func(i int) {
		c <- replicas[i](id)
	}
	for i := range replicas {
		go searchReplica(i)
	}
	return <- c
}
