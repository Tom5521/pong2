package main

import (
	"strconv"

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

func (g *Game) checkCollisions() {
	// Checking for collisions

	if rl.CheckCollisionCircleRec(g.Ball.Vector2, g.Ball.Radius, g.CPU.Rectangle) {
		// Check if the ball hits the top or the bottom of the paddle.
		if g.Ball.Y < g.CPU.Y || g.Ball.Y > g.CPU.Y+g.CPU.Height {
			g.Ball.SpeedY *= -1
		}

		// Check if the ball hits the left/right side.
		if g.Ball.X < g.CPU.X || g.Ball.X > g.CPU.X+g.CPU.Width {
			g.Ball.SpeedX *= -1
		}

		audio.Play("beep")
	}
	if rl.CheckCollisionCircleRec(g.Ball.Vector2, g.Ball.Radius, g.Player.Rectangle) {
		// Check if the ball hits the top or the bottom of the paddle.
		if g.Ball.Y < g.Player.Y || g.Ball.Y > g.Player.Y+g.Player.Height {
			g.Ball.SpeedY *= -1
		}
		// Check if the ball hits the left/right side.
		if g.Ball.X < g.Player.X || g.Ball.X > g.Player.X+g.Player.Width {
			g.Ball.SpeedX *= -1
		}

		audio.Play("beep")
	}
}

func (g *Game) checkPoints() {
	player := rl.CheckCollisionCircleRec(
		g.Ball.Vector2,
		g.Ball.Radius,
		rl.Rectangle{
			Width:  1,
			Height: Height,
			X:      Width,
			Y:      0,
		},
	)

	cpu := rl.CheckCollisionCircleRec(
		g.Ball.Vector2,
		g.Ball.Radius,
		rl.Rectangle{
			Width:  1,
			Height: Height,
			X:      1,
			Y:      0,
		},
	)

	if player {
		g.CPU.Score++
		g.Texts.CPUScore.SetText(strconv.Itoa(g.CPU.Score))

		audio.Play("lose")
	}

	if cpu {
		g.Player.Score++
		g.Texts.PlayerScore.SetText(strconv.Itoa(g.Player.Score))

		audio.Play("victory")
	}
	if player || cpu {
		g.ResetToDefaultPositions()
		g.Options.Waiting4Play = true
	}
}
