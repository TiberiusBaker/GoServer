package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string `gorm:"size:255;unique" json:"username"`
	ProfileName string `gorm:"size:255" json:"profileName"`
	Password    string `gorm:"size:255" json:"-"`
	Email       string `gorm:"size:255;unique" json:"email"`
	Age         uint8 `json:"age"`
}
