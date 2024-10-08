package models

import (
	"fmt"

	"github.com/TiberiusBaker/GoServer/pkg/config"
	"gorm.io/gorm"
)

type Console struct {
	gorm.Model
	Name           string `gorm:"size:255" json:"name"`
	Image          string `gorm:"size:512" json:"image"`
	PlayerBaseSize uint64 `json:"playerBaseSize"`
	Games          []Game `gorm:"many2many:game_consoles;" json:"games"`
}

func (c *Console) CreateConsole() *Console {
	db := config.GetDB()
	result := db.Create(&c)

	if result.Error != nil {
        fmt.Println("Error creating console:", result.Error)
    } else {
        fmt.Println("Game inserted successfully with ID:", c.ID)
    }
	return c
}

func (c *Console) GetConsoleFromId(consoleId string) (*Console, error) {
	db := config.GetDB()
	if err := db.Where("ID = ?", consoleId).First(&c).Error; err != nil {
		return nil, fmt.Errorf("No such console for id " + consoleId)
	}
	return c, nil
}