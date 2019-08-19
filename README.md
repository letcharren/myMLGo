# myMLGo

Api myML para practica ejercicio en Go. Realiza distintas consultas a servicio externos para devolver la informacion sobre un usario

## Dependencias
Implementado el uso de los servicios externo con circuit breaker. Para que el proyecto funcione descargar 
https://gopkg.in/eapache/go-resiliency.v1

## Uso
Para hacer uso de la api hacer una peticion GET del siguien endPoint: 
http://localhost:8080/usersync/{user} 
donde {user} es un numero entero.

Para caso de uso con fallo de una de las API externas puede usar el siguiente mock server:
https://github.com/letcharren/mockServer

Para modificar parametros circuit breaker o que se llame a api mercado libre y no a mockServer modificar el archivo:
/src/api/domain/externalApis.go
