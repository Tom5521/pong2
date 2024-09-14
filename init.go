package main

import (
	"pong2/audio"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewGame() Game {
	var g Game

	g.initOptions()

	g.Ball = NewBall()
	g.Player = NewPaddle()
	g.CPU = NewCPU(&g.Ball.Y)

	return g
}

func (g *Game) initOptions() {
	o := &g.Options
	o.Muted = &audio.Mute
	*o.Muted = true
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

	t.Texts = append(t.Texts,
		&t.PlayerScore,
		&t.CPUScore,
		&t.Wait4Play,
		&t.Muted,
		&t.Pause,
	)
}
