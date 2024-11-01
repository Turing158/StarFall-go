package entity

type Notice struct {
	Content string `json:"content,omitempty"`
	ID      int64  `json:"id,omitempty"`
}
