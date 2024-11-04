package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"starfall-go/dao"
	"starfall-go/entity"
	"starfall-go/util"
	"strings"
)

var result = entity.Result{}
var userDao = dao.UserDao{}

type User entity.User
type UserService struct {
}

func (UserService) Login(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")
	codeAndId := c.PostForm("code")
	codeId := strings.Split(codeAndId, ":")[0]
	code := strings.Split(codeAndId, ":")[1]
	util.VerifyCaptchaCode(codeId, code)
	if true {
		user := userDao.FindUserWithUserOrEmail(account)
		fmt.Println(user.User)
		if user.User != "" {
			passwordE, _ := util.AesDecrypt(user.Password)
			if passwordE == password {
				token := util.GenerateToken(user.User)
				c.JSON(200, result.OkWithObj(token))
				return
			}
			c.AbortWithStatusJSON(http.StatusNotAcceptable, result.ErrorWithMsg("password is wrong"))
			return
		}
		c.AbortWithStatusJSON(http.StatusNotAcceptable, result.ErrorWithMsg("account is not exist"))
		return
	}
	c.AbortWithStatusJSON(http.StatusNotAcceptable, result.ErrorWithMsg("code is wrong"))
}

func (UserService) GetUserInfo(c *gin.Context) {
	token := c.GetHeader("Authorization")
	claim, _ := util.ParseToken(token)
	user := claim["user"].(string)
	if user == "" {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, result.ErrorWithMsg("The User is null in the token"))
		return
	}
	userObj := userDao.FindUserWithUser(user)
	userObj.Password = "***"
	c.JSON(200, result.OkWithObj(userObj))
}
