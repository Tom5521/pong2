package main

type (
	float = float32
	long  = int64
)

const (
	Width  = 800
	Height = 500
)

var devel bool

func main() {
	game := NewGame()
	game.Run()
}
