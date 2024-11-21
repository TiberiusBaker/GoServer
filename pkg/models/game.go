package models

import (
	"github.com/TiberiusBaker/GoServer/pkg/config"
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Name        string    `gorm:"size:255" json:"name"`
	Description string    `json:"description"`
	Publisher   string    `gorm:"size:512" json:"publisher"`
	Image       string    `gorm:"size:512" json:"image"`
	NumPlayers  uint64    `json:"numPlayers"`
	Consoles    []Console `gorm:"many2many:game_consoles;" json:"consoles"`
}

func init() {
	db := config.GetDB()
	db.AutoMigrate(&Game{})
}

func (game *Game) CreateGame() (*Game, error) {
	db := config.GetDB()
	if err := db.Create(&game).Error; err != nil {
		return nil, err
	}
	return game, nil
}

func (game *Game) getConsoleAssociation() *gorm.Association {
	return config.GetDB().Model(&game).Association("Consoles")
}

type AssociationAction func(*gorm.Association, interface{}) error

func (game *Game) doConsoleAssociationAction(consoleId interface{}, action AssociationAction) (*Game, error) {
	console := &Console{}
	if _, err := GetFromId(consoleId, console); err != nil {
		return nil, err
	}
	if err := action(game.getConsoleAssociation(), console); err != nil {
		return nil, err
	}
	return game, nil
}


func (game *Game) AddConsole(consoleId interface{}) (*Game, error) {
	return game.doConsoleAssociationAction(consoleId, func(association *gorm.Association, console interface{}) error {
		return association.Append(console)
	})
}

func (game *Game) DeleteConsole(consoleId interface{}) (*Game, error) {
	return game.doConsoleAssociationAction(consoleId, func(association *gorm.Association, console interface{}) error {
		return association.Delete(console)
	})
}

func (game *Game) GetConsoles(gameId string) ([]Console, error) {
	if _, err := GetFromId(gameId, game); err != nil {
		return nil, err
	}
	var consoles []Console 
	if err := game.getConsoleAssociation().Find(&consoles); err != nil {
		return nil, err
	}
	return consoles, nil
}
