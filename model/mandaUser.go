package model

import "time"

type MandaUser struct {
	ID             string     `json:"id" gorm:"primary_key;unique"`
	Username       string     `json:"username" gorm:"-"`
	Password       string     `json:"password" gorm:"-"`
	FullName       string     `json:"full_name" binding:"required"`
	NickName       string     `json:"nick_name" binding:"required"`
	Title          string     `json:"title"`
	BirthPlace     string     `json:"birth_place" binding:"required"`
	BirthDate      time.Time  `json:"birth_date"`
	Gender         string     `json:"gender" binding:"required" gorm:"index"`
	Address        string     `json:"address" binding:"required" gorm:"type:text"`
	Phone          string     `json:"phone" binding:"required"`
	Email          string     `json:"email" binding:"required"`
	MarriageStatus string     `json:"marriage_status" binding:"required" gorm:"index"`
	Generation     string     `json:"generation" binding:"required" gorm:"index"`
	Major          string     `json:"major" binding:"required" gorm:"index"`
	Work           string     `json:"work" binding:"required" gorm:"index"`
	Type           string     `json:"type" binding:"required" gorm:"index"`
	Status         string     `json:"status" binding:"required" gorm:"index"`
	UrlPhoto       string     `json:"url_photo"`
	OneSignalId    string     `json:"one_signal_id"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
}
