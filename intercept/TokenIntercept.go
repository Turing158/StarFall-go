package intercept

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"starfall-go/dao"
	"starfall-go/entity"
	"starfall-go/util"
	"strings"
	"time"
)

var passUrl = []string{
	"/",
	"/getCodeImage",
	"/login",
	"/register",
	"/getEmailCode",
	"/findUserByUser",
	"/findAllNotice",
	"/findAllTopic",
	"/getTopicInfo",
	"/findAllTopicByUser",
	"/findTopicVersion",
	"/findCommentByTopic",
	"/search",
	"/getForgetEmailCode",
	"/checkForgetPassword",
	"/getLike",
}

var dbUser = dao.UserDao{}

func TokenIntercept() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL.Path
		if strings.Contains(url, "/message") {
			c.Next()
			return
		}
		for s := range passUrl {
			if passUrl[s] == url {
				c.Next()
				return
			}
		}
		result := entity.Result{}
		token := c.GetHeader("Authorization")
		if len(token) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, result.ErrorWithMsg("The token is null"))
			return
		}
		claim, userClaim, err := util.ParseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, result.ErrorWithMsg(err.Error()))
			return
		}

		if !claim.VerifyExpiresAt(time.Now().Unix(), true) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, result.ErrorWithMsg("The token has expired"))
			return
		}
		if userClaim.Role == "@ForgetPassword" {
			if url == "/forgetPassword" || url == "/checkForgetPassword" {
				c.Next()
				return
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, result.ErrorWithMsg("The current token does not allow access, please log in again"))
			return
		}
		//token刷新
		userObj := dbUser.FindUserWithUserOrEmail(userClaim.User)
		if userObj.User == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, result.ErrorWithMsg("The user was exist in the old Token"))
			return
		}
		newClaim := util.UserClaim{User: userObj.User, Email: userObj.Email, Role: userObj.Role}
		newToken := util.GenerateToken(newClaim)
		c.Header("Authorization", newToken)
		c.Next()
	}
}
