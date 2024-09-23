package models

import "gorm.io/gorm"

type Game struct {
	gorm.Model
	Name        string `gorm:"size:255"`
	Description string
	Image       string `gorm:"size:512"`
	NumPlayers  uint64
	Consoles    []Console `gorm:"many2many:game_consoles;"`
}
