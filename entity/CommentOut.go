package entity

type CommentOut struct {
	Content string `json:"content,omitempty"`
	Date    string `json:"date,omitempty"`
	TopicID int64  `json:"topicId,omitempty"  gorm:"column:topicId"`
	User    string `json:"user,omitempty"`
	Name    string `json:"name,omitempty"`
	Avatar  string `json:"avatar,omitempty"`
	Level   int64  `json:"level,omitempty"`
	Exp     int64  `json:"exp,omitempty"`
	MaxExp  int64  `json:"maxExp,omitempty"`
}
