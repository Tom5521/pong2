package main

import (
	"strconv"

	"pong2/audio"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewGame() Game {
	var g Game
	g.ResetInstances()

	return g
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
	o.Waiting4Play = true
}

func (g *Game) initTextFields() {
	t := &g.Texts

	// Player Score.
	t.PlayerScore = NewText(
		strconv.Itoa(g.Player.Score),
		50,
		rl.Gray,
		rl.Vector2{},
	)
	t.PlayerScore.Y = Height/2 - t.PlayerScore.Measure().Y/2
	t.PlayerScore.X = Width/2 + t.PlayerScore.Measure().X*3

	// CPU Score.
	t.CPUScore = NewText(
		strconv.Itoa(g.CPU.Score),
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
