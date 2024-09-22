package main

import (
	"pong2/audio"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) Update() {
	g.Frames++

	if g.GUI.Visible {
		g.Ball.Update()
		return
	}

	switch {
	case rl.IsKeyPressed(rl.KeyR):
		g.ResetToDefaultState()
	case rl.IsKeyPressed(rl.KeyM):
		*g.Options.Muted = !*g.Options.Muted
	case rl.IsKeyPressed(rl.KeySpace):
		audio.Play("pause")
		if g.Options.Waiting4Play {
			g.Options.Waiting4Play = !g.Options.Waiting4Play
			break
		}
		g.Options.Paused = !g.Options.Paused
	}

	if g.Options.Paused || g.Options.Waiting4Play {
		return
	}

	for _, u := range g.Updaters {
		u.Update()
	}

	g.checkPoints()
	g.checkCollisions()
}
