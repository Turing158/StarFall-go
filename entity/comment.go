package entity

type Comment struct {
	Content    string `json:"content,omitempty"`
	Date       string `json:"date,omitempty"`
	OldDate    string `json:"oldDate,omitempty"`
	OldTopicID int64  `json:"oldTopicId,omitempty"`
	OldUser    string `json:"oldUser,omitempty"`
	TopicID    int64  `json:"topicId,omitempty"`
	User       string `json:"user,omitempty"`
}
