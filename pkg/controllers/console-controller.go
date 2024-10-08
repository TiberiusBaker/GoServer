package controllers

import (
	"net/http"

	"github.com/TiberiusBaker/GoServer/pkg/models"
	"github.com/TiberiusBaker/GoServer/pkg/utils"
)

var CreateConsole = utils.JsonReturn(func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	console := models.GetFromContext[*models.Console](r.Context(), models.ConsoleKey)
	return console.CreateConsole(), nil
})

var GetConsoleFromId = utils.JsonReturn(func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	console := &models.Console{}
	consoleId := models.GetFromContext[string](r.Context(), models.ConsoleIdKey)
	return console.GetConsoleFromId(consoleId)
})