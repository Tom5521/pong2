package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Game struct {
	Frames long

	Ball     Ball
	Drawers  []Drawer
	Updaters []Updater
}

func NewGame() Game {
	var g Game

	g.Ball = NewBall()

	g.Drawers = append(g.Drawers, &g.Ball)
	g.Updaters = append(g.Updaters, &g.Ball)

	return g
}

func (g *Game) Update() {
	for _, u := range g.Updaters {
		g.Frames++
		u.Update()
	}
}

func (g *Game) Draw() {
	for _, d := range g.Drawers {
		d.Draw()
	}
}

func (g *Game) Run() {
	rl.InitWindow(Width, Height, "Pong")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		g.Update()
		g.Draw()
		rl.EndDrawing()
	}
}
