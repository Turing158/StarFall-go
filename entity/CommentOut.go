package entity

type CommentOut struct {
	Content string `json:"content,omitempty"`
	Date    string `json:"date,omitempty"`
	TopicID int64  `json:"topicId,omitempty"`
	User    string `json:"user,omitempty"`
	Name    string `json:"name,omitempty"`
	Avatar  string `json:"avatar,omitempty"`
	Level   int    `json:"level,omitempty"`
	exp     int    `json:"exp,omitempty"`
	maxExp  int    `json:"maxExp,omitempty"`
}
