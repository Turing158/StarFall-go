package entity

type UserOut struct {
	Avatar   string `json:"avatar,omitempty"`
	Birthday string `json:"birthday,omitempty"`
	Email    string `json:"email,omitempty"`
	Exp      int64  `json:"exp,omitempty"`
	MaxExp   int64  `json:"maxExp"`
	Gender   int64  `json:"gender,omitempty"`
	Level    int64  `json:"level,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role,omitempty"`
	User     string `json:"user,omitempty" gorm:"primaryKey"`
}
