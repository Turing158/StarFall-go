package entity

type CommentCreate struct {
	Content string `json:"content,omitempty"`
	Date    string `json:"date,omitempty"`
	TopicID int64  `json:"topicId,omitempty"  gorm:"column:topicId"`
	User    string `json:"user,omitempty"`
}
