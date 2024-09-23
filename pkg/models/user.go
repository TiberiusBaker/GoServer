package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string `gorm:"size:255;unique"`
	ProfileName string `gorm:"size:255"`
	Password    string `gorm:"size:255"`
	Email       string `gorm:"size:255;unique"`
	Age         uint8
}
