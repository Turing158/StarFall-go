package entity

type TopicCreate struct {
	Comment int64  `json:"comment,omitempty"`
	Date    string `json:"date,omitempty"`
	ID      int64  `json:"id,omitempty" gorm:"primaryKey"`
	Label   string `json:"label,omitempty"`
	Title   string `json:"title,omitempty"`
	User    string `json:"user,omitempty"`
	Version string `json:"version,omitempty"`
	View    int64  `json:"view,omitempty"`
}
