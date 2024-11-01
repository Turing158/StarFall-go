package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"starfall-go/entity"
	"time"
)

const exp = time.Hour * 24
const key = "StarFall"

func GenerateToken(user string) (token string, err error) {
	claims := jwt.MapClaims{
		"user": user,
		"exp":  time.Now().Add(exp).Unix(),
	}
	claimWithSign := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return claimWithSign.SignedString(key)
}

func ParseToken(token string) (claim jwt.MapClaims, err error) {
	claimP := &jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token, claim, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	return *claimP, err
}

var passUrl = []string{
	"/",
	"/findAllNotice",
}

func TokenIntercept() gin.HandlerFunc {
	return func(c *gin.Context) {

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
		//newToken := ""
		//c.Header("Authorization", newToken)

		c.Next()
	}
}
