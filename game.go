package main

import (
	"pong2/audio"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	// Frame counter.
	Frames long

	// Instances.

	Ball   Ball
	Paddle Paddle
	CPU    CPU

	// Text fields.
	Texts struct {
		Drawers []Text

		CPUScore    Text
		PlayerScore Text

		Pause     Text
		Muted     Text
		Wait4Play Text
	}

	Options struct {
		Paused       bool
		Waiting4Play bool
		Muted        *bool
	}

	// Update/Draw slices.

	Drawers  []Drawer
	Updaters []Updater
}

func NewGame() Game {
	var g Game

	g.initOptions()

	g.Ball = NewBall()
	g.Paddle = NewPaddle()
	g.CPU = NewCPU(&g.Ball.Y)

	g.addInstance(&g.Ball)
	g.addInstance(&g.Paddle)
	g.addInstance(&g.CPU)

	return g
}

func (g *Game) initOptions() {
	o := &g.Options
	o.Muted = &audio.Mute
	o.Waiting4Play = true
}

func (g *Game) initTextFields() {
	t := &g.Texts

	// Player Score.
	t.PlayerScore = NewText("0", 50, rl.Gray, rl.Vector2{})
	t.PlayerScore.Y = Height/2 - t.PlayerScore.Measure().Y/2
	t.PlayerScore.X = Width/2 + t.PlayerScore.Measure().X*3

	// CPU Score.
	t.CPUScore = NewText("0", 50, rl.Gray, rl.Vector2{})
	t.CPUScore.Y = t.PlayerScore.Y
	t.CPUScore.X = Width/2 - t.CPUScore.Measure().X*4

	// Wait for play.
	t.Wait4Play = NewText("Press [SPACE] to start", 50, rl.White, rl.Vector2{})
	t.Wait4Play.Y = Height/3 - t.Wait4Play.Measure().Y/2
	t.Wait4Play.X = Width/2 - t.Wait4Play.Measure().X/2
	t.Wait4Play.Visible = false

	// Muted.
	t.Muted = NewText("MUTED", 50, rl.White, rl.Vector2{})
	t.Muted.Visible = false

	// Paused.
	t.Pause = NewText("PAUSED", 50, rl.White, rl.Vector2{})
	t.Pause.Y = Height/3 - t.Pause.Measure().Y/2
	t.Pause.Visible = false

	t.Drawers = append(
		t.Drawers,
		// Scoring.
		t.PlayerScore,
		t.CPUScore,
		// Options.
		t.Wait4Play,
		t.Muted,
		t.Pause,
	)
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

func (g *Game) Draw() {
	rl.DrawLine(
		Width/2,
		0,
		Width/2,
		Height,
		rl.DarkGray,
	)

	// Center line for debugging.
	rl.DrawLine(
		0,
		Height/2,
		Width,
		Height/2,
		rl.Red,
	)

	// Draw texts.
	for _, d := range g.Texts.Drawers {
		d.Draw()
	}
	// Draw instances.
	for _, d := range g.Drawers {
		d.Draw()
	}
}

func (g *Game) Run() {
	rl.InitWindow(Width, Height, "Pong")
	rl.InitAudioDevice()
	defer rl.CloseWindow()
	defer rl.CloseAudioDevice()

	rl.SetTargetFPS(60)

	audio.Load()
	g.initTextFields()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		g.Update()
		g.Draw()

		rl.EndDrawing()
	}
}
