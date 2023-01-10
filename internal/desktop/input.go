package desktop

import (
	"bombom/internal/game"
	"bombom/internal/pkg/event"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var keyMap = map[ebiten.Key]game.Key{
	ebiten.KeyLeft:  game.KeyLeft,
	ebiten.KeyRight: game.KeyRight,
	ebiten.KeyUp:    game.KeyUp,
	ebiten.KeyDown:  game.KeyDown,
}

func readInputEvents() []event.Event {
	var events []event.Event

	for ebitenKey, gameKey := range keyMap {
		if inpututil.IsKeyJustPressed(ebitenKey) {
			events = append(events, event.New(game.KeyPressed{Key: gameKey}))
		}
		if inpututil.IsKeyJustReleased(ebitenKey) {
			events = append(events, event.New(game.KeyReleased{Key: gameKey}))
		}
	}

	return events
}
