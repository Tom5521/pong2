package main

import (
	"pong2/audio"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	// Frame counter.
	Frames long

	// Instances.

	Ball    Ball
	Player  Paddle
	Player2 Paddle2

	// Text fields.
	Texts struct {
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
	g.Texts.PlayerScore.Draw()
	g.Texts.CPUScore.Draw()
	g.Texts.Wait4Play.Draw()
	g.Texts.Muted.Draw()
	g.Texts.Pause.Draw()

	g.Player2.Draw()
	g.Player.Draw()
	g.Ball.Draw()
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
