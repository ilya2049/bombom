package desktop

import (
	"bombom/internal/game/input"
	"bombom/internal/pkg/event"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var keyMap = map[ebiten.Key]input.Key{
	ebiten.KeyLeft:  input.KeyLeft,
	ebiten.KeyRight: input.KeyRight,
	ebiten.KeyUp:    input.KeyUp,
	ebiten.KeyDown:  input.KeyDown,
	ebiten.KeySpace: input.KeySpace,
}

func readInputEvents() []event.Event {
	var events []event.Event

	for ebitenKey, gameKey := range keyMap {
		if inpututil.IsKeyJustPressed(ebitenKey) {
			events = append(events, event.New(input.KeyPressed{Key: gameKey}))
		}
		if inpututil.IsKeyJustReleased(ebitenKey) {
			events = append(events, event.New(input.KeyReleased{Key: gameKey}))
		}
	}

	return events
}
