package routes

import (
	"github.com/TiberiusBaker/GoServer/pkg/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var RegisterRoutes = func (router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.AllowContentType("application/json"))

	router.Route("/game", func (r chi.Router) {
		r.With(GameBodyCtx).Post("/", controllers.CreateGame)
	})
}