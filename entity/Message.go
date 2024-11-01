package entity

type Message struct {
	Content    string `json:"content,omitempty"`
	Date       string `json:"date,omitempty"`
	FromAvatar string `json:"fromAvatar,omitempty"`
	FromName   string `json:"fromName,omitempty"`
	FromUser   string `json:"fromUser,omitempty"`
	ToAvatar   string `json:"toAvatar,omitempty"`
	ToName     string `json:"toName,omitempty"`
	ToUser     string `json:"toUser,omitempty"`
}
