package model

type User struct {
	ID       string `json:"id" gorm:"id"`
	Password string `json:"password" gorm:"password"`
	Type     int    `json:"type" gorm:"type"`
	Name     string `json:"name" gorm:"name"`
}

type KnowledgePoint struct {
	ID            int    `json:"id" gorm:"id"`
	Name          string `json:"name" gorm:"name"`
	FacilityValue int    `json:"facility_value" gorm:"facility_value"`
}
