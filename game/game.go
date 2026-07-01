package game

import (
	"image/color"
)

type app struct {
	window Page
	fps    int32
	gameBG color.RGBA
}

func New(Window Page,
	Fps int32,
	GameBG color.RGBA) *app {
	return &app{
		window: Window,
		fps:    Fps,
		gameBG: GameBG,
	}
}
