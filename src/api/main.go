package main
import (
	ginTonic "github.com/gin-gonic/gin"
	myML "github.com/mercadolibre/myML/src/api/controller/myMLController"
	ping "github.com/mercadolibre/myML/src/api/controller/pingController"
)

const(
	port = ":8080"
)

var(
	router = ginTonic.Default()
)

func main(){

	router.GET("/user/:id", myML.GetMLChain)
	router.GET("/userreplica/:id", myML.GetMLChainReplica)
	router.GET("/usersync/:id", myML.GetMLSync)
	router.GET("/ping", ping.Ping)
	err := router.Run(port)
	if err!=nil{
		panic("error al inciar puerto")
	}
}
