package entity

type TopicOut struct {
	Address    string `json:"address,omitempty"`
	Author     string `json:"author,omitempty"`
	Avatar     string `json:"avatar,omitempty"`
	Comment    int64  `json:"comment,omitempty"`
	Content    string `json:"content,omitempty"`
	Date       string `json:"date,omitempty"`
	Download   string `json:"download,omitempty"`
	EnTitle    string `json:"enTitle,omitempty"`
	Exp        int64  `json:"exp,omitempty"`
	ID         int64  `json:"id,omitempty" gorm:"primaryKey"`
	Label      string `json:"label,omitempty"`
	Language   string `json:"language,omitempty"`
	Level      int64  `json:"level,omitempty"`
	MaxExp     int64  `json:"maxExp,omitempty"`
	Name       string `json:"name,omitempty"`
	OldID      int64  `json:"oldId,omitempty"`
	Source     string `json:"source,omitempty"`
	Title      string `json:"title,omitempty"`
	TopicTitle string `json:"topicTitle,omitempty" gorm:"column:topicTitle"`
	User       string `json:"user,omitempty"`
	Version    string `json:"version,omitempty"`
	View       int64  `json:"view,omitempty"`
}
