package entity

type LikeItem struct {
	Date    string `json:"date,omitempty"`
	Status  int64  `json:"status,omitempty"`
	TopicID int64  `json:"topicId,omitempty" gorm:"column:topicId"`
	Name    string `json:"name,omitempty"`
	User    string `json:"user,omitempty"`
}
