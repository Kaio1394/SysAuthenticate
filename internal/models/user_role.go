package models

type UserRole struct {
	UserID int `gorm:"primaryKey" json:"user_id"`
	RoleID int `gorm:"primaryKey" json:"role_id"`

	User User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Role Role `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (UserRole) TableName() string { return "t_user_role" }
