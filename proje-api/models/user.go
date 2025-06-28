package models

type User struct {
	ID       uint64 `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"not null"`
	Name     string `json:"name" gorm:"not null"`
	Surname  string `json:"surname" gorm:"not null"`
	Email    string `json:"email" gorm:"not null;unique"`
	RoleID   int    `json:"role_id" gorm:"not null"`
}
