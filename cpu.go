package main

type CPU struct {
	*Paddle

	ballYPtr *float
}

func NewCPU(ballYPtr *float) *CPU {
	cpu := CPU{NewPaddle(), ballYPtr}

	cpu.Speed /= 2
	cpu.X = 0

	return &cpu
}

func (c *CPU) Update() {
	if *c.ballYPtr > c.Y {
		c.Y += c.Speed
	}
	if *c.ballYPtr < c.Y {
		c.Y -= c.Speed
	}

	c.limitMovement()
}
