package desktop

import (
	"bombom/internal/game"
	"bombom/internal/pkg/event"

	"context"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewGame() *Game {
	eventDispatcher := event.NewDispatcher()

	game.RegisterKeyHandlers(eventDispatcher)

	return &Game{
		eventDispatcher: eventDispatcher,
	}
}

type Game struct {
	eventDispatcher *event.Dispatcher
}

func (g *Game) Update() error {
	var events []event.Event

	events = append(events, readInputEvents()...)

	return g.eventDispatcher.Dispatch(context.Background(), events...)
}

func (g *Game) Draw(_ *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
