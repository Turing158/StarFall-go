package dao

import (
	"starfall-go/entity"
	"starfall-go/util"
)

var dbMsg = util.DB.Table("chat_notice c")
var selectMsg = "from_user as fromUser,fu.name as fromName,fu.avatar as fromAvatar,to_user as toUser,tu.name as toName,tu.avatar as toAvatar,date,content"

type MessageDao struct {
}

func (MessageDao) FindAllMsgByToUser(user string) (msgs []entity.Message) {
	util.DB.Table("chat_notice c").Select(selectMsg).Where("to_user = ? or from_user = ?", user, user).Joins("join starfall.user fu on c.from_user = fu.user join starfall.user tu on c.to_user = tu.user").Order("date desc").Find(&msgs)
	return
}

func (MessageDao) FindMsgByToUserAndFromUser(toUser, fromUser string) (msgs []entity.Message) {
	util.DB.Table("chat_notice c").Select(selectMsg).Where("(to_user = ?1 and from_user = ?2) or (to_user = ?2 and from_user= ?1)", toUser, fromUser).Joins("join starfall.user fu on c.from_user = fu.user join starfall.user tu on c.to_user = tu.user").Order("date desc").Find(&msgs)
	return
}

func (MessageDao) FindFromUserMsgByFromUserAndToUser(fromUser, toUser string) (msgs []entity.Message) {
	util.DB.Table("chat_notice c").Select(selectMsg).Where("from_user = ? and to_user = ?", fromUser, toUser).Joins("join starfall.user fu on c.from_user = fu.user join starfall.user tu on c.to_user = tu.user").Order("date desc").Find(&msgs)
	return
}

func (MessageDao) UpdateMsgContent(message entity.Message) bool {
	re := util.DB.Table("chat_notice c").Where("from_user = ? and to_user = ? and date = ?", message.FromUser, message.ToUser, message.Date).First(&entity.Message{}).Update("content", message.Content).RowsAffected
	return util.Int64ToBool(re)
}

func (MessageDao) InsertMsg(message entity.MessageCreate) bool {
	re := util.DB.Table("chat_notice c").Create(message).RowsAffected
	return util.Int64ToBool(re)
}
