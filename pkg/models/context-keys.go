package models

import "context"

type ContextKey string

var GameKey = ContextKey("game")
var GameIdKey = ContextKey("gameId")

var ConsoleKey = ContextKey("console")
var ConsoleIdKey = ContextKey("consoleId")

func GetFromContext[T any](ctx context.Context, key ContextKey) T {
	return ctx.Value(key).(T)
}
