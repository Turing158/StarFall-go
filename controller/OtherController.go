package controller

import (
	"github.com/gin-gonic/gin"
	"starfall-go/service"
)

type OtherController struct {
}

var otherService = service.OtherService{}

func (OtherController) Register(engine *gin.Engine) {
	engine.GET("/getCodeImage", otherService.GetCodeImage)
}
