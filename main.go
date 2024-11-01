package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"starfall-go/controller"
	"starfall-go/entity"
)

func main() {
	engine := gin.Default()
	engine.Use(cors.Default())
	engine.GET("/", func(context *gin.Context) {
		context.JSON(200, entity.Result{}.Ok())
	})
	engine.NoRoute(func(context *gin.Context) {
		context.JSON(404, "ERROR")
	})

	//注册控制类
	controller.UserController{}.Register(engine)

	engine.Run()
}
