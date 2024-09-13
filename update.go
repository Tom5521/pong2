package main

import (
	"strconv"

	"pong2/audio"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) Update() {
	g.Frames++

	g.Player2.Update()
	g.Player.Update()
	g.Ball.Update()

	g.checkPoints()
	g.checkCollisions()
}

func (g *Game) checkCollisions() {
	// Checking for collisions

	if rl.CheckCollisionCircleRec(g.Ball.Vector2, g.Ball.Radius, g.Player2.Rectangle) {
		// Check if the ball hits the top or the bottom of the paddle.
		if g.Ball.Y < g.Player2.Y || g.Ball.Y > g.Player2.Y+g.Player2.Height {
			g.Ball.SpeedY *= -1
		}

		// Check if the ball hits the left/right side.
		if g.Ball.X < g.Player2.X || g.Ball.X > g.Player2.X+g.Player2.Width {
			g.Ball.SpeedX *= -1
		}

		audio.Play(audio.Beep)
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

		audio.Play(audio.Beep)
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
		g.Player2.Score++
		g.Texts.CPUScore.SetText(strconv.Itoa(g.Player2.Score))

		audio.Play(audio.Lose)
	}

	if cpu {
		g.Player.Score++
		g.Texts.PlayerScore.SetText(strconv.Itoa(g.Player.Score))

		audio.Play(audio.Victory)
	}
	// TODO: Finish this.
	// if player || cpu {
	// 	g.ResetToDefaultState()
	// }
}
