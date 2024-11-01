package entity

type UserAndOldUser struct {
	OldEmail string `json:"oldEmail,omitempty"`
	OldUser  string `json:"oldUser,omitempty"`
	User     User   `json:"user,omitempty"`
}
