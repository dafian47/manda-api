package model

type MandaMarriageStatus struct {
	ID  int `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Label string `json:"label" gorm:"type:text;not null"`
}

type MandaWork struct {
	ID  int `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Label string `json:"label" gorm:"type:text;not null"`
}

type MandaMajor struct {
	Code  string `json:"code" gorm:"primary_key;unique"`
	Label string `json:"label" gorm:"type:text;not null"`
}

type MandaRole struct {
	Code  string `json:"code" gorm:"primary_key;unique"`
	Label string `json:"label" gorm:"type:text;not null"`
}

type MandaUserStatus struct {
	Code  string `json:"code" gorm:"primary_key;unique"`
	Label string `json:"label" gorm:"type:text;not null"`
}

type MandaThreadStatus struct {
	Code  string `json:"code" gorm:"primary_key;unique"`
	Label string `json:"label" gorm:"type:text;not null"`
}
