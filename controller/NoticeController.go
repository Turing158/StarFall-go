package controller

import (
	"github.com/gin-gonic/gin"
	"starfall-go/service"
)

var noticeService = service.NoticeService{}

type NoticeController struct {
}

func (NoticeController) Register(engine *gin.Engine) {
	engine.Any("/findAllNotice", noticeService.FindAllNotice)
}
