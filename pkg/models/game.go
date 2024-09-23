package models

import (
	"fmt"

	"github.com/TiberiusBaker/GoServer/pkg/config"
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Name        string `gorm:"size:255"`
	Description string
	Publisher   string `gorm:"size:512"`
	Image       string `gorm:"size:512"`
	NumPlayers  uint64
	Consoles    []Console `gorm:"many2many:game_consoles;"`
}

func init() {
	db := config.GetDB()
	db.AutoMigrate(&Game{})
}

func (g *Game) CreateGame() *Game {
	db := config.GetDB()
	result := db.Create(&g)

	if result.Error != nil {
        fmt.Println("Error inserting game:", result.Error)
    } else {
        fmt.Println("Game inserted successfully with ID:", g.ID)
    }
	return g
}