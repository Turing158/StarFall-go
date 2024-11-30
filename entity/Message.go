package entity

type Message struct {
	Content    string `json:"content,omitempty" gorm:"column:content"`
	Date       string `json:"date,omitempty" gorm:"column:date"`
	FromAvatar string `json:"fromAvatar,omitempty" gorm:"column:fromAvatar"`
	FromName   string `json:"fromName,omitempty" gorm:"column:fromName"`
	FromUser   string `json:"fromUser,omitempty" gorm:"column:fromUser"`
	ToAvatar   string `json:"toAvatar,omitempty" gorm:"column:toAvatar"`
	ToName     string `json:"toName,omitempty" gorm:"column:toName"`
	ToUser     string `json:"toUser,omitempty" gorm:"column:toUser"`
}
