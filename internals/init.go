package internals

import (
	"net/http"

	"github.com/CooffeeCoder-Go/agenda-emails/internals/configs"
	"github.com/gin-gonic/gin"
)


func Init(){
	r := gin.Default()

	configs.InitEnv()
	configs.Connect()

	r.GET("/",func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,gin.H{
			"msg":"Bem vindo ao meu app!",
			"warning": "Este primeiro commit é só um teste, isto não será uma api.",
		})
	})

	r.Run(":8080")
}