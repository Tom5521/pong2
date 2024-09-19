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
					int32(r.X),
					int32(r.Y),
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

	rl.DrawText("PONG", int32(ui.grid[2][2].X), int32(ui.grid[2][2].Y), 140, rl.White)
	rl.DrawText("By Tom5521", int32(ui.grid[2][5].X), int32(ui.grid[2][5].Y), 30, rl.Gray)

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

func (g *Game) Draw() {
	if devel {
		rl.DrawLine(
			0,
			Height/2,
			Width,
			Height/2,
			rl.DarkGray,
		)

		rl.DrawFPS(0, 0)
	}

	if g.GUI.Visible {
		g.Ball.Draw()
		g.RenderGUI()
		return
	}

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

	for _, d := range g.Drawers {
		d.Draw()
	}
}
