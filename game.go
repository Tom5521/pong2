package main

import (
	"fmt"

	"pong2/audio"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	// Frame counter.
	Frames long // TODO: Remove this.

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

	devel bool

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

func (g *Game) DrawGUI() {
	ui := &g.GUI

	if g.devel {
		for i := range ui.grid {
			for j := range ui.grid[i] {
				r := ui.grid[i][j]
				rl.DrawRectangleLinesEx(r, 1, rl.Red)

				rl.DrawText(
					fmt.Sprintf("x:%v,y:%v", i, j),
					int32(r.X),
					int32(r.Y),
					20,
					rl.White,
				)
			}
		}
	}

	ui.Play = raygui.Button(ui.grid[0][3], "Play")
	ui.Settings = raygui.Button(ui.grid[0][5], "Settings")

	muteTxt := "Mute"
	if *g.Options.Muted {
		muteTxt = "Muted"
	}
	*g.Options.Muted = raygui.Toggle(ui.grid[0][7], muteTxt, *g.Options.Muted)

	if ui.Play {
		g.Options.Waiting4Play = false
		g.GUI.Visible = false
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

	for _, t := range g.Texts.Texts {
		t.Draw()
	}

	if g.GUI.Visible {
		g.DrawGUI()
		return
	}

	for _, d := range g.Drawers {
		d.Draw()
	}
}

func (g *Game) Run() {
	rl.InitWindow(Width, Height, "Pong")
	rl.InitAudioDevice()
	defer rl.CloseWindow()
	defer rl.CloseAudioDevice()

	rl.SetTargetFPS(60) // This probably fits better in 30.

	audio.Load()
	g.initTextFields()
	g.initGUI()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		g.Update()
		g.Draw()

		rl.EndDrawing()
	}
}
