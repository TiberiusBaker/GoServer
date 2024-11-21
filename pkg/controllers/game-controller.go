package controllers

import (
	"net/http"

	"github.com/TiberiusBaker/GoServer/pkg/models"
	"github.com/TiberiusBaker/GoServer/pkg/utils"
)

var CreateGame = utils.JsonReturn(func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	game := models.GetFromContext[*models.Game](r.Context(), models.GameKey)
	return game.CreateGame()
}, http.StatusCreated)

var DeleteGame = utils.JsonReturn(func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	gameId := models.GetFromContext[string](r.Context(), models.GameIdKey)
	return models.Delete(gameId, &models.Game{})
}, http.StatusOK)

var GetGameFromId = utils.JsonReturn(func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	gameId := models.GetFromContext[string](r.Context(), models.GameIdKey)
	return models.GetFromId(gameId, &models.Game{})
}, http.StatusOK)

var GetGameConsoles = utils.JsonReturn(func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	game := models.Game{}
	gameId := models.GetFromContext[string](r.Context(), models.GameIdKey)
	return game.GetConsoles(gameId)
}, http.StatusOK)

var AddConsoleRelation = utils.JsonReturn(func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	game := models.Game{}
	gameId := models.GetFromContext[string](r.Context(), models.GameIdKey)
	if _, err := models.GetFromId(gameId, &game); err != nil {
		return nil, err
	}
	console := models.GetFromContext[*models.Console](r.Context(), models.ConsoleKey)
	return game.AddConsole(console.ID)
}, http.StatusOK)

var DeleteConsoleRelation = utils.JsonReturn(func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	game := models.Game{}
	gameId := models.GetFromContext[string](r.Context(), models.GameIdKey)
	if _, err := models.GetFromId(gameId, &game); err != nil {
		return nil, err
	}
	console := models.GetFromContext[*models.Console](r.Context(), models.ConsoleKey)
	return game.DeleteConsole(console.ID)
}, http.StatusOK)

