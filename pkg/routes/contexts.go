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