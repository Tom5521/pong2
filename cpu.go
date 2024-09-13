//go:build ignore
// +build ignore

package main

type CPU struct {
	Paddle

	BallY *float
}

func NewCPU(ballYPtr *float) CPU {
	cpu := CPU{NewPaddle(), ballYPtr}

	cpu.Speed /= 2
	cpu.X = 0

	return cpu
}

func (c *CPU) Update() {
	if *c.BallY > c.Y {
		c.Y += c.Speed
	}
	if *c.BallY < c.Y {
		c.Y -= c.Speed
	}
	c.limitMovement()
}
