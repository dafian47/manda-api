package model

import "time"

type MandaChannel struct {
	ID           string     `json:"id" gorm:"primary_key;unique"`
	Name         string     `json:"name" binding:"required"`
	Status       string     `json:"status" binding:"required" gorm:"index"`
	UrlIcon      string     `json:"url_icon"`
	UserID       string     `json:"user_id" gorm:"index"`
	CreatorName  string     `json:"creator_name" gorm:"-"`
	CreatorPhoto string     `json:"creator_photo" gorm:"-"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}
