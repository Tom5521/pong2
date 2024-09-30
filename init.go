package main

import (
	"strconv"

	"pong2/audio"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewGame() (g Game) {
	g.ResetInstances()

	return
}

func (g *Game) appendUpdateDrawers(appends ...any) {
	for _, appnd := range appends {
		if v, ok := appnd.(Drawer); ok {
			g.Drawers = append(g.Drawers, v)
		}
		if v, ok := appnd.(Updater); ok {
			g.Updaters = append(g.Updaters, v)
		}
	}
}

func (g *Game) initOptions() {
	o := &g.Options
	o.Muted = &audio.Mute
}

func (g *Game) initGUI() {
	g.GUI.Mute = &audio.Mute
	g.GUI.Visible = true
	grid := &g.GUI.grid
	// https://www.reddit.com/r/programminghorror/comments/1fhgc1p/it_just_came_to_my_mind_that_i_could_do_this/
	// lol.
	// sudo: this incident will be reported.

	// Create the buttons grid for better management.
	for i := range cap(grid) {
		for j := range cap(grid[i]) {
			grid[i][j] = rl.Rectangle{
				Width:  200,
				Height: 50,
				X:      float(i) * 200,
				Y:      float(j) * 50,
			}
		}
	}
}

func (g *Game) initTextFields() {
	t := &g.Texts

	// Player Score.
	t.PlayerScore = NewText(
		strconv.Itoa(int(g.Player.Score)),
		50,
		rl.Gray,
		rl.Vector2{},
	)
	t.PlayerScore.Y = Height/2 - t.PlayerScore.Measure().Y/2
	t.PlayerScore.X = Width/2 + t.PlayerScore.Measure().X*3

	// CPU Score.
	t.CPUScore = NewText(
		strconv.Itoa(int(g.CPU.Score)),
		50,
		rl.Gray,
		rl.Vector2{},
	)
	t.CPUScore.Y = t.PlayerScore.Y
	t.CPUScore.X = Width/2 - t.CPUScore.Measure().X*4

	// Wait for play.
	t.Wait4Play = NewTextVPtr(
		"-- Press [SPACE] to start --",
		40,
		rl.White,
		rl.Vector2{},
		&g.Options.Waiting4Play,
	)
	t.Wait4Play.Y = Height/3 - t.Wait4Play.Measure().Y/2
	t.Wait4Play.X = Width/2 - t.Wait4Play.Measure().X/2

	// Muted.
	t.Muted = NewTextVPtr(
		"MUTED",
		40,
		rl.DarkGray,
		rl.Vector2{},
		&audio.Mute,
	)

	// Paused.
	t.Pause = NewTextVPtr(
		"-- PAUSED --",
		40,
		rl.White,
		rl.Vector2{},
		&g.Options.Paused,
	)
	t.Pause.Y = Height/3 - t.Pause.Measure().Y/2
	t.Pause.X = Width/2 - t.Pause.Measure().X/2

	t.Texts = append(t.Texts,
		&t.PlayerScore,
		&t.CPUScore,
		&t.Wait4Play,
		&t.Muted,
		&t.Pause,
	)
}
