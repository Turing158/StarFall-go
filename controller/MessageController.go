package controller

import (
	"github.com/gin-gonic/gin"
	"starfall-go/service"
)

var messageService = service.MessageService{}

func MessageControllerRegister(engine *gin.Engine) {
	engine.GET("/message/*token", messageService.HandleWebSocket)
	engine.POST("/findMessageList", messageService.GetMessageList)
	engine.POST("/findMsgByToUserAndFromUser", messageService.GetMsgByToUserAndFromUser)
	engine.POST("/sendMessage", messageService.SendMessage)
}
