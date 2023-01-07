package desktop

import (
	"path"

	"github.com/gonutz/prototype/draw"
)

var bombermanSprite = path.Join("resources", "bomberman.png")

func DrawGameIteration(window draw.Window) {
	if window.WasKeyPressed(draw.KeyEscape) {
		window.Close()
	}

	err := window.DrawImageFile(bombermanSprite, 320, 320)
	if err != nil {
		panic(err)
	}
}
