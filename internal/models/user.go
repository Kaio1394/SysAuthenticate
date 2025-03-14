package models

type User struct {
	Id       uint   `gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

func (User) TableName() string {
	return "t_auth_user"
}
