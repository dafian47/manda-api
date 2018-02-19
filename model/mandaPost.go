package model

import "time"

type MandaPost struct {
	ID           string     `json:"id" gorm:"primary_key;unique"`
	Title        string     `json:"name" binding:"required"`
	Content      string     `json:"content" binding:"required" gorm:"type:text"`
	Status       string     `json:"status" binding:"required" gorm:"index"`
	UrlImage     string     `json:"url_image"`
	CountViewer  int        `json:"count_viewer"`
	CountComment int        `json:"count_comment" gorm:"-"`
	ChannelID    string     `json:"channel_id" gorm:"index"`
	ChannelName  string     `json:"channel_name" gorm:"-"`
	ChannelIcon  string     `json:"channel_icon" gorm:"-"`
	UserID       string     `json:"user_id" gorm:"index"`
	CreatorName  string     `json:"creator_name" gorm:"-"`
	CreatorPhoto string     `json:"creator_photo" gorm:"-"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}
