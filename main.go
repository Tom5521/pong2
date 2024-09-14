package main

type (
	float = float32
	long  = int64
)

const (
	Width  = 800
	Height = 500
)

func main() {
	game := NewGame()
	game.Run()
}
