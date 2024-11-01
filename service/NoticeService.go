package service

import (
	"github.com/gin-gonic/gin"
	"starfall-go/dao"
	"starfall-go/entity"
)

type NoticeService struct {
}

var noticeDao = dao.NoticeDao{}

func (NoticeService) FindAllNotice(c *gin.Context) {
	c.JSON(200, entity.Result{}.OkWithObj(noticeDao.FindAllNotice()))
}
