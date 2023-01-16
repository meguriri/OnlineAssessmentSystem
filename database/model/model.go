package model

type User struct {
	ID       string `json:"id" gorm:"id"`
	Password string `json:"password" gorm:"password"`
	Type     int    `json:"type" gorm:"type"`
	Name     string `json:"name" gorm:"name"`
}
