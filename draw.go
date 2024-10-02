package main

import (
	"fmt"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) RenderGUI() {
	ui := &g.GUI

	if devel {
		for i := range ui.grid {
			for j := range ui.grid[i] {
				r := ui.grid[i][j]
				rl.DrawRectangleLinesEx(r, 0.4, rl.Red)

				rl.DrawText(
					fmt.Sprintf("x:%v,y:%v", i, j),
					rint(r.X),
					rint(r.Y),
					20,
					rl.White,
				)
			}
		}
	}

	ui.Play = raygui.Button(ui.grid[0][3], "Play")
	if devel {
		ui.Settings = raygui.Button(ui.grid[0][5], "Settings")
	}

	rl.DrawText("PONG", rint(ui.grid[2][2].X), rint(ui.grid[2][2].Y), 140, rl.White)
	rl.DrawText("By Tom5521", rint(ui.grid[2][5].X), rint(ui.grid[2][5].Y), 30, rl.Gray)

	muteTxt := "Mute"
	if *g.Options.Muted {
		muteTxt = "Muted"
	}

	if devel {
		*g.Options.Muted = raygui.Toggle(ui.grid[0][7], muteTxt, *g.Options.Muted)
	} else {
		*g.Options.Muted = raygui.Toggle(ui.grid[0][5], muteTxt, *g.Options.Muted)
	}

	if ui.Play {
		g.Options.Waiting4Play = true
		g.GUI.Visible = false

		g.ResetInstances()
	}
}

func (g *Game) DrawBackground() {
	rl.ClearBackground(rl.DarkGreen)

	// Draw window borders
	rl.DrawRectangleLinesEx(
		rl.Rectangle{
			Width:  Width,
			Height: Height,
		},
		10,
		rl.White,
	)

	// Draw center-line
	rl.DrawLineEx(
		rl.Vector2{
			X: Width / 2,
			Y: 0,
		},
		rl.Vector2{
			X: Width / 2,
			Y: Height,
		},
		10,
		rl.White,
	)
}

func (g *Game) Draw() {
	g.DrawBackground()

	if devel {
		rl.DrawLine(
			0,
			Height/2,
			Width,
			Height/2,
			rl.DarkGray,
		)

		rl.DrawFPS(11, 11)
	}

	if g.GUI.Visible {
		g.Ball.Draw()
		g.RenderGUI()
		return
	}

	for _, t := range g.Texts.Texts {
		t.Draw()
	}

	for _, d := range g.Drawers {
		d.Draw()
	}
}
