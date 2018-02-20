package model

import "time"

type MandaComment struct {
	ID           string     `json:"id" gorm:"primary_key;unique"`
	Comment      string     `json:"comment" binding:"required"`
	ThreadID     string     `json:"thread_id" gorm:"index"`
	UserID       string     `json:"user_id" gorm:"index"`
	CreatorName  string     `json:"creator_name" gorm:"-"`
	CreatorPhoto string     `json:"creator_photo" gorm:"-"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}
