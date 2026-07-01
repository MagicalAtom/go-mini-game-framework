package game

import (
	event "game/game/keys"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Page struct {
	Width  int32
	Height int32
	Title  WindowInformation
}
type WindowInformation struct {
	Text string
}

type Game struct {
	Update func(key string)
	Start  func()
}

func (a *app) Open(game Game) {
	rl.InitWindow(a.window.Width, a.window.Height, a.window.Title.Text)
	defer rl.CloseWindow()
	rl.SetTargetFPS(a.fps)
	game.Start()
	for !rl.WindowShouldClose() {
		key := event.Listen()
		rl.BeginDrawing()
		game.Update(key)
		rl.ClearBackground(a.gameBG)
		rl.EndDrawing()
	}
}
