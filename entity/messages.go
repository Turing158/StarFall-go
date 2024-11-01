package entity

type Messages struct {
	NewMessage Message `json:"newMessage,omitempty"`
	OldMessage Message `json:"oldMessage,omitempty"`
}
