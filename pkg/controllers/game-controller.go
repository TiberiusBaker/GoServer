package controllers

import (
	"net/http"

	"github.com/TiberiusBaker/GoServer/pkg/models"
	"github.com/TiberiusBaker/GoServer/pkg/utils"
)

var CreateGame = utils.JsonReturn(func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	ctx := r.Context()
	game := models.GetGame(ctx)
	return game.CreateGame(), nil
})