package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Ball struct {
	rl.Vector2
	Radius float

	SpeedX, SpeedY float
}

func NewBall() Ball {
	return Ball{
		Vector2: rl.Vector2{
			X: Width / 2,
			Y: Height / 2,
		},

		Radius: 20,
		SpeedX: 7,
		SpeedY: 7,
	}
}

func (b *Ball) Update() {
	b.X += b.SpeedX
	b.Y += b.SpeedY

	if b.X+b.Radius >= Width || b.X-b.Radius <= 0 {
		b.SpeedX = -b.SpeedX
	}

	if b.Y+b.Radius >= Height || b.Y-b.Radius <= 0 {
		b.SpeedY = -b.SpeedY
	}
}

func (b *Ball) Draw() {
	rl.DrawCircleV(b.Vector2, b.Radius, rl.White)
}
