package models

import (
	"fmt"

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

func (game *Game) Delete(id interface{}) (*Game, error) {
	db := config.GetDB()
	if err := db.Delete(&game, id).Error; err != nil {
		return nil, err
	}
	return game, nil
}

func (game *Game) GetFromId(id interface{}) (*Game, error) {
	db := config.GetDB()
	if err := db.Where("ID = ?", id).Limit(1).First(&game).Error; err != nil {
		return nil, fmt.Errorf("No item found for id: %s, %w", id, err)
	}
	return game, nil
}

func (game *Game) doConsoleAssociationAction(consoleId interface{}, action AssociationAction) (*Game, error) {
	console := &Console{}
	if _, err := console.GetFromId(consoleId); err != nil {
		return nil, err
	}
	if err := action(game.getConsoleAssociation(), console); err != nil {
		return nil, err
	}
	return game, nil
}


func (game *Game) AddConsoleRel(consoleId interface{}) (*Game, error) {
	return game.doConsoleAssociationAction(consoleId, func(association *gorm.Association, console interface{}) error {
		return association.Append(console)
	})
}

func (game *Game) DeleteConsoleRel(consoleId interface{}) (*Game, error) {
	return game.doConsoleAssociationAction(consoleId, func(association *gorm.Association, console interface{}) error {
		return association.Delete(console)
	})
}

func (game *Game) GetConsoles() ([]Console, error) {
	var consoles []Console 
	if err := game.getConsoleAssociation().Find(&consoles); err != nil {
		return nil, err
	}
	return consoles, nil
}


//------------------HELPERS--------------------
func (game *Game) getConsoleAssociation() *gorm.Association {
	return config.GetDB().Model(&game).Association("Consoles")
}

type AssociationAction func(association *gorm.Association, model interface{}) error