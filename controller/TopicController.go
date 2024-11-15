package controller

import (
	"github.com/gin-gonic/gin"
	"starfall-go/service"
)

var topicService = service.TopicService{}

func TopicControllerRegister(engine *gin.Engine) {
	engine.Any("/findAllTopic", topicService.FindAllTopic)
	engine.POST("/getTopicInfo", topicService.GetTopicInfo)
	engine.POST("/findAllTopicByUser", topicService.FindAllTopicByUser)
	engine.Any("/findTopicVersion", topicService.FindTopicVersion)
	engine.POST("/getLike", topicService.GetLike)
	engine.POST("/like", topicService.Like)
	engine.POST("/findCommentByTopic", topicService.FindCommentById)
	engine.POST("/appendComment", topicService.AppendComment)
	engine.POST("/deleteComment", topicService.DeleteComment)
	engine.POST("/appendTopic", topicService.AppendTopic)
	engine.POST("/isPromiseToEdit", topicService.IsPromiseToEditTopic)
	engine.POST("/hasToPromiseToEdit", topicService.FindTopicInfoToEdit)
	engine.POST("/editTopic", topicService.UpdateTopic)
	engine.POST("/deleteTopic", topicService.DeleteTopic)
	engine.POST("/search", topicService.Search)

}
