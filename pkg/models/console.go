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

func (console *Console) Delete(id interface{}) (*Console, error) {
	db := config.GetDB()
	if err := db.Delete(&console, id).Error; err != nil {
		return nil, err
	}
	return console, nil
}

func (console *Console) GetFromId(id interface{}) (*Console, error) {
	db := config.GetDB()
	if err := db.Where("ID = ?", id).Limit(1).First(&console).Error; err != nil {
		return nil, fmt.Errorf("No item found for id: %s, %w", id, err)
	}
	return console, nil
}

func (c *Console) CreateConsole() (*Console, error) {
	db := config.GetDB()
	if err := db.Create(&c).Error; err != nil {
		return nil, err
	}
	return c, nil
}

func (console *Console) doGameAssociationAction(gameId interface{}, action AssociationAction) (*Game, error) {
	gameItem := Game{}
	game, err := gameItem.GetFromId(gameId)
	if err != nil {
		return nil, err
	}
	if err := action(console.getGameAssociation(), game); err != nil {
		return nil, err
	}
	return game, nil
}

func (console *Console) AddGameRel(gameId interface{}) (*Game, error) {
	return console.doGameAssociationAction(gameId, func(association *gorm.Association, game interface{}) error {
		return association.Append(game)
	})
}

func (console *Console) DeleteGameRel(gameId interface{}) (*Game, error) {
	return console.doGameAssociationAction(gameId, func(association *gorm.Association, game interface{}) error {
		return association.Delete(game)
	})
}

func (console *Console) GetGames() (*[]Game, error) {
	var games []Game
	if err := console.getGameAssociation().Find(&games); err != nil {
		return nil, err
	}
	return &games, nil
}

//-----------------------HELPERS----------------------
func (console *Console) getGameAssociation() *gorm.Association {
	return config.GetDB().Model(&console).Association("Games")
}