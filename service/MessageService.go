package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"starfall-go/util"
	"strings"
)

type MessageService struct {
}

func (MessageService) HandleWebSocket(c *gin.Context) {
	token := strings.Trim(c.Param("token"), "/")
	_, claim, err := util.ParseToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("Unable to connect to the server"))
		return
	}
	conn, err := util.WebSocketUpGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("Unable to connect to the server"))
		return
	}
	util.WsMap.Store(claim.User, conn)
}
