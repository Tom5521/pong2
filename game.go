package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Game struct {
	Frames long

	Ball   Ball
	Paddle Paddle
	CPU    CPU

	Drawers  []Drawer
	Updaters []Updater
}

func NewGame() Game {
	var g Game

	g.Ball = NewBall()
	g.Paddle = NewPaddle()
	g.CPU = NewCPU(&g.Ball.Y)

	g.addInstance(&g.Ball)
	g.addInstance(&g.Paddle)
	g.addInstance(&g.CPU)

	return g
}

func (g *Game) addInstance(appends ...any) {
	for _, a := range appends {
		u, ok := a.(Updater)
		if ok {
			g.Updaters = append(g.Updaters, u)
		}
		d, ok := a.(Drawer)
		if ok {
			g.Drawers = append(g.Drawers, d)
		}
	}
}

func (g *Game) Update() {
	// Update instances.
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
