package models

import (
	"fmt"

	"github.com/TiberiusBaker/GoServer/pkg/config"
)

type IdHaver interface {
	SetID(uint) 
}

func (c *Console) SetID(id uint) {
	c.ID = id
}

func (game *Game) SetID(id uint) {
	game.ID = id
}

func Delete(id string, model interface{}) (interface{}, error) {
	db := config.GetDB()
	if err := db.Delete(&model, id).Error; err != nil {
		return model, err
	}
	return model, nil
}

func GetFromId(id interface{}, model interface{}) (interface{}, error) {
	db := config.GetDB()
	if err := db.Where("ID = ?", id).Limit(1).First(&model).Error; err != nil {
		return nil, fmt.Errorf("No item found for id: %s, %w", id, err)
	}
	return model, nil
}
