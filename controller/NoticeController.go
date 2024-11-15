package controller

import (
	"github.com/gin-gonic/gin"
	"starfall-go/service"
)

var noticeService = service.NoticeService{}

func NoticeControllerRegister(engine *gin.Engine) {
	engine.Any("/findAllNotice", noticeService.FindAllNotice)
}
