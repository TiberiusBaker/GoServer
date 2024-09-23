package models

import "gorm.io/gorm"

type Console struct {
	gorm.Model
	Name           string `gorm:"size:255"`
	Image          string `gorm:"size:512"`
	PlayerBaseSize uint64
	Games          []Game `gorm:"many2many:game_consoles;"`
}
