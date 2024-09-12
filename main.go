package main

type (
	float = float32
	rint  = int32
	long  = int64
)

const (
	Width  = 800
	Height = 400
)

func main() {
	game := NewGame()
	game.Run()
}
