package main

import (
	"pong2/audio"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	// Frame counter.
	Frames uint64

	// Instances.

	Ball   *Ball
	Player *Paddle
	CPU    *CPU

	// Text fields.
	Texts struct {
		Texts []Drawer

		CPUScore    Text
		PlayerScore Text

		Pause     TextVPtr
		Muted     TextVPtr
		Wait4Play TextVPtr
	}

	Options struct {
		Paused       bool
		Waiting4Play bool
		Muted        *bool // WARN: I think this should not be a pointer, but it can still be passed up.
	}

	GUI struct {
		grid [Width / 200][Height / 50]rl.Rectangle

		Visible bool

		Play     bool
		Settings bool
		Mute     *bool
	}

	Drawers  []Drawer
	Updaters []Updater
}

func (g *Game) ResetToDefaultPositions() {
	cpuScore := g.CPU.Score
	playerScore := g.Player.Score

	g.ResetInstances()

	g.CPU.Score = cpuScore
	g.Player.Score = playerScore

	g.ResetTextFields()
}

func (g *Game) ResetInstances() {
	g.initOptions()

	g.Ball = NewBall()
	g.Player = NewPaddle()
	g.CPU = NewCPU(&g.Ball.Y)

	g.Drawers = nil
	g.Updaters = nil

	g.appendUpdateDrawers(
		g.Ball,
		g.Player,
		g.CPU,
	)
}

func (g *Game) ResetTextFields() {
	g.Texts.Texts = nil
	g.initTextFields()
}

func (g *Game) ResetToDefaultState() {
	g.ResetInstances()
	g.ResetTextFields()
}

func (g *Game) Run() {
	rl.InitWindow(Width, Height, "Pong")
	rl.InitAudioDevice()
	defer rl.CloseWindow()
	defer rl.CloseAudioDevice()

	rl.SetTargetFPS(60)

	audio.Load()
	g.initTextFields()
	g.initGUI()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		g.Update()
		g.Draw()

		rl.EndDrawing()
	}
}
