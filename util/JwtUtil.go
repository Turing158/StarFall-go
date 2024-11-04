package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"starfall-go/dao"
	"starfall-go/entity"
	"time"
)

const exp = time.Hour * 24

var jwtKey = []byte("StarFall")

func GenerateToken(user string) (token string, err error) {
	claims := jwt.MapClaims{
		"user": user,
		"exp":  time.Now().Add(exp).Unix(),
	}
	claimWithSign := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return claimWithSign.SignedString(jwtKey)
}

func ParseToken(token string) (claim jwt.MapClaims, err error) {
	_, err = jwt.ParseWithClaims(token, &claim, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return claim, err
}

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
		claim, err := ParseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, result.ErrorWithMsg("Unknown error in token"))
			return
		}

		if !claim.VerifyExpiresAt(time.Now().Unix(), true) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, result.ErrorWithMsg("The token has expired"))
			return
		}
		//token刷新
		//userObj := dbUser.FindUserWithUserOrEmail(claim["user"].(string))
		//if userObj.User == ""{
		//	c.AbortWithStatusJSON(http.StatusUnauthorized, result.ErrorWithMsg("The user was exist in the old Token"))
		//}
		//newToken, _ := GenerateToken(userObj.User)
		//c.Header("Authorization", newToken)

		c.Next()
	}
}
