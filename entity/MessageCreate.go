package entity

type MessageCreate struct {
	Content  string `json:"content,omitempty"`
	Date     string `json:"date,omitempty"`
	FromUser string `json:"fromUser,omitempty"`
	ToUser   string `json:"toUser,omitempty"`
}
