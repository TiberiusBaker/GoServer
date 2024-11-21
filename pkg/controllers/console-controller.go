package controllers

import (
	"net/http"

	"github.com/TiberiusBaker/GoServer/pkg/models"
	"github.com/TiberiusBaker/GoServer/pkg/utils"
)

//------------------PATHS-------------------------

var CreateConsole = utils.JsonReturn(func(r *http.Request) (interface{}, error) {
	return ExtractConsole(r).CreateConsole()
}, http.StatusCreated)

var DeleteConsole = utils.JsonReturn(func(r *http.Request) (interface{}, error) {
	console := &models.Console{}
	return console.Delete(ExtractConsoleId(r))
}, http.StatusOK)

var GetConsoleFromId = utils.JsonReturn(func(r *http.Request) (interface{}, error) {
	console := &models.Console{}
	return console.GetFromId(ExtractConsoleId(r))
}, http.StatusOK)

var GetConsoleGames = utils.JsonReturn(func(r *http.Request) (interface{}, error) {
	console := models.Console{}
	if _, err := console.GetFromId(ExtractConsoleId(r)); err != nil {
		return nil, err
	}
	return console.GetGames()
}, http.StatusOK)

var AddGameRelation = utils.JsonReturn(func(r *http.Request) (interface{}, error) {
	console := &models.Console{}
	if _, err := console.GetFromId(ExtractConsoleId(r)); err != nil {
		return nil, err
	}
	game := ExtractGame(r)
	return console.AddGameRel(game.ID)
}, http.StatusOK)

var DeleteGameRelation = utils.JsonReturn(func(r *http.Request) (interface{}, error) {
	console := &models.Console{}
	if _, err := console.GetFromId(ExtractConsoleId(r)); err != nil {
		return nil, err
	}
	game := ExtractGame(r)
	return console.DeleteGameRel(game.ID)
}, http.StatusOK)

//-----------------HELPERS-----------------------------
func ExtractConsole(r *http.Request) (*models.Console) {
    return models.GetFromContext[*models.Console](r.Context(), models.ConsoleKey)
}

func ExtractConsoleId(r *http.Request) (string) {
	return models.GetFromContext[string](r.Context(), models.ConsoleIdKey)
}