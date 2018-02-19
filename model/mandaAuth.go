package model

type MandaAuth struct {
	UserID   string `json:"user_id" gorm:"unique"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
