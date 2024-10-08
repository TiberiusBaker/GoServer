package routes

import (
	"context"
	"net/http"

	"github.com/TiberiusBaker/GoServer/pkg/models"
	"github.com/TiberiusBaker/GoServer/pkg/utils"
)

func GameBodyCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		game := &models.Game{}
		err := utils.ParseBody(r, game)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return 
		}
		ctx := context.WithValue(r.Context(), models.GameKey, game)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ConsoleBodyCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		console := &models.Console{}
		err := utils.ParseBody(r, console)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		ctx := context.WithValue(r.Context(), models.ConsoleKey, console)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GameIdCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gameId, err := utils.ParseId(r, models.GameIdKey)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		ctx := context.WithValue(r.Context(), models.GameIdKey, gameId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ConsoleIdCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		consoleId, err := utils.ParseId(r, models.ConsoleIdKey)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		ctx := context.WithValue(r.Context(), models.ConsoleIdKey, consoleId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}