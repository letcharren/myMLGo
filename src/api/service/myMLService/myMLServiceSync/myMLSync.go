package myMLServiceSync

import (
	"github.com/mercadolibre/myML/src/api/domain/externalApis/countryDomain"
	"github.com/mercadolibre/myML/src/api/domain/externalApis/siteDomain"
	"github.com/mercadolibre/myML/src/api/domain/externalApis/userDomain"
	"github.com/mercadolibre/myML/src/api/domain/myMLDomain"
	"github.com/mercadolibre/myML/src/api/utils/errors"
	"net/http"
)

func FindMLInformationSync(id int64) (*myMLDomain.MyML, *errors.ApiError){

	if id<=0 {
		return nil, &errors.ApiError{
			Message:"ID invalido",
			Status: http.StatusBadRequest,
		}
	}
	var myML myMLDomain.MyML
	myML.User = &userDomain.User{
		ID:id,
	}
	apiError := myML.User.Get()
	if apiError != nil{
		return nil, &errors.ApiError{
			Message:apiError.Message,
			Status: apiError.Status,
		}
	}
	myML.Country= &countryDomain.Country{
		ID:myML.User.CountryID,
	}
	apiError = myML.Country.Get()
	if apiError != nil{
		return nil, &errors.ApiError{
			Message:apiError.Message,
			Status: apiError.Status,
		}
	}
	myML.Site= &siteDomain.Site{
		ID:myML.User.SiteID,
	}
	apiError = myML.Site.Get()
	if apiError != nil{
		return nil, &errors.ApiError{
			Message:apiError.Message,
			Status: apiError.Status,
		}
	}
	return &myML,nil
}
