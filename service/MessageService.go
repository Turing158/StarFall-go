package service

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
	"starfall-go/dao"
	"starfall-go/entity"
	"starfall-go/util"
	"strings"
	"time"
)

type MessageService struct {
}

var messageDao = dao.MessageDao{}

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

func (MessageService) GetMessageList(c *gin.Context) {
	token := c.GetHeader("Authorization")
	_, claim, _ := util.ParseToken(token)
	messages := messageDao.FindAllMsgByToUser(claim.User)
	var messageList []entity.MessageTerm
	for i := range messages {
		contents := strings.Split(messages[i].Content, "[&divide&]")
		lastContent := contents[len(contents)-1]
		if len(messageList) == 0 {
			messageListAdd(&messageList, messages[i], claim.User, lastContent)
		} else {
			flag := true
			for j := range messageList {
				if messageList[j].User == messages[i].ToUser || messageList[j].User == messages[i].FromUser {
					flag = false
					break
				}
			}
			if flag {
				messageListAdd(&messageList, messages[i], claim.User, lastContent)
			}
		}
	}
	c.JSON(200, result.OkWithObj(messageList))
}

func messageListAdd(messageList *[]entity.MessageTerm, messages entity.Message, user, lastContent string) {
	var messageTerm entity.MessageTerm
	if messages.FromUser == user {
		messageTerm.User = messages.ToUser
		messageTerm.Name = messages.ToName
		messageTerm.Avatar = messages.ToAvatar
		messageTerm.LastContent = lastContent
	} else {
		messageTerm.User = messages.FromUser
		messageTerm.Name = messages.FromName
		messageTerm.Avatar = messages.FromAvatar
		messageTerm.LastContent = lastContent
	}
	*messageList = append(*messageList, messageTerm)
}

func (MessageService) GetMsgByToUserAndFromUser(c *gin.Context) {
	token := c.GetHeader("Authorization")
	_, claim, _ := util.ParseToken(token)
	fromUser := c.PostForm("fromUser")
	messages := messageDao.FindMsgByToUserAndFromUser(claim.User, fromUser)
	c.JSON(200, result.OkWithObj(messages))
}

func (MessageService) SendMessage(c *gin.Context) {
	token := c.GetHeader("Authorization")
	_, claim, _ := util.ParseToken(token)
	content := c.PostForm("content")
	toUser := c.PostForm("toUser")
	fromUser := claim.User
	fromUserObj := userDao.FindUserWithUser(fromUser)
	toUserObj := userDao.FindUserWithUser(toUser)
	if fromUserObj.User != "" {
		date := time.Now().Format("2006-01-02 15:04:05")
		message := entity.Message{
			Content:    content,
			Date:       date,
			FromAvatar: fromUserObj.Avatar,
			FromName:   fromUserObj.Name,
			FromUser:   fromUserObj.User,
			ToAvatar:   toUserObj.Avatar,
			ToName:     toUserObj.Name,
			ToUser:     toUserObj.User,
		}
		messageData, _ := json.Marshal(message)
		util.WsSendMessage(toUser, messageData)
		fromUserMsgs := messageDao.FindFromUserMsgByFromUserAndToUser(fromUser, toUser)
		if len(fromUserMsgs) == 0 {
			messageDao.InsertMsg(entity.MessageCreate{
				Content:  content,
				Date:     date,
				FromUser: fromUser,
				ToUser:   toUser,
			})
		} else {
			fromUserMsg := fromUserMsgs[0]
			oldDateTime, _ := time.Parse("2006-01-02 15:04:05", fromUserMsg.Date)
			newDateTime := time.Now()
			if newDateTime.Sub(oldDateTime).Minutes() < 2 {
				newContent := fromUserMsg.Content + "[&divide&]" + content
				messageDao.UpdateMsgContent(entity.Message{
					Content:  newContent,
					Date:     fromUserMsg.Date,
					FromUser: fromUser,
					ToUser:   toUser,
				})
			} else {
				messageDao.InsertMsg(entity.MessageCreate{
					Content:  content,
					Date:     date,
					FromUser: fromUser,
					ToUser:   toUser,
				})
			}

		}
		c.JSON(200, result.OkWithObj(message))
		return
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("Unknown user"))
}
