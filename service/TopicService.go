package service

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"starfall-go/dao"
	"starfall-go/entity"
	"starfall-go/util"
	"strconv"
	"time"
)

type TopicService struct {
}

var topicDao = dao.TopicDao{}

type TopicOut entity.TopicOut

func (TopicService) FindAllTopic(c *gin.Context) {
	pageArg := c.PostForm("page")
	page, err := strconv.Atoi(pageArg)
	if err != nil || page <= 0 {
		page = 1
	}
	label := c.PostForm("label")
	version := c.PostForm("version")
	if label == "无" {
		label = ""
	}
	if version == "无" {
		version = ""
	}
	topics := topicDao.FindAllTopic((page-1)*10, label, version)
	topicsNum := topicDao.CountAllTopic(label, version)
	c.JSON(200, result.OkWithObj(gin.H{
		"topics": topics,
		"count":  topicsNum,
	}))
}

func (TopicService) GetTopicInfo(c *gin.Context) {
	idStr := c.PostForm("id")
	token := c.GetHeader("Authorization")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("ID is not a number"))
		return
	}
	topicOut := topicDao.FindTopicById(id)
	if topicOut.ID != 0 {
		if token != "" {
			topicDao.UpdateTopicView(int64(id), topicOut.View+1)
		}
		topicOut.MaxExp = util.GetMaxExp(topicOut.Level)
		c.JSON(200, result.OkWithObj(topicOut))
		return
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("This topic does not exist"))
	return
}

func (TopicService) FindAllTopicByUser(c *gin.Context) {
	pageStr := c.PostForm("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("page is not a number"))
		return
	}
	user := c.PostForm("user")
	topics := topicDao.FindTopicByUser((page-1)*10, user)
	count := topicDao.CountTopicByUser(user)
	c.JSON(200, result.OkWithObj(gin.H{
		"topics": topics,
		"count":  count,
	}))
}

func (TopicService) FindTopicVersion(c *gin.Context) {
	c.JSON(200, result.OkWithObj(topicDao.FindTopicVersion()))
}

func (TopicService) GetLike(c *gin.Context) {
	topicIdStr := c.PostForm("id")
	topicId, err := strconv.Atoi(topicIdStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("ID is not a number"))
		return
	}
	token := c.GetHeader("Authorization")
	_, claim, _ := util.ParseToken(token)
	user := claim.User
	likeLog := topicDao.FindLikeLogByTopicIdAndUser(topicId, user)
	status := "NOT_LIKE"
	var count int64 = 0
	if likeLog.Status == 1 {
		count = topicDao.CountLikeLogByTopicIdAndLike(topicId)
		status = "IS_LIKE"
	} else if likeLog.Status == 2 {
		status = "IS_DISLIKE"
	}
	c.JSON(200, result.OkWithObj(gin.H{
		"status": status,
		"count":  count,
	}))
}

func (TopicService) Like(c *gin.Context) {
	token := c.GetHeader("Authorization")
	idStr := c.PostForm("id")
	likeStr := c.PostForm("like")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("ID is not a number"))
		return
	}
	like, err := strconv.Atoi(likeStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("like value is not a number"))
		return
	}
	_, claim, _ := util.ParseToken(token)
	user := claim.User
	date := time.Now().Format("2006-01-02 15:04:05")
	likeObj := topicDao.FindLikeLogByTopicIdAndUser(id, user)
	if likeObj.TopicID != 0 {
		if int(likeObj.Status) != like {
			topicDao.UpdateLikeStateByTopicAndUser(int64(id), int64(like), user, date)
			if like == 1 {
				c.JSON(200, result.OkWithObj(topicDao.CountLikeLogByTopicIdAndLike(id)))
				return
			}
			c.JSON(200, result.Ok())
			return
		}
		topicDao.UpdateLikeStateByTopicAndUser(int64(id), 0, user, date)
		c.JSON(200, result.OkWithMsgAndObj("ALREADY_LIKE", nil))
		return
	}
	topicDao.InsertLike(entity.LikeLog{TopicID: int64(id), Status: int64(like), User: user, Date: date})
	if like == 1 {
		c.JSON(200, result.OkWithObj(topicDao.CountLikeLogByTopicIdAndLike(id)))
		return
	}
	c.JSON(200, result.Ok())
}

func (TopicService) FindCommentById(c *gin.Context) {
	idStr := c.PostForm("id")
	pageStr := c.PostForm("page")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("ID is not a number"))
		return
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("page is not a number"))
		return
	}
	comments := topicDao.FindCommentByTopicId(id, (page-1)*10)
	for i := range comments {
		comments[i].MaxExp = util.GetMaxExp(comments[i].Level)
	}
	count := topicDao.CountCommentByTopicId(id)
	c.JSON(200, result.OkWithObj(gin.H{
		"comments": comments,
		"count":    count,
	}))
}

func (TopicService) AppendComment(c *gin.Context) {
	token := c.GetHeader("Authorization")
	_, claim, _ := util.ParseToken(token)
	codeStr := c.PostForm("code")
	idStr := c.PostForm("id")
	content := c.PostForm("content")
	id, _ := strconv.Atoi(idStr)
	codeId, code := util.GetCodeAndIdByCode(codeStr)
	if util.VerifyCaptchaCode(codeId, code) {
		date := time.Now().Format("2006-01-02 15:04:05")
		topicDao.InsertComment(entity.CommentCreate{TopicID: int64(id), Date: date, User: claim.User, Content: content})
		count := topicDao.CountCommentByTopicId(id)
		topicDao.UpdateTopicComment(int64(id), count)
		c.JSON(200, result.OkWithObj(count))
		return
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("The code is wrong"))
}

func (TopicService) DeleteComment(c *gin.Context) {
	token := c.GetHeader("Authorization")
	_, claim, _ := util.ParseToken(token)
	idStr := c.PostForm("id")
	dateStr := c.PostForm("date")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("Id must be a number"))
		return
	}
	_, err = time.Parse("2006-01-02 15:04:05", dateStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("The format of date is wrong"))
		return
	}
	status1 := topicDao.DeleteCommentByIdAndUserAndDate(id, claim.User, dateStr)
	count := topicDao.CountCommentByTopicId(id)
	status2 := topicDao.UpdateTopicComment(int64(id), count)
	if !(status1 && status2) {
		c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("DataSource error"))
		return
	}
	c.JSON(200, result.Ok())
}

func (TopicService) AppendTopic(c *gin.Context) {
	token := c.GetHeader("Authorization")
	_, claim, _ := util.ParseToken(token)
	var topic entity.TopicIn
	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, result.ErrorWithMsg("Date Error"))
		return
	}
	codeId, code := util.GetCodeAndIdByCode(topic.Code)
	if util.VerifyCaptchaCode(codeId, code) {
		user := userDao.FindUserWithUser(claim.User)
		level := user.Level
		if user.Level < 5 {
			c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("You are not allowed to post"))
			return
		}
		id := topicDao.FindALl()[0].ID + 1
		date := time.Now().Format("2006-01-02 15:04:05")
		status1 := topicDao.InsertTopic(entity.Topic{
			ID:      id,
			Title:   topic.Title,
			User:    claim.User,
			Date:    date,
			Label:   topic.Label,
			Version: topic.Version,
			View:    0,
			Comment: 0,
		})
		status2 := topicDao.InsertTopicItem(entity.TopicItem{
			TopicId:    id,
			TopicTitle: topic.TopicTitle,
			EnTitle:    topic.EnTitle,
			Content:    topic.Content,
			Author:     topic.Author,
			Source:     topic.Source,
			Language:   topic.Language,
			Address:    topic.Address,
			Download:   topic.Download,
		})
		if !(status1 && status2) {
			c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("DataSource error"))
			return
		}
		addExp := int64(rand.Intn(100) + 50)
		exp := user.Exp + addExp
		diff := util.CheckAndLevelUp(exp, level)
		if diff >= 0 {
			exp = diff
			level++
		}
		userDao.UpdateExp(user.User, exp, level)
		c.JSON(200, result.Ok())
		return
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("The code is wrong"))
}

func isPromiseToEditTopicFunc(token string, id int) (r string) {
	_, claim, _ := util.ParseToken(token)
	topicUser := topicDao.FindTopicById(id).User
	if topicUser == claim.User {
		return "ACCEPT"
	}

	return "You are not allowed to edit this topic"
}

func (TopicService) IsPromiseToEditTopic(c *gin.Context) {
	token := c.GetHeader("Authorization")
	idStr := c.PostForm("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("ID is not a number"))
		return
	}
	r := isPromiseToEditTopicFunc(token, id)
	if r == "ACCEPT" {
		c.JSON(200, result.Ok())
		return
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("You are not allowed to edit this topic"))
}

func (TopicService) FindTopicInfoToEdit(c *gin.Context) {
	token := c.GetHeader("Authorization")
	idStr := c.PostForm("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("ID is not a number"))
		return
	}
	r := isPromiseToEditTopicFunc(token, id)
	if r == "ACCEPT" {
		topicOut := topicDao.FindTopicById(id)
		if topicOut.ID == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("This topic does not exist"))
			return
		}
		c.JSON(200, result.OkWithObj(topicOut))
		return
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("You are not allowed to edit this topic"))
}

func (TopicService) UpdateTopic(c *gin.Context) {
	token := c.GetHeader("Authorization")
	_, claim, _ := util.ParseToken(token)
	var topic entity.TopicIn
	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, result.ErrorWithMsg("Date Error"))
		return
	}
	if isPromiseToEditTopicFunc(token, int(topic.ID)) == "ACCEPT" {
		codeId, code := util.GetCodeAndIdByCode(topic.Code)
		if util.VerifyCaptchaCode(codeId, code) {
			date := time.Now().Format("2006-01-02 15:04:05")
			topicOut := topicDao.FindTopicById(int(topic.ID))
			status1 := topicDao.UpdateTopicExpectCommentAndView(entity.Topic{
				ID:      topic.ID,
				Title:   topic.Title,
				User:    claim.User,
				Label:   topic.Label,
				Date:    date,
				Version: topic.Version,
				View:    topicOut.View,
				Comment: topicOut.Comment,
			})
			status2 := topicDao.UpdateTopicItem(entity.TopicItem{
				TopicId:    topic.ID,
				TopicTitle: topic.TopicTitle,
				EnTitle:    topic.EnTitle,
				Content:    topic.Content,
				Author:     topic.Author,
				Source:     topic.Source,
				Address:    topic.Address,
				Download:   topic.Download,
				Language:   topic.Language,
			})
			if !(status1 && status2) {
				c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("DataSource error"))
				return
			}
			c.JSON(200, result.Ok())
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("The code is wrong"))
		return
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("You are not allowed to edit this topic"))
}

func (TopicService) DeleteTopic(c *gin.Context) {
	token := c.GetHeader("Authorization")
	idStr := c.PostForm("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("ID is not a number"))
		return
	}
	if isPromiseToEditTopicFunc(token, id) == "ACCEPT" {
		status1 := topicDao.DeleteCommentByTopicId(id)
		status2 := topicDao.DeleteLikeLog(id)
		status3 := topicDao.DeleteTopicItem(id)
		status4 := topicDao.DeleteTopic(id)
		if !(status1 && status2 && status3 && status4) {
			c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("DataSource error"))
			return
		}
		c.JSON(200, result.Ok())
		return
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("You are not allowed to delete this topic"))
}

func (TopicService) Search(c *gin.Context) {
	keyStr := c.PostForm("key")
	classification := c.PostForm("classification")
	pageStr := c.PostForm("page")
	if keyStr == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("The key can not empty"))
		return
	}
	if pageStr == "" || pageStr == "0" {
		c.JSON(200, result.OkWithObj(""))
		return
	}
	page, _ := strconv.Atoi(pageStr)
	keyStr = "%" + keyStr + "%"
	c.JSON(200, result.OkWithObj(gin.H{
		"search": topicDao.SearchByKey(keyStr, classification, (page-1)*10),
		"count":  topicDao.CountSearchByKey(keyStr, classification),
	}))
}
