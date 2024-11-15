package controller

import (
	"github.com/gin-gonic/gin"
	"starfall-go/service"
)

var otherService = service.OtherService{}

func OtherControllerRegister(engine *gin.Engine) {
	engine.GET("/getCodeImage", otherService.GetCodeImage)
}
