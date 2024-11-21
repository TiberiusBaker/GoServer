package routes

import (
	"github.com/TiberiusBaker/GoServer/pkg/controllers"
	"github.com/TiberiusBaker/GoServer/pkg/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var RegisterRoutes = func(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.AllowContentType("application/json"))

	router.Route("/game", func(r chi.Router) {
		r.With(GameBodyCtx).Post("/", controllers.CreateGame)

		r.Route("/{"+string(models.GameIdKey)+"}", func(r chi.Router) {
			r.Use(GameIdCtx)
			r.Get("/", controllers.GetGameFromId)
			r.Delete("/", controllers.DeleteGame)

			r.Route("/consoles", func(r chi.Router) {
				r.Get("/", controllers.GetGameConsoles)
				r.With(ConsoleBodyCtx).Post("/", controllers.AddConsoleRelation)
				r.With(ConsoleBodyCtx).Delete("/", controllers.DeleteConsoleRelation)
			})
		})
	})

	router.Route("/console", func(r chi.Router) {
		r.With(ConsoleBodyCtx).Post("/", controllers.CreateConsole)

		r.Route("/{"+string(models.ConsoleIdKey)+"}", func(r chi.Router) {
			r.Use(ConsoleIdCtx)
			r.Get("/", controllers.GetConsoleFromId)
			r.Delete("/", controllers.DeleteConsole)

			r.Route("/games", func(r chi.Router) {
				//r.Get("/", controllers.GetConsoleGames)
				r.With(GameBodyCtx).Post("/", controllers.AddGameRelation)
				//r.With(GameBodyCtx).Delete("/", controllers.DeleteGameRelation)
			})

		})
	})
}
