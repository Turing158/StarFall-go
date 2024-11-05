package entity

type Topic struct {
	Avatar  string `json:"avatar,omitempty"`
	Comment int64  `json:"comment,omitempty"`
	Date    string `json:"date,omitempty"`
	ID      int64  `json:"id,omitempty" gorm:"primaryKey"`
	Label   string `json:"label,omitempty"`
	Name    string `json:"name,omitempty"`
	Title   string `json:"title,omitempty"`
	User    string `json:"user,omitempty"`
	Version string `json:"version,omitempty"`
	View    int64  `json:"view,omitempty"`
}
