package model

import "time"

type MandaAuth struct {
	UserID   string `json:"user_id" gorm:"unique"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
}
