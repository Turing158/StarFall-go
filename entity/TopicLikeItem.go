package entity

type TopicLikeItem struct {
	Id      int    `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	Label   string `json:"label,omitempty"`
	User    string `json:"user,omitempty"`
	Name    string `json:"name,omitempty"`
	Like    int    `json:"like,omitempty"`
	Dislike int    `json:"dislike,omitempty"`
}
