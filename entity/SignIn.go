package entity

type SignIn struct {
	Date    string `json:"date,omitempty"`
	Emotion string `json:"emotion,omitempty"`
	Message string `json:"message,omitempty"`
	//Name    string `json:"name,omitempty"`
	User string `json:"user,omitempty"`
}
