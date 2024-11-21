package controllers

import (
	"net/http"

	"github.com/TiberiusBaker/GoServer/pkg/models"
	"github.com/TiberiusBaker/GoServer/pkg/utils"
)

var CreateConsole = utils.JsonReturn(func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	console := models.GetFromContext[*models.Console](r.Context(), models.ConsoleKey)
	return console.CreateConsole(), nil
}, http.StatusCreated)

var GetConsoleFromId = utils.JsonReturn(func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return getConsole(r)
}, http.StatusOK)

func getConsole(r *http.Request) (interface{}, error) {
	consoleId := models.GetFromContext[string](r.Context(), models.ConsoleIdKey)
	return models.GetFromId(consoleId, &models.Console{})
}

var DeleteConsole = utils.JsonReturn(func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	consoleId := models.GetFromContext[string](r.Context(), models.ConsoleIdKey)
	return models.Delete(consoleId, &models.Console{}) 
}, http.StatusOK)

var AddGameRelation = utils.JsonReturn(func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	consoleId := models.GetFromContext[string](r.Context(), models.ConsoleIdKey)
	console := &models.Console{}
	if _, err := models.GetFromId(consoleId, &console); err != nil {
		return nil, err
	}
	game := models.GetFromContext[*models.Game](r.Context(), models.GameKey)
	return console.AddGame(game.ID)
}, http.StatusOK)

