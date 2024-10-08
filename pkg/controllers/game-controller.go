package controllers

import (
	"net/http"

	"github.com/TiberiusBaker/GoServer/pkg/models"
	"github.com/TiberiusBaker/GoServer/pkg/utils"
)

var CreateGame = utils.JsonReturn(func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	game := models.GetFromContext[*models.Game](r.Context(), models.GameKey)
	return game.CreateGame(), nil
})

var GetGameFromId = utils.JsonReturn(func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	game := models.Game{}
	gameId := models.GetFromContext[string](r.Context(), models.GameIdKey)
	return game.GetGameFromId(gameId)
})

// var addConsoleRelation = utils.JsonReturn(func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
// 	// TODO: Continue
// 	ctx := r.Context()
// 	models.GetGameKey()
// })