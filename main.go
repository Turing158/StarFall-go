package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"starfall-go/controller"
	"starfall-go/entity"
	"starfall-go/util"
)

func main() {
	engine := gin.Default()
	engine.Use(cors.Default())
	engine.Use(Logger())
	engine.Use(util.TokenIntercept())
	engine.GET("/", func(context *gin.Context) {
		context.JSON(200, entity.Result{}.Ok())
	})
	engine.NoRoute(func(context *gin.Context) {
		context.JSON(404, entity.Result{}.ErrorWithMsg("Unknown Path"))
	})

	//注册控制类
	controller.UserController{}.Register(engine)
	controller.NoticeController{}.Register(engine)

	engine.Run()
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Request URL:", c.Request.URL)
		c.Next()
		log.Println("Response Status:", c.Writer.Status())
	}
}
