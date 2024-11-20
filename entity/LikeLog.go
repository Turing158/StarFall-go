package entity

type LikeLog struct {
	Date    string `json:"date,omitempty"`
	Status  int64  `json:"status,omitempty"`
	TopicID int64  `json:"topicId,omitempty"  gorm:"column:topicId"`
	User    string `json:"user,omitempty"`
}
