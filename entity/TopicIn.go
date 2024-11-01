package entity

type TopicIn struct {
	Address    string `json:"address,omitempty"`
	Author     string `json:"author,omitempty"`
	Code       string `json:"code,omitempty"`
	Content    string `json:"content,omitempty"`
	Download   string `json:"download,omitempty"`
	EnTitle    string `json:"enTitle,omitempty"`
	ID         int64  `json:"id,omitempty"`
	Label      string `json:"label,omitempty"`
	Language   string `json:"language,omitempty"`
	Source     string `json:"source,omitempty"`
	Title      string `json:"title,omitempty"`
	TopicTitle string `json:"topicTitle,omitempty"`
	Version    string `json:"version,omitempty"`
}
