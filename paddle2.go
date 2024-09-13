package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Paddle2 struct {
	Paddle
}

func NewPaddle2() Paddle2 {
	p := Paddle2{NewPaddle()}
	p.X = 0

	return p
}

func (p *Paddle2) Update() {
	switch {
	case rl.IsKeyDown(rl.KeyW):
		p.Y -= p.Speed
	case rl.IsKeyDown(rl.KeyS):
		p.Y += p.Speed
	}

	p.limitMovement()
}
