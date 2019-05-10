package myMLServiceAsync

import (
	"github.com/mercadolibre/myML/src/api/domain/externalApis/countryDomain"
	"github.com/mercadolibre/myML/src/api/domain/externalApis/siteDomain"
	"github.com/mercadolibre/myML/src/api/domain/externalApis/userDomain"
	"github.com/mercadolibre/myML/src/api/domain/myMLDomain"
)

const RoutinesCount  = 3

func loadUser(userId int64) *myMLDomain.Result{

	var user *userDomain.User= &userDomain.User{
		ID:userId,
	}
	err:= user.Get()
	if err != nil {
		return &(myMLDomain.Result{nil,err})
	}
	myML:= &myMLDomain.MyML{User:user}
	return &(myMLDomain.Result{myML,nil})
}

func loadCountry(countryID string) *myMLDomain.Result{

	var country *countryDomain.Country= &countryDomain.Country{
		ID:countryID,
	}
	err:= country.Get()
	if err != nil {
		return &(myMLDomain.Result{nil,err})
	}
	myML:= &myMLDomain.MyML{Country:country}
	return &(myMLDomain.Result{myML,nil})
}

func loadSite(siteID string) *myMLDomain.Result{

	var site *siteDomain.Site= &siteDomain.Site{
		ID:siteID,
	}
	err:= site.Get()
	if err != nil {
		return &(myMLDomain.Result{nil,err})
	}
	myML:= &myMLDomain.MyML{Site:site}
	return &(myMLDomain.Result{myML,nil})
}

