package entity

type Search struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Label   string `json:"label"`
	Content string `json:"content"`
	View    int64  `json:"view"`
	Comment int64  `json:"comment"`
	Date    string `json:"date"`
	User    string `json:"user"`
	Name    string `json:"name"`
}
