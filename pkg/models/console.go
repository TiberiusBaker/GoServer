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

func (console *Console) getGameAssociation() *gorm.Association {
	return config.GetDB().Model(&console).Association("Games")
}

func (console *Console) doGameAssociationAction(gameId interface{}, action AssociationAction) (*Game, error) {
	game := &Game{}
	if _, err := GetFromId(gameId, game); err != nil {
		return nil, err
	}
	if err := action(console.getGameAssociation(), game); err != nil {
		return nil, err
	}
	return game, nil
}


func (console *Console) AddGame (gameId interface{}) (*Game, error) {
	return console.doGameAssociationAction(gameId, func(association *gorm.Association, game interface{}) error {
		return association.Append(game)
	})
}
