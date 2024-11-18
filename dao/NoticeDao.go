package dao

import (
	"starfall-go/entity"
	"starfall-go/util"
)

type NoticeDao struct {
}
type Notice entity.Notice

func (NoticeDao) FindAllNotice() (notices []Notice) {
	util.DB.Table("notice").Find(&notices)
	return notices
}
