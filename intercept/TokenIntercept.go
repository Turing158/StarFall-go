package intercept

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"starfall-go/dao"
	"starfall-go/entity"
	"starfall-go/util"
	"time"
)

var passUrl = []string{
	"/",
	"/login",
	"/findAllNotice",
}

var dbUser = dao.UserDao{}

func TokenIntercept() gin.HandlerFunc {
	return func(c *gin.Context) {
		//c.Next()
		//return

		url := c.Request.URL.Path
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
		claim, err := util.ParseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, result.ErrorWithMsg(err.Error()))
			return
		}

		if !claim.VerifyExpiresAt(time.Now().Unix(), true) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, result.ErrorWithMsg("The token has expired"))
			return
		}
		//token刷新
		userObj := dbUser.FindUserWithUserOrEmail(claim["user"].(string))
		if userObj.User == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, result.ErrorWithMsg("The user was exist in the old Token"))
		}
		newToken := util.GenerateToken(userObj.User)
		c.Header("Authorization", newToken)

		c.Next()
	}
}
