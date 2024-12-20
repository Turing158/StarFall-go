package entity

type TopicItem struct {
	Address    string `json:"address,omitempty"`
	Author     string `json:"author,omitempty"`
	Content    string `json:"content,omitempty"`
	Download   string `json:"download,omitempty"`
	EnTitle    string `json:"enTitle,omitempty"  gorm:"column:enTitle"`
	TopicId    int64  `json:"topicId,omitempty"  gorm:"column:topicId"`
	Language   string `json:"language,omitempty"`
	Source     string `json:"source,omitempty"`
	TopicTitle string `json:"topicTitle,omitempty"  gorm:"column:topicTitle"`
}
