package main

type (
	float = float32
	rint  = int32
	long  = int64
)

const (
	Width  = 800
	Height = 600
)

func main() {
	game := NewGame()
	game.Run()
}
