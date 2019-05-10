package pingController
import ginTonic "github.com/gin-gonic/gin"

func Ping (ctx *ginTonic.Context){

	ctx.String(200,"pong")
}