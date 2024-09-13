package main

type (
	float = float32
	rint  = int32 // TODO: Remove this.
	long  = int64
)

const (
	Width  = 800
	Height = 450
)

func main() {
	game := NewGame()
	game.Run()
}
