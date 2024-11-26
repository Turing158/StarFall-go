package controller

import (
	"github.com/gin-gonic/gin"
	"starfall-go/service"
)

var userService = service.UserService{}

func UserControllerRegister(engine *gin.Engine) {
	engine.POST("/login", userService.Login)
	engine.POST("/getUserInfo", userService.GetUserInfo)
	engine.POST("/register", userService.Register)
	engine.POST("/getEmailCode", userService.GetRegEmailCode)
	engine.POST("/getForgetEmailCode", userService.GetForgerPasswordEmailCode)
	engine.POST("/checkForgetPassword", userService.CheckForgetPassword)
	engine.POST("/forgetPassword", userService.ForgerPassword)
	engine.POST("/findUserByUser", userService.FindUserByUser)
	engine.POST("/updateUserInfo", userService.SettingInfo)
	engine.POST("/updatePassword", userService.SettingPassword)
	engine.POST("/updateAvatar", userService.SettingAvatar)
	engine.POST("/getOldEmailCode", userService.GetOldEmailCode)
	engine.POST("/getNewEmailCode", userService.GetNewEmailCode)
	engine.POST("/updateEmail", userService.SettingEmail)
	engine.POST("/findAllSignIn", userService.FindAlreadySignIn)
	engine.POST("/signIn", userService.SignIn)
}
