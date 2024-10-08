package models

import (
	"fmt"

	"github.com/TiberiusBaker/GoServer/pkg/config"
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Name        string `gorm:"size:255" json:"name"`
	Description string `json:"description"`
	Publisher   string `gorm:"size:512" json:"publisher"`
	Image       string `gorm:"size:512" json:"image"`
	NumPlayers  uint64 `json:"numPlayers"`
	Consoles    []Console `gorm:"many2many:game_consoles;" json:"consoles"`
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

func (g *Game) GetGameFromId(gameId string) (*Game, error) {
	db := config.GetDB()
	if err := db.Where("ID = ?", gameId).First(&g).Error; err != nil {
		return nil, fmt.Errorf("game")
	}
	return g, nil
}

func (g *Game) AddConsole(consoleName string) error {
	db := config.GetDB()

	console := Console{}
	if err := db.Where("name = ?", consoleName).First(&console).Error; err != nil {
        return fmt.Errorf("console %s not found: %w", consoleName, err)
    }
	if err := db.Model(&g).Association("Consoles").Append(&console); err != nil {
        return fmt.Errorf("failed to add console to game: %w", err)
    }

	return nil
}