package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"starfall-go/controller"
	"starfall-go/entity"
	"starfall-go/intercept"
	"starfall-go/util"
)

func main() {
	engine := gin.Default()
	util.InitRedis()
	defer util.CloseRedis()
	//engine.Use(cors.Default())
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}
	engine.Use(cors.New(config))
	engine.Use(Logger())
	engine.Use(intercept.TokenIntercept())
	engine.GET("/", func(context *gin.Context) {
		context.JSON(200, entity.Result{}.Ok())
	})

	engine.NoRoute(func(context *gin.Context) {
		context.JSON(404, entity.Result{}.ErrorWithMsg("Unknown Path"))
	})

	//注册控制类
	controller.UserControllerRegister(engine)
	controller.OtherControllerRegister(engine)
	controller.NoticeControllerRegister(engine)
	controller.TopicControllerRegister(engine)

	engine.Run(":9090")
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Request URL:", c.Request.URL)
		c.Next()
		log.Println("Response Status:", c.Writer.Status())
	}
}
