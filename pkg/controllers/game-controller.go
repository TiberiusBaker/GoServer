package controllers

import (
	"net/http"

	"github.com/TiberiusBaker/GoServer/pkg/models"
	"github.com/TiberiusBaker/GoServer/pkg/utils"
)

//------------------PATHS-------------------------
var CreateGame = utils.JsonReturn(func(r *http.Request) (interface{}, error) {
	return ExtractGame(r).CreateGame()
}, http.StatusCreated)

var DeleteGame = utils.JsonReturn(func(r *http.Request) (interface{}, error) {
	game := models.Game{}
	return game.Delete(ExtractGameId(r))
}, http.StatusOK)

var GetGameFromId = utils.JsonReturn(func(r *http.Request) (interface{}, error) {
	game := &models.Game{}
	return game.GetFromId(ExtractGameId(r))
}, http.StatusOK)

var GetGameConsoles = utils.JsonReturn(func(r *http.Request) (interface{}, error) {
	game := models.Game{}
	if _, err := game.GetFromId(ExtractGameId(r)); err != nil {
		return nil, err
	}
	return game.GetConsoles()
}, http.StatusOK)

var AddConsoleRelation = utils.JsonReturn(func(r *http.Request) (interface{}, error) {
	game := &models.Game{}
	if _, err := game.GetFromId(ExtractGameId(r)); err != nil {
		return nil, err
	}
	console := ExtractConsole(r)
	return game.AddConsoleRel(console.ID)
}, http.StatusOK)

var DeleteConsoleRelation = utils.JsonReturn(func(r *http.Request) (interface{}, error) {
	game := &models.Game{}
	if _, err := game.GetFromId(ExtractGameId(r)); err != nil {
		return nil, err
	}
	console := ExtractConsole(r)
	return game.DeleteConsoleRel(console.ID)
}, http.StatusOK)

//-----------------HELPERS-----------------------------
func ExtractGameId(r *http.Request) string {
	return models.GetFromContext[string](r.Context(), models.GameIdKey)
}

func ExtractGame(r *http.Request) *models.Game {
	return models.GetFromContext[*models.Game](r.Context(), models.GameKey)
}

