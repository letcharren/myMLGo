package externalApis

import "time"

/*
*url base para consultas a api externas
 */
//const BaseUrl  = "https://api.mercadolibre.com/"
const BaseUrl = "http://localhost:8081/"

/**
*Constantes para circuit breaker
*ErrorThreshold = Cantidad de errores para poner Closed circuit breaker
*SuccessThreshold = Cantidad de solicitudes acepdtadas para pasar de half Open a Open
**/
const ErrorThreshold = 3
const SuccessThreshold = 1
const TimeOpen = 2 * time.Minute
