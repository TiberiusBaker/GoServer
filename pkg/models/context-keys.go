package models

import "context"

type contextKey string

var GameKey = contextKey("game")

func GetGame(ctx context.Context) *Game {
	return ctx.Value(GameKey).(*Game)
}