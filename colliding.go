package main

import (
	"fmt"
	"strconv"

	"pong2/audio"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) checkCollisions() {
	// Checking for collisions
	switch {
	case rl.CheckCollisionCircleRec(g.Ball.Vector2, g.Ball.Radius, g.CPU.Rectangle):
		if !g.Ball.hitLast {
			// Check if the ball hits the top or the bottom of the paddle.
			if g.Ball.Y < g.CPU.Y || g.Ball.Y > g.CPU.Y+g.CPU.Height {
				g.Ball.SpeedY *= -1
				g.Ball.hitLast = true
			} else
			// Check if the ball hits the left/right side.
			if g.Ball.X < g.CPU.X || g.Ball.X > g.CPU.X+g.CPU.Width {
				g.Ball.SpeedX *= -1
				g.Ball.hitLast = true
			}
		}

		audio.Play("beep")
	case rl.CheckCollisionCircleRec(g.Ball.Vector2, g.Ball.Radius, g.Player.Rectangle):
		if !g.Ball.hitLast {
			// Check if the ball hits the top or the bottom of the paddle.
			if g.Ball.Y < g.Player.Y || g.Ball.Y > g.Player.Y+g.Player.Height {
				g.Ball.SpeedY *= -1
				g.Ball.hitLast = true
			} else
			// Check if the ball hits the left/right side.
			if g.Ball.X < g.Player.X || g.Ball.X > g.Player.X+g.Player.Width {
				g.Ball.SpeedX *= -1
				g.Ball.hitLast = true
			}
		}

		audio.Play("beep")
	// devel.
	case devel && g.Ball.hitLast:
		if g.Ball.hitLast {
			g.Ball.hits++
		}

		text := fmt.Sprintf(`hits: %v  |//^\\|  state: %v`, g.Ball.hits, g.Ball.hitLast)
		measure := rl.MeasureTextEx(rl.GetFontDefault(), text, 20, 2)
		color := rl.Red
		if g.Ball.hitLast {
			color = rl.Green
		}
		rl.DrawTextEx(
			rl.GetFontDefault(),
			text,
			rl.Vector2{
				X: Width - measure.X,
				Y: Height - measure.Y,
			},
			20,
			2,
			color,
		)
		fallthrough // Run default.
	default:
		g.Ball.hitLast = false
	}
}

func (g *Game) checkPoints() {
	prect := rl.Rectangle{
		Width:  1,
		Height: Height,
		X:      Width,
		Y:      0,
	}
	player := rl.CheckCollisionCircleRec(
		g.Ball.Vector2,
		g.Ball.Radius,
		prect,
	)

	crect := rl.Rectangle{
		Width:  1,
		Height: Height,
		X:      0,
		Y:      0,
	}
	cpu := rl.CheckCollisionCircleRec(
		g.Ball.Vector2,
		g.Ball.Radius,
		crect,
	)

	if devel {
		rl.DrawRectangleRec(prect, rl.Red)
		rl.DrawRectangleRec(crect, rl.Red)
	}

	if player {
		g.CPU.Score++
		g.Texts.CPUScore.SetText(strconv.Itoa(int(g.CPU.Score)))

		audio.Play("lose")
	}

	if cpu {
		g.Player.Score++
		g.Texts.PlayerScore.SetText(strconv.Itoa(int(g.Player.Score)))

		audio.Play("victory")
	}
	if player || cpu {
		g.ResetToDefaultPositions()
		g.Options.Waiting4Play = true

		// Invert ball direction if the CPU wins a point.
		if player {
			g.Ball.SpeedX *= -1
			g.Ball.SpeedY *= -1
		}
	}
}
