package dao

import (
	"starfall-go/entity"
	"starfall-go/util"
)

type NoticeDao struct {
}
type Notice entity.Notice

var dbNotice = util.DB.Table("notice")

func (NoticeDao) FindAllNotice() []Notice {
	var notices []Notice
	dbNotice.Find(&notices)
	return notices
}
