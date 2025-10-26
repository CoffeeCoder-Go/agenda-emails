package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()

	r.GET("/",func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,gin.H{
			"msg":"Bem vindo ao meu app!",
			"warning": "Este primeiro commit é só um teste, isto não será uma api.",
		})
	})

	r.Run(":8080")
}