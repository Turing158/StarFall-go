package dao

import (
	"starfall-go/entity"
	"starfall-go/util"
)

type TopicDao struct {
}

const topicSelect = "t.id,t.title,t.label,t.user,t.date,t.view,t.comment,t.version"
const userSelect = "u.name,u.exp,u.level,u.avatar"
const topicItemSelect = "ti.topicTitle,ti.enTitle,ti.source,ti.author,ti.language,ti.address,ti.download,ti.content"
const commentSelect = "c.topicId,c.user,c.date,c.content"

func (TopicDao) FindALl() (topic []entity.Topic) {
	util.DB.Table("topic t").Order("id desc").Find(&topic)
	return
}

func (TopicDao) FindAllTopic(offset int, label, version string) (topic []entity.Topic) {
	table := util.DB.Table("topic t").Select(topicSelect + "," + userSelect).Joins("join starfall.user u on t.user = u.user")
	if label != "" && version == "" {
		table.Where("label = ?", label)
	} else if label == "" && version != "" {
		table.Where("version = ?", version).Order("date desc")
	} else if label != "" && version != "" {
		table.Where("label = ? and version = ?", label, version)
	}
	table.Order("date desc").Offset(offset).Limit(10).Find(&topic)
	return
}

func (TopicDao) CountAllTopic(label, version string) (count int64) {
	table := util.DB.Table("topic t")
	if label != "" && version == "" {
		table.Where("label = ?", label)
	} else if label == "" && version != "" {
		table.Where("version = ?", version)
	} else if label != "" && version != "" {
		table.Where("label = ? and version = ?", label, version)
	}
	table.Find(&entity.TopicOut{}).Count(&count)
	return
}

func (TopicDao) FindTopicVersion() (versions []string) {
	util.DB.Table("topic t").Distinct("version").Find(&versions)
	return
}

func (TopicDao) FindTopicById(id int) (topic entity.TopicOut) {
	util.DB.Table("topic t").Select(topicSelect+","+userSelect+","+topicItemSelect).Joins("join starfall.topicitem ti on t.id = ti.topicId").Joins("join starfall.user u on u.user = t.user").Where("id = ?", id).Find(&topic)
	return
}

func (TopicDao) FindTopicByUser(offset int, user string) (topics []entity.Topic) {
	util.DB.Table("topic t").Select(topicSelect+","+userSelect).Joins("join starfall.user u on t.user = u.user").Where("t.user = ?", user).Order("date desc").Offset(offset).Limit(10).Find(&topics)
	return
}

func (TopicDao) CountTopicByUser(user string) (count int64) {
	util.DB.Table("topic t").Where("user = ?", user).Find(&entity.TopicOut{}).Count(&count)
	return
}

func (TopicDao) FindCommentByTopicId(id, offset int) (comments []entity.CommentOut) {
	util.DB.Table("comment c").Select(commentSelect+","+userSelect).Joins("join starfall.user u on c.user = u.user").Where("topicId = ?", id).Order("date").Offset(offset).Limit(10).Find(&comments)
	return
}

func (TopicDao) CountCommentByTopicId(id int) (count int64) {
	util.DB.Table("comment c").Where("topicId = ?", id).Find(&entity.CommentOut{}).Count(&count)
	return
}

func (TopicDao) CountLikeLogByTopicIdAndLike(id int) (count int64) {
	util.DB.Table("likelog l").Where("topicId = ? and status = 1", id).Find(&entity.LikeLog{}).Count(&count)
	return
}

func (TopicDao) FindLikeLogByTopicIdAndUser(id int, user string) (likeLog entity.LikeLog) {
	util.DB.Table("likelog l").Where("topicId = ? and user = ?", id, user).Find(&likeLog)
	return
}

func (TopicDao) SearchByKey(key, classification string, offset int) (searches []entity.Search) {
	util.DB.Table("topic t").
		Select("t.id,"+
			"t.title,"+
			"t.label,"+
			"regexp_replace(regexp_replace(regexp_replace(regexp_replace(regexp_replace(regexp_replace(content,'<.+?>',''),'\\\\*\\\\*(.*?)\\\\*\\\\*', '$1'),'\\\\[(.*?)\\\\]\\\\((.*?)\\\\)', '$1'),'\\>\\\\s', ''),'#', ''),'\\\\*{3}|\\\\*\\\\s\\\\*\\\\s\\\\*', '') as content,"+
			"t.view,"+
			"t.comment,"+
			"t.date,"+
			"t.user,"+
			"u.name").
		Joins("join topicitem ti on t.id = ti.topicId join user u on t.user = u.user").
		Where(checkClassificationToWhereStr(classification), key).
		Order("t.date desc").
		Limit(10).
		Offset(offset).
		Find(&searches)
	return
}

func (TopicDao) CountSearchByKey(key, classification string) (count int64) {
	util.DB.Table("topic t").
		Joins("join topicitem ti on t.id = ti.topicId join user u on t.user = u.user").
		Where(checkClassificationToWhereStr(classification), key).
		Count(&count)
	return
}

func checkClassificationToWhereStr(classification string) string {
	if classification == "作者" {
		return "u.name like ?"
	} else if classification == "主题" {
		return "t.title like ?"
	} else if classification == "内容" {
		return "ti.content like ?"
	}
	return "t.title like ? or u.name like ? or ti.content like ?"
}

func (TopicDao) InsertLike(likeLog entity.LikeLog) bool {
	re := util.DB.Table("likelog").Create(likeLog).RowsAffected
	return util.Int64ToBool(re)
}

func (TopicDao) InsertComment(comment entity.CommentCreate) bool {
	re := util.DB.Table("comment").Create(comment).RowsAffected
	return util.Int64ToBool(re)
}

func (TopicDao) InsertTopic(topic entity.Topic) bool {
	re := util.DB.Table("topic").Create(topic).RowsAffected
	return util.Int64ToBool(re)
}

func (TopicDao) InsertTopicItem(topicItem entity.TopicItem) bool {
	re := util.DB.Table("topicItem").Create(topicItem).RowsAffected
	return util.Int64ToBool(re)
}

func (TopicDao) UpdateTopicView(id, view int64) bool {
	re := util.DB.Table("topic t").Where("id = ?", id).First(&entity.Topic{}).Update("view", view).RowsAffected
	return util.Int64ToBool(re)
}

func (TopicDao) UpdateTopicComment(id, comment int64) bool {
	re := util.DB.Table("topic t").Where("id = ?", id).First(&entity.Topic{}).Update("comment", comment).RowsAffected
	return util.Int64ToBool(re)
}

func (TopicDao) UpdateLikeStateByTopicAndUser(id, status int64, user, date string) bool {
	re := util.DB.Table("likelog l").Where("topicId = ? and user = ?", id, user).
		First(&entity.LikeLog{}).
		Update("date", date).
		Update("status", status).
		RowsAffected
	return util.Int64ToBool(re)
}

func (TopicDao) UpdateTopicExpectCommentAndView(topic entity.Topic) bool {
	re := util.DB.Table("topic t").Where("id = ?", topic.ID).First(&entity.Topic{}).Updates(entity.Topic{
		Title:   topic.Title,
		Label:   topic.Label,
		User:    topic.User,
		Date:    topic.Date,
		Version: topic.Version,
	}).RowsAffected
	return util.Int64ToBool(re)
}

func (TopicDao) UpdateTopicItem(item entity.TopicItem) bool {
	re := util.DB.Table("topicItem ti").Where("topicId = ?", item.TopicId).First(&entity.TopicItem{}).Updates(item).RowsAffected
	return util.Int64ToBool(re)
}

func (TopicDao) DeleteCommentByIdAndUserAndDate(topicId int, user, date string) bool {
	re := util.DB.Table("comment c").Where("topicId = ? and user = ? and date = ?", topicId, user, date).Delete(&entity.Comment{}).RowsAffected
	return util.Int64ToBool(re)
}

func (TopicDao) DeleteTopic(id int) bool {
	re := util.DB.Table("topic t").Where("id = ?", id).Delete(&entity.Topic{}).RowsAffected
	return util.Int64ToBool(re)
}

func (TopicDao) DeleteTopicItem(topicId int) bool {
	re := util.DB.Table("topicItem ti").Where("topicId = ?", topicId).Delete(&entity.TopicItem{}).RowsAffected
	return util.Int64ToBool(re)
}

func (TopicDao) DeleteLikeLog(topicId int) bool {
	re := util.DB.Table("likelog l").Where("topicId = ?", topicId).Delete(&entity.LikeLog{}).RowsAffected
	return util.Int64ToBool(re)
}

func (TopicDao) DeleteCommentByTopicId(topicId int) bool {
	re := util.DB.Table("comment c").Where("topicId = ?", topicId).Delete(&entity.Comment{}).RowsAffected
	return util.Int64ToBool(re)
}
