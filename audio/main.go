// TODO:This could be improved by using a map and ready to go.

package audio

import (
	"embed"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//go:embed assets
var assets embed.FS

var Mute bool

var S map[string]rl.Sound

var (
	Beep    rl.Sound
	Lose    rl.Sound
	Pause   rl.Sound
	Victory rl.Sound
)

func Load() {
	S = make(map[string]rl.Sound)

	dirs, _ := assets.ReadDir("assets")
	for _, d := range dirs {
		file, _ := assets.ReadFile("assets/" + d.Name())

		parts := strings.Split(d.Name(), ".")
		wave := rl.LoadWaveFromMemory("."+parts[len(parts)-1], file, int32(len(file)))

		sound := rl.LoadSoundFromWave(wave)
		S[parts[0]] = sound
	}
}

func Play(n string) {
	if Mute {
		return
	}
	rl.PlaySound(S[n])
}
