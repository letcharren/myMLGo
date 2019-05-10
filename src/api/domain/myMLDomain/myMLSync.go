package myMLDomain

import (
	"github.com/mercadolibre/myML/src/api/domain/externalApis/siteDomain"
	"github.com/mercadolibre/myML/src/api/domain/externalApis/countryDomain"
	"github.com/mercadolibre/myML/src/api/domain/externalApis/userDomain"
	"github.com/mercadolibre/myML/src/api/utils/errors"
)

type MyML struct {
	User *userDomain.User
	Site *siteDomain.Site
	Country *countryDomain.Country
}

type Result struct{
	Data *MyML
	Error *errors.ApiError
}

