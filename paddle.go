package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Paddle struct {
	rl.Rectangle

	Speed float
	Score byte
}

func NewPaddle() *Paddle {
	var p Paddle

	p.Width = 20
	p.Height = 100

	p.X = Width - p.Width
	p.Y = Height/2 - p.Height/2

	p.Speed = 8

	return &p
}

func (p *Paddle) limitMovement() {
	if p.Y+p.Height >= Height {
		p.Y = Height - p.Height
	}
	if p.Y <= 0 {
		p.Y = 0
	}
}

func (p *Paddle) Update() {
	switch {
	case rl.IsKeyDown(rl.KeyUp), rl.IsKeyDown(rl.KeyW):
		p.Y -= p.Speed
	case rl.IsKeyDown(rl.KeyDown), rl.IsKeyDown(rl.KeyS):
		p.Y += p.Speed
	}

	p.limitMovement()
}

func (p Paddle) Draw() {
	rl.DrawRectangleRec(p.Rectangle, rl.White)
	rl.DrawRectangleLines(
		rint(p.X),
		rint(p.Y),
		rint(p.Width),
		rint(p.Height),
		rl.Black,
	)
}
