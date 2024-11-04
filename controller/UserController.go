package controller

import (
	"github.com/gin-gonic/gin"
	"starfall-go/service"
)

var userService = service.UserService{}

type UserController struct {
}

func (UserController) Register(engine *gin.Engine) {
	engine.POST("/login", userService.Login)
	engine.POST("/getUserInfo", userService.GetUserInfo)
}
