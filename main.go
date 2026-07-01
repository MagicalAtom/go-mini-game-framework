package main

import (
	"fmt"
	"image/color"
	"math/rand"

	game "game/game"
	event "game/game/keys"
	"game/game/system"
	"game/game/trigger"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Point struct {
	X int32
	Y int32
}

type SnakeProperty struct {
	system.CommonProperty
	Body      []Point
	Direction string
}

func (s *SnakeProperty) GetBody() []Point      { return s.Body }
func (s *SnakeProperty) SetBody(b []Point)     { s.Body = b }
func (s *SnakeProperty) GetDirection() string  { return s.Direction }
func (s *SnakeProperty) SetDirection(d string) { s.Direction = d }

type FoodProperty struct {
	system.CommonProperty
	X int32
	Y int32
}

func (f *FoodProperty) GetX() int32  { return f.X }
func (f *FoodProperty) GetY() int32  { return f.Y }
func (f *FoodProperty) SetX(x int32) { f.X = x }
func (f *FoodProperty) SetY(y int32) { f.Y = y }

type StatusProperty struct {
	system.CommonProperty
	Score     int32
	GameOver  bool
	TickTimer float32
	TickRate  float32
}

func (g *StatusProperty) GetScore() int32        { return g.Score }
func (g *StatusProperty) SetScore(s int32)       { g.Score = s }
func (g *StatusProperty) IsGameOver() bool       { return g.GameOver }
func (g *StatusProperty) SetGameOver(o bool)     { g.GameOver = o }
func (g *StatusProperty) GetTickTimer() float32  { return g.TickTimer }
func (g *StatusProperty) SetTickTimer(t float32) { g.TickTimer = t }
func (g *StatusProperty) GetTickRate() float32   { return g.TickRate }
func (g *StatusProperty) SetTickRate(r float32)  { g.TickRate = r }

func main() {
	world := game.New(
		game.Page{
			Width:  800,
			Height: 600,
			Title: game.WindowInformation{
				Text: "Venuro Engine - Pure Trigger Snake",
			},
		},
		60,
		color.RGBA{R: 12, G: 18, B: 12, A: 255},
	)

	event.AddEvent(event.AddKey("move_up", event.W))
	event.AddEvent(event.AddKey("move_down", event.S))
	event.AddEvent(event.AddKey("move_left", event.A))
	event.AddEvent(event.AddKey("move_right", event.D))

	mSnake := system.New(system.ModelName{Name: "Snake"})
	mSnake.AddProperty(&SnakeProperty{
		CommonProperty: system.CommonProperty{Name: "snake_data", ModelName: "Snake"},
		Body: []Point{
			{X: 20, Y: 15},
			{X: 19, Y: 15},
			{X: 18, Y: 15},
		},
		Direction: "move_right",
	})

	mFood := system.New(system.ModelName{Name: "Food"})
	mFood.AddProperty(&FoodProperty{
		CommonProperty: system.CommonProperty{Name: "food_data", ModelName: "Food"},
		X:              10,
		Y:              10,
	})

	mStatus := system.New(system.ModelName{Name: "Status"})
	mStatus.AddProperty(&StatusProperty{
		CommonProperty: system.CommonProperty{Name: "status_data", ModelName: "Status"},
		Score:          0,
		GameOver:       false,
		TickTimer:      0.0,
		TickRate:       0.10,
	})

	hSnake, _ := system.Make(system.Hook[*SnakeProperty]{Model: mSnake, PropertyName: "snake_data"})
	hFood, _ := system.Make(system.Hook[*FoodProperty]{Model: mFood, PropertyName: "food_data"})
	hStatus, _ := system.Make(system.Hook[*StatusProperty]{Model: mStatus, PropertyName: "status_data"})

	trigger.T.AddTrigger(trigger.CreateTrigger{
		EventName: "move_up",
		Handler: func() {
			if hSnake.Property.GetDirection() != "move_down" {
				hSnake.Property.SetDirection("move_up")
			}
		},
	})

	trigger.T.AddTrigger(trigger.CreateTrigger{
		EventName: "move_down",
		Handler: func() {
			if hSnake.Property.GetDirection() != "move_up" {
				hSnake.Property.SetDirection("move_down")
			}
		},
	})

	trigger.T.AddTrigger(trigger.CreateTrigger{
		EventName: "move_left",
		Handler: func() {
			if hSnake.Property.GetDirection() != "move_right" {
				hSnake.Property.SetDirection("move_left")
			}
		},
	})

	trigger.T.AddTrigger(trigger.CreateTrigger{
		EventName: "move_right",
		Handler: func() {
			if hSnake.Property.GetDirection() != "move_left" {
				hSnake.Property.SetDirection("move_right")
			}
		},
	})

	trigger.T.AddTrigger(trigger.CreateTrigger{
		EventName: "game_physics",
		Handler: func() {
			if hStatus.Property.IsGameOver() {
				return
			}

			hStatus.Property.SetTickTimer(hStatus.Property.GetTickTimer() + rl.GetFrameTime())
			if hStatus.Property.GetTickTimer() < hStatus.Property.GetTickRate() {
				return
			}
			hStatus.Property.SetTickTimer(0)

			body := hSnake.Property.GetBody()
			head := body[0]

			switch hSnake.Property.GetDirection() {
			case "move_up":
				head.Y--
			case "move_down":
				head.Y++
			case "move_left":
				head.X--
			case "move_right":
				head.X++
			}

			if head.X < 0 || head.X >= 40 || head.Y < 0 || head.Y >= 30 {
				hStatus.Property.SetGameOver(true)
				return
			}

			for _, segment := range body {
				if head.X == segment.X && head.Y == segment.Y {
					hStatus.Property.SetGameOver(true)
					return
				}
			}

			newBody := append([]Point{head}, body...)

			if head.X == hFood.Property.GetX() && head.Y == hFood.Property.GetY() {
				hStatus.Property.SetScore(hStatus.Property.GetScore() + 10)
				hFood.Property.SetX(rand.Int31n(38) + 1)
				hFood.Property.SetY(rand.Int31n(28) + 1)
			} else {
				newBody = newBody[:len(newBody)-1]
			}

			hSnake.Property.SetBody(newBody)
		},
	})

	trigger.T.AddTrigger(trigger.CreateTrigger{
		EventName: "game_render",
		Handler: func() {
			if hStatus.Property.IsGameOver() {
				rl.DrawText("GAME OVER", 270, 240, 50, rl.Red)
				rl.DrawText(fmt.Sprintf("Final Score: %d", hStatus.Property.GetScore()), 330, 310, 24, rl.LightGray)
				return
			}

			rl.DrawRectangle(hFood.Property.GetX()*20, hFood.Property.GetY()*20, 18, 18, rl.Red)

			body := hSnake.Property.GetBody()
			for i, segment := range body {
				c := rl.Lime
				if i == 0 {
					c = rl.Green
				}
				rl.DrawRectangle(segment.X*20, segment.Y*20, 18, 18, c)
			}

			rl.DrawText(fmt.Sprintf("SCORE: %d", hStatus.Property.GetScore()), 20, 20, 20, rl.RayWhite)
		},
	})

	world.Open(game.Game{
		Start: func() {},
		Update: func(key string) {
			trigger.T.Listen(key)
			trigger.T.Listen("game_physics")
			trigger.T.Listen("game_render")
		},
	})
}
