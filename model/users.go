package model

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Uid      string `json:"uid"` // UserID as specified by OIDC OP
	Email    string `json:"email"`
	Username string `json:"username"`
}
